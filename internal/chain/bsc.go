package chain

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/gorilla/websocket"
	"github.com/konstellation/swap/internal/config"
	"github.com/konstellation/swap/internal/model"
	"github.com/konstellation/swap/internal/mongo"
	BEP20Token "github.com/konstellation/swap/internal/types"
	"github.com/konstellation/swap/internal/util"
	"github.com/shopspring/decimal"
	"golang.org/x/crypto/sha3"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gammazero/deque"
	mongodrv "go.mongodb.org/mongo-driver/mongo"
)

const (
	amountBscUnit      = "1000000000000000000"
	userTransactionFee = "0.0001"
)

type BSCConnection struct {
	client                *ethclient.Client
	konConn               *KnstlConnection
	MongoDB               *mongo.Connection
	headerChan            chan *types.Header
	logChan               chan types.Log
	ctx                   context.Context
	sub                   ethereum.Subscription
	msgChan               chan string
	transactionScanApiUrl string
	InputChan             chan *model.Tx
	q                     deque.Deque

	signer           types.Signer
	corporateAddress common.Address
	contractAddress  common.Address
	privKey          *ecdsa.PrivateKey
	pubKey           *ecdsa.PublicKey
	contractAbi      abi.ABI
}

type TransactionData struct {
	SwapNonce int `json:"swap_nonce"`
}

func (b *BSCConnection) InitConnection(ctx context.Context, c *config.BscInfo, mg *mongo.Connection, konConn *KnstlConnection, msgChan chan string) error {
	var err error
	b.client, err = ethclient.Dial(c.BscNodeUrl)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("BSC node connected:", c.BscNodeUrl)

	b.ctx = context.Background()

	b.corporateAddress = common.HexToAddress(c.BscCorporateAddr)
	b.contractAddress = common.HexToAddress(c.BEP20ContractAddr)
	b.konConn = konConn
	b.logChan = make(chan types.Log)
	b.InputChan = make(chan *model.Tx)
	b.msgChan = msgChan
	b.MongoDB = mg
	b.transactionScanApiUrl = c.BscTransactionScanApiUrl

	query := ethereum.FilterQuery{
		Addresses: []common.Address{b.contractAddress},
	}

	sub, err := b.client.SubscribeFilterLogs(ctx, query, b.logChan)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("BSC: subscribed to %s\n", query.Addresses)

	b.sub = sub

	b.privKey, err = crypto.HexToECDSA(c.BscCorporateAddrPrivKey)
	// b.privKey, err = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Fatalln(err)
	}
	b.pubKey = (b.privKey.Public()).(*ecdsa.PublicKey)
	b.signer = types.NewEIP155Signer(big.NewInt(1000))

	contractAbi, err := abi.JSON(strings.NewReader(BEP20Token.BEP20TokenMetaData.ABI))
	if err != nil {
		log.Fatal(err)
	}
	b.contractAbi = contractAbi

	return nil
}

func (b *BSCConnection) FillInputData() {
	for {
		select {
		case inputData := <-b.InputChan:
			b.q.PushBack(inputData)
		}
	}
}

func (b *BSCConnection) StoreTransactions() {
	for {
		select {
		case err := <-b.sub.Err():
			log.Println(err)
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				c := config.New()
				err := b.InitConnection(b.ctx, c.SwapInfo.Bsc, b.MongoDB, b.konConn, b.msgChan)
				if err != nil {
					log.Println(err)
					// TODO: what do we do if reconnection is failed?
				}
			}
		case vLog := <-b.logChan:
			amountBscTransaction := b.getAmount(vLog.Data)
			log.Println("###### Get source bsc transaction data:", vLog, ", amount:", amountBscTransaction, "######")
			if strings.ToLower("0x"+vLog.Topics[2].String()[26:]) != strings.ToLower(crypto.PubkeyToAddress(*b.pubKey).String()) {
				log.Println("Not swap transaction")
				continue
			}
			_, err := b.MongoDB.InsertBscTx(&vLog)
			if err != nil {
				log.Println(err)
			}
		}
	}
}

