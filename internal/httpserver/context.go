package httpserver

import (
	"github.com/konstellation/swap/internal/chain"
	"github.com/konstellation/swap/internal/config"
	"github.com/konstellation/swap/internal/mongo"
	"github.com/labstack/echo/v4"
)

// CCtx custom echo context
type CCtx struct {
	// public
	Ctx       echo.Context
	MongoDB   *mongo.Connection
	KnstlConn *chain.KnstlConnection
	BscConn   *chain.BSCConnection

	// private
	cfg    *config.Config
	cfgSet bool
}

// NewCCtx create new context instance
func NewCCtx(ectx echo.Context, cfg *config.Config, mg *mongo.Connection, bsc *chain.BSCConnection, knstl *chain.KnstlConnection) *CCtx {
	appCtx := &CCtx{
		Ctx:       ectx,
		MongoDB:   mg,
		KnstlConn: knstl,
		BscConn:   bsc,
	}

	appCtx.SetConfig(cfg)
	return appCtx
}

// SetConfig set config to context
func (cctx *CCtx) SetConfig(c *config.Config) {
	if cctx.cfgSet {
		panic("config already set")
	}
	cctx.cfg = c
	cctx.cfgSet = true
}

// GetConfig get config from context
func (cctx *CCtx) GetConfig() *config.Config {
	if !cctx.cfgSet {
		panic("must set config before using")
	}
	return cctx.cfg
}
