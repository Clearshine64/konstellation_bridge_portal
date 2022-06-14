package errors

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

// HTTPErrorHandler is a default http errors handler
func HTTPErrorHandler(err error, ctx echo.Context) {
	// handle default echo error
	if echoErr, ok := err.(*echo.HTTPError); ok {
		Prepare(echoErr.Code, fmt.Errorf("%v", echoErr.Message)).Response(ctx)
		return
	}

	// handle AppError - release version
	if appError, ok := err.(*AppError); ok {
		appError.Response(ctx)
		return
	}

	// unhandled error
	Prepare(http.StatusInternalServerError, err).Response(ctx)
}
