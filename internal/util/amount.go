package util

import (
	"log"

	"github.com/shopspring/decimal"
)

const (
	UserBscTransactionFee   = 2
	UserKnstlTransactionFee = 0.0001

	AmountBscUnit                 = "1000000000000000000"
	AmountKnstlUnit               = "1000000"
	BlacklistAllowThresholdAmount = "1000000"
)

func GetTransactionAmount(network string, amount float64) decimal.Decimal {
	var amountUnit string
	if network == "bsc" {
		amountUnit = AmountBscUnit
	} else if network == "knstl" {
		amountUnit = AmountKnstlUnit
	}
	unitDecimalAmount, _ := decimal.NewFromString(amountUnit)
	log.Println(network, "unit decimal amount:", unitDecimalAmount)
	userTransactionDecimalFee := decimal.NewFromFloat(amount)
	log.Println("User transaction decimal amount:", userTransactionDecimalFee)
	userTransactionTotalFee := userTransactionDecimalFee.Mul(unitDecimalAmount)
	log.Println("User transaction total", network, "amount:", userTransactionTotalFee)
	return userTransactionTotalFee
}
