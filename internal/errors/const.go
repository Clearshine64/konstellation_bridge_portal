package errors

import (
	"net/http"
)

type ErrorCode struct {
	Code     int
	HttpCode int
	Text     string
}

const (
	ECInternalServerError = 500
	ECValidation          = 10001
	ECBindError           = 10002

	ECTxInsertFailed = 11000
	ECTxSelectFailed = 11001

	ECInvalidID = 10105
)

var (
	// CodeMap of errors
	CodeMap = map[int]int{
		ECInternalServerError: http.StatusInternalServerError,
		ECValidation:          http.StatusBadRequest,
		ECBindError:           http.StatusBadRequest,
		ECInvalidID:           http.StatusBadRequest,
		ECTxInsertFailed:      http.StatusInternalServerError,
		ECTxSelectFailed:      http.StatusInternalServerError,
	}

	CodeText = map[int]string{
		ECInternalServerError: "internal server error",

		ECBindError:  "bind error",
		ECValidation: "validation error",

		ECInvalidID: "invalid id",

		ECTxInsertFailed: "failed to insert new tx",
		ECTxSelectFailed: "failed to select tx",
	}
)
