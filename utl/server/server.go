package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/go-playground/validator/v10"

	"github.com/swagftw/stripe_pay_service/utl/config"
	"github.com/swagftw/stripe_pay_service/utl/constant"
	"github.com/swagftw/stripe_pay_service/utl/logger"
)

// NewServer creates a new echo http server.
func NewServer() *echo.Echo {
	e := echo.New()

	// using default middlewares
	e.Use(middleware.Logger(), middleware.Recover(), middleware.CORS())
	e.HTTPErrorHandler = ErrorHandler

	e.Validator = &CustomValidator{V: validator.New()}
	e.Binder = &CustomBinder{b: &echo.DefaultBinder{}}

	return e
}

// StartServer starts the echo http server.
func StartServer(e *echo.Echo) {
	cfg := config.GetGlobalConfig()

	// ping server to check if it is running
	e.GET("/api/v1/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	errChan := make(chan error)

	// run server in background.
	go func() {
		srv := &http.Server{
			Addr:         ":" + cfg.GetServerConfig().Port,
			ReadTimeout:  time.Duration(cfg.GetServerConfig().Timeout) * time.Second,
			WriteTimeout: time.Duration(cfg.GetServerConfig().Timeout) * time.Second,
			IdleTimeout:  time.Duration(cfg.GetServerConfig().Timeout) * time.Second,
		}

		// listen and serve
		err := e.StartServer(srv)
		if err != nil {
			errChan <- err
		}
	}()

	signalChan := make(chan os.Signal, 1)

	// listen for SIGINT or SIGTERM
	go func() {
		signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
		sign := <-signalChan
		logger.Logger.Info(context.Background(), fmt.Sprintf("Shutting down server: %v", sign))
	}()

	// waits for the error channel or signal channel to return and logs the error
	select {
	case err := <-errChan:
		logger.Logger.Error(context.Background(), "shutting down server", err)
	case sign := <-signalChan:
		logger.Logger.Info(context.Background(), fmt.Sprintf("Shutting down server: %v", sign))
	}
}

func ToGoContext(e echo.Context) context.Context {
	return context.WithValue(e.Request().Context(), constant.TxKey(constant.RequestIDKey), e.Request().Header.Get("X-Request-ID"))
}

// CustomValidator holds custom validator
type CustomValidator struct {
	V *validator.Validate
}

func (cv CustomValidator) Validate(i interface{}) error {
	return cv.V.Struct(i)
}

// CustomBinder struct
type CustomBinder struct {
	b echo.Binder
}

// Bind tries to bind request into ledger_payment, and if it does then validate it
func (cb *CustomBinder) Bind(i interface{}, c echo.Context) error {
	if err := cb.b.Bind(i, c); err != nil && err != echo.ErrUnsupportedMediaType {
		return err
	}

	return c.Validate(i)
}
