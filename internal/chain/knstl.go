package chain

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/cosmos/cosmos-sdk/client/tx"
	cryptokeyring "github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/golang/protobuf/proto"
	"github.com/konstellation/swap/internal/config"
	keyring "github.com/konstellation/swap/internal/key"
	"github.com/konstellation/swap/internal/model"
	"github.com/konstellation/swap/internal/mongo"
	"github.com/konstellation/swap/internal/util"
	"github.com/shopspring/decimal"
	tenderminthttp "github.com/tendermint/tendermint/rpc/client/http"
	tendermintrpctypes "github.com/tendermint/tendermint/rpc/core/types"
	mongodrv "go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

const amountKnstlUnit = "1000000"

var (
	cryptoKeyring cryptokeyring.Keyring
	keyringInfo   cryptokeyring.Info
	Denom         = "udarc"
	ChainID       = "darchub"
)

type KnstlConnection struct {
	conn         *tenderminthttp.HTTP
	resChan      <-chan tendermintrpctypes.ResultEvent
	bscConn      *BSCConnection
	MongoDB      *mongo.Connection
	swapAddr     string
	msgChan      chan string
	keyring      cryptokeyring.Info
	knstlGrpcUrl string
	knstlUrl     string
}

func (k *KnstlConnection) InitConnection(_ context.Context, c *config.KnstlInfo, mg *mongo.Connection, bscConn *BSCConnection, msgChan chan string) error {
	var err error

	cryptoKeyring, err = keyring.NewKeyring("portal")
	if err != nil {
		log.Fatal(err)
	}

	keyringInfo, err = keyring.CreateKey(cryptoKeyring, "portal_keyring", c.KnstlSwapMnemonic)
	if err != nil {
		log.Fatalf("Failed to create keyring: %v", err)
	}

	k.bscConn = bscConn
	k.msgChan = msgChan
	k.swapAddr = c.KnstlSwapAddr
	k.keyring = keyringInfo
	k.MongoDB = mg
	k.knstlGrpcUrl = c.KnstlNodeGrpcUrl
	k.knstlUrl = c.KnstlNodeUrl

	k.conn, err = tenderminthttp.New(k.knstlUrl, "/websocket")
	if err != nil {
		log.Fatalln("Konstellation: Failed to initialize RPC Connection: ", err)
	}

	if err := k.conn.Start(); err != nil {
		log.Fatalln("Konstellation: Failed to initialize Websocket Connection: ", err)
	}

	log.Println("Konstellation: RPC Websocket connection established")

	query := fmt.Sprintf(`tm.event='Tx' AND transfer.recipient = '%s'`, k.swapAddr)
	k.resChan, err = k.conn.WSEvents.Subscribe(
		context.Background(),
		"swap",
		query,
	)
	if err != nil {
		log.Fatalln("Konstellation: Failed to subscribe: ", err)
	}

	log.Println(fmt.Sprintf("Konstellation: Subscribed to %s", query))

	return err
}

type Log struct {
	Events []struct {
		Type       string `json:"type"`
		Attributes []struct {
			Key   string `json:"key"`
			Value string `json:"value"`
		} `json:"attributes"`
	} `json:"events"`
}

