package routes

import (
	"fmt"
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation"
)

type TxInput struct {
	FromAddress string  `json:"from_address"`
	FromNetwork string  `json:"from_network"`
	ToAddress   string  `json:"to_address"`
	ToNetwork   string  `json:"to_network"`
	Amount      float64 `json:"amount"`
}

// Validate struct
func (i TxInput) Validate() error {
	bscRe := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	knstlRe := regexp.MustCompile("^darc1[0-9a-zA-Z]{38}$")
	if i.FromNetwork == "bsc" && len(i.FromAddress) > 0 && !bscRe.MatchString(i.FromAddress) {
		return fmt.Errorf("Not valid bsc address: %s", i.FromAddress)
	} else if i.ToNetwork == "bsc" && len(i.ToAddress) > 0 && !bscRe.MatchString(i.ToAddress) {
		return fmt.Errorf("Not valid bsc address: %s", i.ToAddress)
	} else if i.FromNetwork == "knstl" && len(i.FromAddress) > 0 && !knstlRe.MatchString(i.FromAddress) {
		return fmt.Errorf("Not valid knstl address: %s", i.FromAddress)
	} else if i.ToNetwork == "knstl" && len(i.ToAddress) > 0 && !knstlRe.MatchString(i.ToAddress) {
		return fmt.Errorf("Not valid knstl address: %s", i.ToAddress)
	}

	return validation.ValidateStruct(&i,
		validation.Field(
			&i.FromAddress,
			validation.Required,
		),
		validation.Field(
			&i.FromNetwork,
			validation.Required,
		),
		validation.Field(
			&i.ToAddress,
			validation.Required,
		),
		validation.Field(
			&i.ToNetwork,
			validation.Required,
		),
		validation.Field(
			&i.Amount,
			validation.Required,
		),
	)
}

type Response struct {
	//Result  *model.Tx `json:"result"`
	Result  interface{} `json:"result"`
	Success bool        `json:"success"`
}

type BlacklistRequest struct {
	Address string `json:"address"`
}

// Validate struct
func (br BlacklistRequest) Validate() error {
	bscRe := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	knstlRe := regexp.MustCompile("^darc1[0-9a-zA-Z]{38}$")
	if !bscRe.MatchString(br.Address) && !knstlRe.MatchString(br.Address) {
		return fmt.Errorf("Not valid address: %s", br.Address)
	}
	return nil
}