func (b *BSCConnection) HandleMessage() {
	for {
		time.Sleep(1 * time.Minute)
		if b.q.Len() > 0 {
			front := b.q.PopFront()
			target, _ := front.(*model.Tx)
			log.Printf("*********** queue front target: %+v\n", target)
			has_match := false
			filter := map[string]string{
				"removed": "false",
			}
			cur, err := b.MongoDB.FindBscTx(filter) // this case cannot update transaction complete case
			if err != nil {
				log.Println(err)
				continue
			}
			defer cur.Close(b.MongoDB.Ctx)
			for cur.Next(b.MongoDB.Ctx) {
				result := types.Log{}
				err := cur.Decode(&result)
				if err != nil {
					log.Println(err)
					break
				}
				log.Printf("*********** bsc tx status: %+v\n", result)
				amountBscTransaction := b.getAmount(result.Data)
				targetAmount := decimal.NewFromFloat(target.Amount)
				log.Println("********** amountBscTransaction", amountBscTransaction, "targetAmount", targetAmount, "equal:", targetAmount.Equal(amountBscTransaction))
				if strings.ToLower(target.FromAddress) == strings.ToLower("0x"+result.Topics[1].String()[26:]) && targetAmount.Equal(amountBscTransaction) {
					isBlackList, err := b.processTransaction(target, &result)
					if err == nil || isBlackList {
						result.Removed = true
						_, err := b.MongoDB.UpdateBscTx(&result) // this case cannot update transaction complete case
						if err != nil {
							log.Println(err)
							break
						}
						has_match = true
					}
				} else {
					continue
				}
			}
			target.TxTryCount++
			if target.TxTryCount == 20 {
				log.Printf("*********** Transaction %+v timeout!", target)
				amount := decimal.NewFromFloat(target.Amount)
				filter = map[string]string{
					"from_address":                  target.FromAddress,
					"to_address":                    target.ToAddress,
					"source_network":                "bsc",
					"source_network_completed":      "false",
					"destination_network_completed": "false",
					"destination_network":           "knstl",
					"created_at":                    target.CreatedAt,
					"amount":                        amount.String(),
				}
				result, _ := b.MongoDB.FindTx(filter)
				tx, _ := result.(model.Tx)
				tx.Completed = true
				_, _ = b.MongoDB.UpdateTx(&tx)
			}
			if !has_match && target.TxTryCount < 20 {
				log.Printf("*********** Unfinished tx: %+v\n", target)
				b.q.PushBack(target)
			}
		}
	}
}

func (b *BSCConnection) processTransaction(inputData *model.Tx, vLog *types.Log) (isBlackList bool, err error) {
	log.Printf("###### Get source bsc transaction data: %+v ######\n", vLog)

	from := vLog.Topics[1]
	fromAddr := common.BytesToAddress(from[len(from)-20:])
	// Bsc transaction listener gets the transaction after knstl->bsc transaction.
	// The listener can detect the bsc from knstl->bsc transaction because knstl->bsc ends with bsc transaction.
	// This needs to distingish from normal bsc->knstl transaction.
	// To handle that, if the source address is equal to BSC_CORPORATE_ADDR then stops transaction flow.
	amountInDB := b.getAmount(vLog.Data)
	if fromAddr == crypto.PubkeyToAddress(*b.pubKey) { // BSC_CORPORATE_ADDR
		err := fmt.Errorf("###### Bsc -> Knstl transaction with the amount %+v is successful. ######", amountInDB)
		log.Printf("###### %s ######", err.Error())
		return false, err
	}
	log.Println("from address: ", fromAddr)

	log.Println("Checking if address in POST request is blacklist address")
	filter := map[string]string{
		"address": fromAddr.String(),
	}
	blacklistResult, err := b.MongoDB.FindBlacklist(filter)
	if err != nil {
		if err != mongodrv.ErrNoDocuments {
			log.Println(err)
			return false, err
		}
	}
	isblacklistAmountbigger := false
	if blacklistResult != nil {
		threshold, _ := decimal.NewFromString(util.BlacklistAllowThresholdAmount)
		if amountInDB.GreaterThanOrEqual(threshold) {
			log.Println("blacklist address amount request is more than 1000000 DARC")
			isblacklistAmountbigger = true
		}
	}

	filter = map[string]string{
		"from_address":                  fromAddr.String(),
		"to_address":                    inputData.ToAddress,
		"source_network":                "bsc",
		"source_network_completed":      "false",
		"destination_network_completed": "false",
		"destination_network":           "knstl",
		"created_at":                    inputData.CreatedAt,
		"amount":                        amountInDB.String(),
	}
	log.Printf("Bsc transaction data to search in the DB: %+v\n", filter)
	result, err := b.MongoDB.FindTx(filter) // this case cannot update transaction complete case
	if err != nil {
		log.Println(err)
		return false, err
	}
	log.Printf("Bsc transaction POST request is found in DB: %+v\n", result)
	tx, ok := result.(model.Tx) // this case cannot update transaction complete case
	if !ok {
		log.Println("The result is not transaction type")
		return false, err
	}
	tx.SourceNetworkHash = vLog.TxHash.String()
	tx.SourceNetworkCompleted = true
	tx.UpdatedAt = time.Now().Format(util.TimeFormat)

	if isblacklistAmountbigger { // Update transaction status is completed
		tx.Completed = true
	}
	_, err = b.MongoDB.UpdateTx(&tx) // this case cannot update transaction complete case
	if err != nil {
		log.Println(err)
		return false, err
	}
	if isblacklistAmountbigger { // Protect blacklist swap
		log.Println("blacklist address amount request is more than 1000000 DARC. Cannot conitnue to swap")
		return true, nil
	}
	log.Printf("Bsc transaction data is updated in the DB: %+v\n", result)
	log.Println("The bsc source network operation is finished. $$$$$$")
	b.konConn.DisburseFunds(&tx)
	return false, nil
}