func (k *KnstlConnection) HandleMessage() {
	for {
		msg := <-k.resChan
		log.Printf("****** Get source knstl transaction data: %+v ******\n", msg.Events)
		if len(msg.Events["transfer.amount"]) == 1 {
			log.Println("There is no fee for transaction")
			continue
		}
		amountStr := strings.ReplaceAll(msg.Events["transfer.amount"][1], "udarc", "")
		amount, err := decimal.NewFromString(amountStr)
		if err != nil {
			log.Println(err)
			continue
		}
		log.Println("Knstl amount decimal conversion: ", amount)
		knstlUnit, err := decimal.NewFromString(amountKnstlUnit)
		if err != nil {
			log.Println(err)
			continue
		}
		log.Println("Knstl basic unit amount decimal conversion: ", knstlUnit)
		amountInDB := amount.Div(knstlUnit)

		log.Println("Checking if address in POST request is blacklist address")
		filter := map[string]string{
			"address": msg.Events["message.sender"][0],
		}
		blacklistResult, err := k.MongoDB.FindBlacklist(filter)
		if err != nil {
			if err != mongodrv.ErrNoDocuments {
				log.Println(err)
				continue
			}
		}
		isblacklistAmountbigger := false
		if blacklistResult != nil {
			threshold, _ := decimal.NewFromString(util.BlacklistAllowThresholdAmount)
			if amountInDB.GreaterThanOrEqual(threshold) {
				err = fmt.Errorf("blacklist address amount request is more than 1000000 DARC")
				log.Println(err)
				isblacklistAmountbigger = true
			}
		}

		filter = map[string]string{
			"from_address":                  msg.Events["message.sender"][0],
			"source_network":                "knstl",
			"destination_network":           "bsc",
			"source_network_completed":      "false",
			"destination_network_completed": "false",
			"amount":                        amountInDB.String(),
		}
		log.Printf("Knstl transaction data to search in the DB: %+v\n", filter)
		result, err := k.MongoDB.FindTx(filter) // this case cannot update transaction complete case
		if err != nil {
			log.Println(err)
			continue
		}
		log.Printf("Knstl transaction POST request is found in DB: %+v\n", result)
		tx, ok := result.(model.Tx) // this case cannot update transaction complete case
		if !ok {
			log.Println("The result is not transaction type")
			continue
		}
		log.Printf("Knstl transaction data in the DB: %+v\n", tx)
		tx.SourceNetworkHash = msg.Events["tx.hash"][0]
		tx.SourceNetworkCompleted = true
		tx.UpdatedAt = time.Now().Format(util.TimeFormat)

		if isblacklistAmountbigger { // Update transaction status is completed
			tx.Completed = true
		}
		_, err = k.MongoDB.UpdateTx(&tx) // this case cannot update transaction complete case
		if err != nil {
			log.Println(err)
			continue
		}
		if isblacklistAmountbigger { // Protect blacklist swap
			log.Println("blacklist address amount request is more than 1000000 DARC. Cannot conitnue to swap")
			continue
		}

		log.Printf("Knstl transaction data is updated in the DB: %+v\n", tx)
		log.Println("The knstl source network operation is finished. $$$$$$")
		k.bscConn.DisburseFunds(&tx)
	}
}

func (k *KnstlConnection) DisburseFunds(t *model.Tx) {
	log.Println("Knstl token conversion start. Destination network operation. $$$$$$")
	toAddress := t.ToAddress
	toAddr, err := types.AccAddressFromBech32(toAddress)
	if err != nil {
		log.Println("Invalid corporate wallet:", t.ToAddress, err)
		k.updateTransactionComplete(t)
		return
	}
	log.Println("Knstl toaddress", toAddress, "Knstl AccAddressFromBech with toaddress", toAddr)
	corporateWallet, err := types.AccAddressFromBech32(k.swapAddr)
	if err != nil {
		log.Println("Invalid corporate wallet:", k.swapAddr, err)
		k.updateTransactionComplete(t)
		return
	}
	log.Println("Knstl swapAddr(KNSTL_CORPORATE_ADDR)", k.swapAddr, "Knstl AccAddressFromBech with swapAddr", corporateWallet)
	knstlGrpcConn, err := grpc.Dial(k.knstlGrpcUrl, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Printf("Did not connect: %v", err)
		k.updateTransactionComplete(t)
		return
	}
	log.Println("knstlGrpcConn", k.knstlGrpcUrl, "is connected")
	defer knstlGrpcConn.Close()
	knstlAcc, err := getAccount(context.Background(), knstlGrpcConn, keyringInfo.GetAddress().String())
	if err != nil {
		log.Printf("Failed to get account %s: %v", keyringInfo.GetAddress().String(), err)
		k.updateTransactionComplete(t)
		return
	}
	decimalAmount := util.GetTransactionAmount("knstl", t.Amount)
	log.Println("user transaction total knstl amount:", decimalAmount)
	//decimalKnstlFee := util.GetTransactionAmount("knstl", float64(util.UserKnstlTransactionFee))
	//log.Println("user transaction fee knstl amount:", decimalKnstlFee)
	//totalAmountDecimal := decimalAmount.Sub(decimalKnstlFee)
	//log.Println("Total decimal amount deducting fee:", totalAmountDecimal)

	msg := banktypes.NewMsgSend(corporateWallet, toAddr, types.NewCoins(types.NewCoin("udarc", types.NewIntFromUint64(decimalAmount.BigInt().Uint64()))))
	err = msg.ValidateBasic()
	if err != nil {
		log.Printf("Invalid tx msg: %v", err)
		k.updateTransactionComplete(t)
		return
	}
	log.Printf("Knstl tx msg: %+v\n", msg)
	encCfg := simapp.MakeTestEncodingConfig()
	txBuilder := encCfg.TxConfig.NewTxBuilder()
	txBuilder.SetGasLimit(140000)
	txBuilder.SetFeeAmount(types.NewCoins(types.NewInt64Coin("udarc", int64(1))))
	err = txBuilder.SetMsgs(msg)
	if err != nil {
		log.Println("Setting messages is failed")
		k.updateTransactionComplete(t)
		return
	}
	txFactory := tx.Factory{}
	txFactory = txFactory.
		WithKeybase(cryptoKeyring).
		WithTxConfig(encCfg.TxConfig).
		WithAccountNumber(knstlAcc.AccountNumber).
		WithSequence(knstlAcc.Sequence).
		WithChainID(ChainID)
	log.Printf("feepayer: %+v\n", txBuilder.GetTx().FeePayer())
	log.Printf("fee: %+v\n", txBuilder.GetTx().GetFee())
	if err := tx.Sign(txFactory, keyringInfo.GetName(), txBuilder, true); err != nil {
		log.Println("Signing transaction is failed")
		k.updateTransactionComplete(t)
		return
	}
	txBytes, err := encCfg.TxConfig.TxEncoder()(txBuilder.GetTx())
	if err != nil {
		log.Println("Encoding transaction is failed")
		k.updateTransactionComplete(t)
		return
	}
	res, err := k.conn.BroadcastTxSync(context.Background(), txBytes)
	if err != nil {
		log.Println("Broadcasting transaction is failed:", err)
		k.updateTransactionComplete(t)
		return
	}
	log.Printf("Knstl transaction response: %+v\n", *res)

	// Confirm that knstl transaction is done
	t.DestinationNetworkHash = res.Hash.String()
	ok := false
	transactionCheckTryCount := 0
	log.Println("Start knstl transaction confirmation check")
	for !ok {
		ok, err = k.IsTransactionSuccessful(res.Hash.String())
		notFoundErrMsg := fmt.Sprintf("Error Message: Internal error, Error Data: tx (%s) not found", res.Hash.String())
		if err != nil && !strings.Contains(err.Error(), notFoundErrMsg) {
			log.Println("Knstl transaction sent has error: ", err)
			k.updateTransactionComplete(t)
			return
		}
		transactionCheckTryCount++
	}
	log.Println("Knstl transaction confirmation check total try: ", transactionCheckTryCount)
	log.Println("Finished knstl transaction confirmation check")

	log.Println("Complete knstl transaction:", res.Hash.String())
	t.DestinationNetworkCompleted = true
	log.Println("Destination network hash:", t.DestinationNetworkHash, "Destination network completed", t.DestinationNetworkCompleted)
	k.updateTransactionComplete(t)
	log.Printf("transaction is updated in the DB: %+v\n", *t)

	k.msgChan <- "###### Disbursing knstl funds for transaction: " + t.ID.String() + " ######"
}

