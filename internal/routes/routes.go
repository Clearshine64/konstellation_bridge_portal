package routes

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/konstellation/swap/internal/errors"
	"github.com/konstellation/swap/internal/httpserver"
	"github.com/konstellation/swap/internal/logger"
	"github.com/konstellation/swap/internal/model"
	"github.com/konstellation/swap/internal/util"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	bsc   = "bsc"
	knstl = "knstl"
)

func InitRoutes(e *echo.Echo) {
	e.POST("/tx", addTx)
	e.POST("/blacklist", addBlacklist)
	e.GET("/tx/:id", getTx)
	e.GET("/log", getLog)
	//e.DELETE("/tx/:id", abortTx)
}

func addTx(ctx echo.Context) error {
	log.Println("====== Get POST transaction request ======")
	cctx := ctx.Get("cctx").(*httpserver.CCtx)

	i := new(TxInput)
	if err := ctx.Bind(i); err != nil {
		err = errors.Prepare(errors.ECBindError, err)
		log.Println(err)
		return ctx.JSON(http.StatusOK, &Response{
			Result:  err.Error(),
			Success: false,
		})
	}
	if err := i.Validate(); err != nil {
		log.Println(err)
		return ctx.JSON(http.StatusOK, &Response{
			Result:  err.Error(),
			Success: false,
		})
	}

	log.Printf("POST request: %+v\n", *i)

	if numDecPlaces(i.Amount) > 5 {
		err := errors.PreparePayload(errors.ECTxInsertFailed, fmt.Sprintf("The %s amount has to be less than 5 decimals like 0.00001", i.FromNetwork))
		log.Println(err)
		return ctx.JSON(http.StatusOK, &Response{
			Result:  err.Error(),
			Success: false,
		})
	}

	var fee float64
	if i.ToNetwork == "bsc" {
		fee = float64(util.UserBscTransactionFee)
	}
	decimalAmount := util.GetTransactionAmount(i.ToNetwork, i.Amount)
	log.Println("user transaction destination network", i.ToNetwork, "total amount:", decimalAmount)
	decimalFee := util.GetTransactionAmount(i.ToNetwork, fee)
	log.Println("user transaction destination network", i.ToNetwork, "fee amount:", decimalFee)
	if decimalFee.GreaterThanOrEqual(decimalAmount) {
		err := fmt.Errorf("Transaction fee %v is bigger than %v", decimalFee, i.Amount)
		return ctx.JSON(http.StatusOK, &Response{
			Result:  err.Error(),
			Success: false,
		})
	}

	log.Println("Checking if POST request is old request failed or in the process")
	filter := map[string]string{
		"from_address": i.FromAddress,
		"completed":    "false",
	}
	result, err := cctx.MongoDB.FindTx(filter)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			log.Println(err)
			return ctx.JSON(http.StatusOK, &Response{
				Result:  err.Error(),
				Success: false,
			})
		}
	}
	if result != nil {
		log.Println("Found old POST request")
		tx, _ := result.(model.Tx) // if the previous request from same address has failed(still completed is false),
		now := time.Now()          // new request more than 3 minutes later will update the previous request completed field to true
		var startedTime string
		if tx.CreatedAt == tx.UpdatedAt {
			startedTime = tx.CreatedAt
		} else {
			startedTime = now.Format(util.TimeFormat)
		}
		log.Println("POST request 3 minutes timeout started. started time:", startedTime)
		tx.UpdatedAt = now.Format(util.TimeFormat)
		_, err := cctx.MongoDB.UpdateTx(&tx)
		if err != nil {
			log.Println(err)
		}
		createdAtTime, _ := time.Parse(util.TimeFormat, tx.CreatedAt)
		expiratedTime := createdAtTime.Add(time.Duration(util.TimeoutMinute * time.Minute))
		log.Println("expirated time for 3 minutes timeout:", expiratedTime, "remained time:", expiratedTime.Sub(now).Minutes(), "minutes")
		if expiratedTime.Sub(now).Minutes() > 0 {
			err = errors.PreparePayload(errors.ECTxInsertFailed, fmt.Sprintf("Former transaction is not complete yet. Please wait for %.2f minutes more", expiratedTime.Sub(now).Minutes()))
			return ctx.JSON(http.StatusOK, &Response{
				Result:  err.Error(),
				Success: false,
			})
		}
		log.Println("3 minutes timeout passed")
		tx.Completed = true
		_, err = cctx.MongoDB.UpdateTx(&tx)
		if err != nil {
			log.Println(err)
		}
		log.Printf("Update old result %+v completed status to true. Because 3 minute timeout passed", result)
	}

	log.Printf("Preparing requeset insertion into DB")
	tx := model.NewTx()
	tx.FromAddress = i.FromAddress
	tx.ToAddress = i.ToAddress
	tx.SourceNetwork = i.FromNetwork
	tx.DestinationNetwork = i.ToNetwork
	tx.Amount = i.Amount
	tx.CreatedAt = time.Now().Format(util.TimeFormat)
	tx.UpdatedAt = tx.CreatedAt

	go func() {
		if tx.SourceNetwork == "bsc" {
			cctx.BscConn.InputChan <- tx
		}
	}()

	_, err = cctx.MongoDB.InsertTx(tx)
	if err != nil {
		err := errors.PreparePayload(errors.ECTxInsertFailed, err)
		log.Println(err)
		return ctx.JSON(http.StatusOK, &Response{
			Result:  err.Error(),
			Success: false,
		})
	}
	log.Printf("====== POST request is inserted in DB: %+v ======\n ", *tx)
	return ctx.JSON(http.StatusOK, &Response{
		Result:  tx,
		Success: true,
	})
}

