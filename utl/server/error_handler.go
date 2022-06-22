package server

import (
	"context"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"

	"github.com/swagftw/stripe_pay_service/types"
	"github.com/swagftw/stripe_pay_service/utl/constant"
	"github.com/swagftw/stripe_pay_service/utl/fault"
	"github.com/swagftw/stripe_pay_service/utl/logger"
)

func ErrorHandler(err error, ctx echo.Context) {
	var errResp fault.ErrResponse
	errResp.StatusCode = http.StatusInternalServerError
	errResp.Message = "Internal Server Error"
	errResp.Err = err.Error()
	errResp.Res = constant.TryAgainLater

	switch e := err.(type) {
	case *fault.HTTPError:
		errResp.StatusCode = e.Status
		errResp.Message = e.Message
		errResp.Res = e.Res
		errResp.Service = e.Service

		if e.Err != nil {
			errResp.Err = e.Err.Error()
		}
	case *echo.HTTPError:
		errResp.StatusCode = e.Code
		errResp.Message = e.Message
		errResp.Res = constant.TryAgainLater
		errResp.Err = e.Error()
	case *validator.ValidationErrors:
		errResp.StatusCode = http.StatusBadRequest
		errResp.Message = "Bad Request"
		errResp.Err = e.Error()
		errResp.Res = "Validation Error"
	case types.CopyError:
		errResp.StatusCode = http.StatusInternalServerError
		errResp.Message = "Internal Server Error"
		errResp.Err = e.Error()
		errResp.Res = constant.TryAgainLater
	}

	if !ctx.Response().Committed {
		err := ctx.JSON(errResp.StatusCode, map[string]interface{}{
			"error": errResp,
		})

		logger.Logger.Error(context.TODO(), "error sending response from echo error handler", err)
	}
}