func (b *BSCConnection) getAmount(data []byte) decimal.Decimal {
	ev, _ := b.contractAbi.Unpack("Transfer", data) // this case cannot update transaction complete case
	amountRaw := ev[0]
	log.Println("amount raw: ", amountRaw)
	amountBigInt, _ := amountRaw.(*big.Int)
	log.Println("Bsc amount float conversion: ", amountBigInt)
	amount, _ := decimal.NewFromString(amountBigInt.String())
	log.Println("Bsc amount decimal conversion: ", amount)
	bscUnit, _ := decimal.NewFromString(amountBscUnit)
	log.Println("Bsc basic unit amount decimal conversion: ", bscUnit)
	return amount.Div(bscUnit)
}

func (b *BSCConnection) DisburseFunds(t *model.Tx) {
	// Reference: https://goethereumbook.org/transfer-eth/
	log.Println("bsc token conversion start. Destination network operation. $$$$$$")
	toAddr := t.ToAddress

	fromAddress := crypto.PubkeyToAddress(*b.pubKey)
	log.Println("Bsc From Address(BSC_CORPORATE_ADDR):", fromAddress)
	nonce, err := b.client.PendingNonceAt(b.ctx, fromAddress)
	if err != nil {
		log.Println(err)
		b.updateTransactionComplete(t)
		return
	}
	log.Println("Nonce:", nonce)
	value := big.NewInt(0)
	gasPrice, err := b.client.SuggestGasPrice(b.ctx)
	if err != nil {
		log.Println(err)
		b.updateTransactionComplete(t)
		return
	}
	log.Println("Gas price:", gasPrice)
	toAddrChecked := common.HexToAddress(toAddr)
	log.Println("ToAddress:", toAddr, "toAddrHexToAddress", toAddrChecked)
	gasLimit := uint64(76708)

	transferFnSignature := []byte("transfer(address,uint256)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]
	log.Println("Method id:", hexutil.Encode(methodID))

	paddedToAddress := common.LeftPadBytes(toAddrChecked.Bytes(), 32)
	log.Println("PaddedToAddress:", hexutil.Encode(paddedToAddress))

	decimalAmount := util.GetTransactionAmount("bsc", t.Amount)
	log.Println("user transaction total bsc amount:", decimalAmount)
	decimalBscFee := util.GetTransactionAmount("bsc", util.UserBscTransactionFee)
	log.Println("user transaction fee bsc amount:", decimalBscFee)
	totalAmountDecimal := decimalAmount.Sub(decimalBscFee)
	log.Println("Total decimal amount deducting fee:", totalAmountDecimal)

	amount := new(big.Int)
	amount.SetString(totalAmountDecimal.String(), 10)
	log.Println("Total amount str:", amount.String())
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	log.Println("PaddedAmount:", hexutil.Encode(paddedAmount))
	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedToAddress...)
	data = append(data, paddedAmount...)

	tx := types.NewTransaction(nonce, b.contractAddress, value, gasLimit, gasPrice, data)
	log.Printf("Transaction: %+v\n", tx)
	chainID, err := b.client.NetworkID(context.Background())
	if err != nil {
		log.Println(err)
		b.updateTransactionComplete(t)
		return
	}
	log.Printf("Chain id: %+v\n", chainID)
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), b.privKey)
	if err != nil {
		log.Println("Error creating and signing transaction: ", err)
		b.updateTransactionComplete(t)
		return
	}
	log.Printf("Signed tx: %+v\n", signedTx.Hash().Hex())

	// Sometimes nonce calculation is not accurate and it can get old hash.
	// To making transaction in this case, check the hash and stop proceeding.
	filter := map[string]string{
		"from_address":                  t.FromAddress,
		"source_network_completed":      "true",
		"destination_network_hash":      signedTx.Hash().Hex(),
		"destination_network_completed": "true",
	}
	row, err := b.MongoDB.FindTx(filter)
	if err != nil {
		if err != mongodrv.ErrNoDocuments {
			log.Println(err)
			b.updateTransactionComplete(t)
			return
		}
	}
	if row != nil {
		log.Println("Failed to send transaction: tx is signed with old nonce. This means that tx is considered as old tx existed.")
		b.updateTransactionComplete(t)
		return
	}

	// if new signed transaction with new nonce, proceed transaction.
	err = b.client.SendTransaction(b.ctx, signedTx)
	if err != nil {
		log.Println("Failed to send transaction: ", err)
		b.updateTransactionComplete(t)
		return
	}

	// Confirm that bsc transaction is done
	t.DestinationNetworkHash = signedTx.Hash().Hex()
	status := "0"
	transactionCheckTryCount := 0
	log.Println("Start bsc transaction confirmation check")
	for status != "0" {
		// https://docs.bscscan.com/support/common-error-messages
		// api rate has limit. Thus, put 30 seconds term
		time.Sleep(util.SleepTimeSeconds * time.Second)
		status, err = b.getTransactionStatus(signedTx.Hash().Hex())
		if err != nil {
			log.Println("Bsc transaction sent has error: ", err)
			b.updateTransactionComplete(t)
			return
		}
		transactionCheckTryCount++
	}
	log.Println("Bsc transaction confirmation check total try: ", transactionCheckTryCount)
	log.Println("Finished bsc transaction confirmation check")

	log.Println("Complete bsc transaction:", signedTx.Hash().Hex())
	t.DestinationNetworkCompleted = true
	log.Println("Destination network hash:", t.DestinationNetworkHash, "Destination network completed", t.DestinationNetworkCompleted)
	b.updateTransactionComplete(t)
	log.Printf("transaction is updated in the DB: %+v\n", *t)

	b.msgChan <- "****** Disbursing bsc funds for transaction: " + t.ID.String() + " ******"
}

