package errors

import (
	errs "errors"
	"fmt"
	"net/http"

	ozzo "github.com/go-ozzo/ozzo-validation"
	"github.com/labstack/echo/v4"
)

// AppError standard client error struct
type AppError struct {
	HTTPCode         int          `json:"http_code"`
	Code             int          `json:"error_code"`
	Message          string       `json:"message"`
	Payload          interface{}  `json:"payload,omitempty"`
	ValidationErrors *ozzo.Errors `json:"validation_errors,omitempty"`
}

func (e AppError) String() string {
	if e.Message == "" {
		e.Message = "undefined internal error"
	}
	return fmt.Sprintf("%d %s %v", e.Code, e.Message, e.Payload)
}

func (e AppError) Error() string {
	return e.String()
}

func New(text string) error {
	return errs.New(text)
}

// NewAppError create error object
func NewAppError(httpCode, code int, err error) *AppError {
	appError := &AppError{
		HTTPCode: httpCode,
		Code:     code,
		Message:  "unknown internal server error",
	}

	if err != nil {
		if vErrs, ok := err.(ozzo.Errors); ok {
			appError.Message = "validation errors"
			appError.ValidationErrors = &vErrs
		} else {
			appError.Message = err.Error()
		}
	}

	return appError
}

// NewAppError create error object
func NewAppErrorPayload(httpCode, code int, err error, payload interface{}) *AppError {
	appError := NewAppError(httpCode, code, err)
	if payload != nil {
		appError.Payload = payload
	}

	return appError
}

func getHttpCode(code int) int {
	httpCode := http.StatusInternalServerError
	if code > 200 && code < http.StatusNetworkAuthenticationRequired { // last 511 http error code
		httpCode = code
	}

	// trying with codes map
	if val, ok := CodeMap[code]; ok {
		httpCode = val
	}

	return httpCode
}

func getTextCode(code int) string {
	err := "unspecified error"
	if text, ok := CodeText[code]; ok {
		err = text
	}

	return err
}

// Prepare client error object by error code
func Prepare(code int, err error) *AppError {
	return NewAppError(getHttpCode(code), code, err)
}

// Prepare client error object by error code and payload [provider id, date, etc...]
func PreparePayload(code int, payload interface{}) *AppError {
	if err, ok := payload.(error); ok {
		payload = err.Error()
	}

	return NewAppErrorPayload(getHttpCode(code), code, fmt.Errorf(getTextCode(code)), payload)
}

// Prepare client error object by error code and payload with custom error text
func PreparePayloadCustom(code int, payload interface{}, err error) *AppError {
	return NewAppErrorPayload(getHttpCode(code), code,
		fmt.Errorf("%s %v", getTextCode(code), err), payload)
}

// Response http with client error object
func (e *AppError) Response(ctx echo.Context) error {
	return ctx.JSON(e.HTTPCode, e)
}