func (k *KnstlConnection) GetTx(hash string) (map[string]interface{}, error) {
	txHash := common.HexToHash(hash)
	knstlTransactionCheckUrl := k.knstlUrl + "/tx?hash=" + txHash.String()
	log.Println("knstlTxUrl:", knstlTransactionCheckUrl)
	tx, err := http.Get(knstlTransactionCheckUrl)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var txResult map[string]interface{}
	err = json.NewDecoder(tx.Body).Decode(&txResult)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer tx.Body.Close()
	return txResult, nil
}

func getAccount(ctx context.Context, conn *grpc.ClientConn, addr string) (*authtypes.BaseAccount, error) {
	authClient := authtypes.NewQueryClient(conn)
	r := &authtypes.QueryAccountRequest{Address: addr}
	acc, err := authClient.Account(ctx, r)
	if err != nil {
		return nil, err
	}

	bacc := &authtypes.BaseAccount{}
	if err := proto.Unmarshal(acc.Account.GetValue(), bacc); err != nil {
		return nil, err
	}

	return bacc, nil
}

func (k *KnstlConnection) updateTransactionComplete(tx *model.Tx) {
	tx.Completed = true
	tx.UpdatedAt = time.Now().Format(util.TimeFormat)
	_, err := k.MongoDB.UpdateTx(tx)
	if err != nil {
		log.Println(err)
		return
	}
}

func (k *KnstlConnection) IsTransactionSuccessful(hash string) (bool, error) {
	knstlResult, err := k.GetTx(hash)
	if err != nil {
		err := fmt.Errorf("Knstl transaction " + err.Error())
		return false, err
	}
	log.Println("KnstlResult:", knstlResult)
	if _, ok := knstlResult["error"]; ok {
		txErr, _ := knstlResult["error"].(map[string]interface{})
		errMsg, _ := txErr["message"].(string)
		errData, _ := txErr["data"].(string)
		err := fmt.Errorf("Error Message: %s, Error Data: %s\n", errMsg, errData)
		return false, err
	} else {
		result, _ := knstlResult["result"].(map[string]interface{})
		txResult, _ := result["tx_result"].(map[string]interface{})
		logStr, _ := txResult["log"].(string)
		if !strings.Contains(logStr, "events") {
			err := fmt.Errorf(logStr)
			return false, err
		}
	}
	return true, nil
}