func (b *BSCConnection) GetTx(hash string) (map[string]interface{}, error) {
	txHash := common.HexToHash(hash)
	bscTransactionCheckUrl := fmt.Sprintf(b.transactionScanApiUrl, txHash.String())
	log.Println("BscTxUrl:", bscTransactionCheckUrl)
	tx, err := http.Get(bscTransactionCheckUrl)
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

func (b *BSCConnection) updateTransactionComplete(tx *model.Tx) {
	tx.Completed = true
	tx.UpdatedAt = time.Now().Format(util.TimeFormat)
	_, err := b.MongoDB.UpdateTx(tx)
	if err != nil {
		log.Println(err)
	}
}

func (b *BSCConnection) IsTransactionSuccessful(hash string) (bool, error) {
	status, err := b.getTransactionStatus(hash)
	if err != nil {
		log.Println(err)
		return false, err
	}
	log.Println("Bsc transaction status:", status)
	if status == "0" {
		err := fmt.Errorf("bsc transaction is failed")
		return false, err
	}
	return true, nil
}

func (b *BSCConnection) getTransactionStatus(hash string) (string, error) {
	bscResult, err := b.GetTx(hash)
	if err != nil {
		err := fmt.Errorf("bsc transaction " + err.Error())
		return "", err
	}
	result, _ := bscResult["result"].(map[string]interface{})
	status, _ := result["status"].(string)
	return status, nil
}