func addBlacklist(ctx echo.Context) error {
	log.Println("====== Get POST blacklist request ======")
	cctx := ctx.Get("cctx").(*httpserver.CCtx)

	br := new(BlacklistRequest)
	if err := ctx.Bind(br); err != nil {
		err = errors.Prepare(errors.ECBindError, err)
		log.Printf("====== %s ======\n", err.Error())
		return ctx.JSON(http.StatusOK, &Response{
			Result:  err.Error(),
			Success: false,
		})
	}
	if err := br.Validate(); err != nil {
		err = fmt.Errorf("The address is not valid")
		log.Printf("====== %s ======\n", err.Error())
		return ctx.JSON(http.StatusOK, &Response{
			Result:  err.Error(),
			Success: false,
		})
	}
	log.Printf("POST request: %+v\n", *br)

	log.Println("Checking if address is already in blacklist")
	filter := map[string]string{
		"address": br.Address,
	}
	result, err := cctx.MongoDB.FindTx(filter)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			log.Println(err)
			return ctx.JSON(http.StatusOK, &Response{
				Result:  err.Error(),
				Success: false,
			})
		}
	}
	if result != nil {
		err = errors.PreparePayload(errors.ECTxInsertFailed, "The address is already in blacklist")
		log.Printf("====== %s ======\n", err.Error())
		return ctx.JSON(http.StatusOK, &Response{
			Result:  err.Error(),
			Success: false,
		})
	}

	log.Printf("Preparing blacklist requeset insertion into DB")

	blacklist := model.NewBlacklist()
	blacklist.Address = br.Address
	blacklist.CreatedAt = time.Now()
	blacklist.UpdatedAt = blacklist.CreatedAt

	_, err = cctx.MongoDB.InsertBlacklist(blacklist)
	if err != nil {
		err := errors.PreparePayload(errors.ECTxInsertFailed, err)
		log.Println(err)
		return ctx.JSON(http.StatusOK, &Response{
			Result:  err.Error(),
			Success: false,
		})
	}
	log.Printf("====== POST blacklist request is inserted in DB: %+v ======\n ", *br)
	return ctx.JSON(http.StatusOK, &Response{
		Result:  br,
		Success: true,
	})
}

func getTx(ctx echo.Context) error {
	log.Println("====== Get GET transaction request ======")
	cctx := ctx.Get("cctx").(*httpserver.CCtx)

	id := ctx.Param("id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		err := errors.PreparePayload(errors.ECInvalidID, err)
		log.Println(err)
		log.Printf("====== Failed to convert id to objectid: %+v ======\n", id)
		return ctx.JSON(http.StatusOK, &Response{
			Result:  err.Error(),
			Success: false,
		})
	}
	log.Println("id", id, "objectId", objectId)
	result, err := cctx.MongoDB.GetTx(objectId)
	if err != nil {
		err := errors.PreparePayload(errors.ECTxSelectFailed, err)
		log.Println(err)
		log.Printf("====== No data in DB with id %+v: %+v ======\n", id, result)
		return ctx.JSON(http.StatusOK, &Response{
			Result:  err.Error(),
			Success: false,
		})
	}
	log.Printf("Get transaction data by id from DB: %+v\n", result)
	tx, ok := result.(model.Tx)
	if !ok {
		err := fmt.Errorf("Getting db tx result is failed")
		log.Println(err)
		log.Printf("====== Failed to convert transaction data %+v to struct model.Tx: %+v ======\n", result, tx)
		return ctx.JSON(http.StatusOK, &Response{
			Result:  err.Error(),
			Success: false,
		})
	}
	log.Printf("transaction conversion is successful: %+v\n", tx)
	errMsgLogFormat := "====== %s ======"

	if !tx.Completed {
		errMsg := "Knstl transaction is on the process"
		log.Printf(errMsgLogFormat, errMsg)
		return ctx.JSON(http.StatusOK, &Response{
			Result:  errMsg,
			Success: false,
		})
	}

	log.Printf("====== Found complete transaction: %+v ======\n", tx)
	return ctx.JSON(http.StatusOK, &Response{
		Result:  tx,
		Success: true,
	})
}

func getLog(ctx echo.Context) error {
	logFile, err := os.ReadFile(logger.LogFileName)
	if err != nil {
		log.Println("Cannot read log file", err)
	}
	logContent := string(logFile)

	return ctx.HTML(http.StatusOK, strings.Replace(logContent, "\n", "<br>", -1))
}

func numDecPlaces(v float64) int {
	s := strconv.FormatFloat(v, 'f', -1, 64)
	i := strings.IndexByte(s, '.')
	if i > -1 {
		return len(s) - i - 1
	}
	return 0
}
