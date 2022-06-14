package httpserver

import (
	"github.com/konstellation/swap/internal/chain"
	"github.com/konstellation/swap/internal/config"
	"github.com/konstellation/swap/internal/mongo"
	echo "github.com/labstack/echo/v4"
)

// AttachCCtx attaches context to each request that will have a
// pointer to mysql connection, datastore connection, and logger
// methods
func AttachCCtx(cfg *config.Config, mg *mongo.Connection, bsc *chain.BSCConnection, knstl *chain.KnstlConnection) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("cctx", NewCCtx(c, cfg, mg, bsc, knstl))
			return next(c)
		}
	}
}
