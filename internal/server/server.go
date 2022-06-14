package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"runtime"
	"time"

	"github.com/konstellation/swap/internal/config"
	"github.com/konstellation/swap/internal/errors"
	"github.com/konstellation/swap/internal/httpserver"
	"github.com/konstellation/swap/internal/logger"
	"github.com/konstellation/swap/internal/mongo"
	"github.com/konstellation/swap/internal/routes"
	"github.com/konstellation/swap/internal/util/echopprof"
	"github.com/labstack/echo/v4/middleware"
	"github.com/neko-neko/echo-logrus/v2/log"
	"github.com/sirupsen/logrus"

	"github.com/konstellation/swap/internal/chain"
	"github.com/labstack/echo/v4"
)

// @Summary Healthcheck endpoint
// @Tags Main
// @Description Make sure that services working
// @Produce json
// @Router / [get]
func healthCheck(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"name":    os.Getenv("APP_NAME"),
		"success": true,
	})
}

func appRecover() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// defer the protection function
			defer func() {
				if r := recover(); r != nil {
					err, ok := r.(error)
					if !ok { // Create an error if we don't have one
						err = fmt.Errorf("%v", r)
					}
					stackSize := 4 << 10 // 4KB
					stack := make([]byte, stackSize)
					length := runtime.Stack(stack, true)
					fmt.Printf("[PANIC RECOVER] err: %v\n%s\n", err, stack[:length])
					c.Error(err)
				}
			}()
			return next(c)
		}
	}
}

// Run server
func (s *Server) Run() {
	go s.ListenChannelErrors()
	// close connections etc.
	defer s.MongoDB.Disconnect(context.Background())

	s.Logger.Infof("Start server on port (%s)", s.Config.Port)
	if s.Config.TLSEnable {
		s.Logger.Fatal(s.Echo.StartTLS(":"+s.Config.Port, s.Config.TLSCertLocation, s.Config.TLSPrivKeyLocation))
	} else {
		s.Logger.Fatal(s.Echo.Start(":" + s.Config.Port))
	}
}

// ListenChannelErrors func
func (s *Server) ListenChannelErrors() {
	for {
		select {
		case err := <-s.ErrChan:
			logger.Err("server error: %v", err)
		}
	}
}

// Server struct
type Server struct {
	Echo      *echo.Echo
	Config    *config.Config
	Logger    *log.MyLogger
	MongoDB   *mongo.Connection
	knstlConn *chain.KnstlConnection
	bscConn   *chain.BSCConnection
	ErrChan   chan error
}

func New(c *config.Config, mg *mongo.Connection, bsc *chain.BSCConnection, knstl *chain.KnstlConnection) *Server {
	srv := &Server{
		Config:    c,
		ErrChan:   make(chan error, 5),
		MongoDB:   mg,
		bscConn:   bsc,
		knstlConn: knstl,
	}
	return srv.init()
}

func (s *Server) SetEcho(v *echo.Echo) *Server {
	s.Echo = v
	return s
}

func (s *Server) SetLogger(v *log.MyLogger) {
	s.Logger = v
}

func (s *Server) init() *Server {
	// connect to postgres

	e := echo.New()
	s.SetEcho(e)

	// grpc server
	// inject logger to echo context
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339,
	})
	e.Logger = log.Logger()
	s.SetLogger(log.Logger())

	// inject errors handler
	e.HTTPErrorHandler = errors.HTTPErrorHandler

	// middleware
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "request=${id} method=${method}, uri=${uri}, status=${status}, error=${error}\n",
	}))
	e.Use(appRecover())
	e.Use(middleware.CORS())
	e.Use(httpserver.AttachCCtx(s.Config, s.MongoDB, s.bscConn, s.knstlConn))

	// endpoints
	e.GET("/", healthCheck) // healthCheck route
	//e.Static("/", "static")

	echopprof.Wrap(e)

	routes.InitRoutes(e)

	return s
}
