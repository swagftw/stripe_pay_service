package payments

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/swagftw/stripe_pay_service/types"
	"github.com/swagftw/stripe_pay_service/utl/server"
)

type HTTP struct {
	service types.PaymentService
}

// InitHTTPHandlers initializes HTTP handlers for payments service
func InitHTTPHandlers(service types.PaymentService, v1 *echo.Group) {
	handler := &HTTP{service: service}

	paymentGroup := v1.Group("/payments")

	paymentGroup.POST("/create_intent", handler.createPaymentIntent)

	paymentGroup.POST("/capture_intent/:id", handler.capturePaymentIntent)

	paymentGroup.GET("/get_intents", handler.getPaymentIntents)

	paymentGroup.POST("/create_refund/:id", handler.refundPaymentIntent)
}

func (h HTTP) createPaymentIntent(c echo.Context) error {
	req := new(types.CreateIntentReq)
	if err := c.Bind(req); err != nil {
		return err
	}

	res, err := h.service.CreatePaymentIntent(server.ToGoContext(c), req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, res)
}

func (h HTTP) capturePaymentIntent(c echo.Context) error {
	id := c.Param("id")

	res, err := h.service.CapturePaymentIntent(server.ToGoContext(c), id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}

func (h HTTP) getPaymentIntents(c echo.Context) error {
	resp, err := h.service.GetPaymentIntents(server.ToGoContext(c))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)
}

func (h HTTP) refundPaymentIntent(c echo.Context) error {
	paymentID := c.Param("id")

	refund, err := h.service.CreateRefund(server.ToGoContext(c), paymentID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, refund)
}
