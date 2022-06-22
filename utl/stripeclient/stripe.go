package stripeclient

import (
	"context"
	"net/http"
	"time"

	"github.com/jinzhu/copier"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/client"

	"github.com/swagftw/stripe_pay_service/types"
	"github.com/swagftw/stripe_pay_service/utl/config"
	"github.com/swagftw/stripe_pay_service/utl/fault"
	"github.com/swagftw/stripe_pay_service/utl/logger"
)

type stripeClient struct {
	client *client.API
}

type StripeService interface {
	CreatePaymentIntent(req *types.CreateIntentReq) (*types.CreateIntentRes, error)
	CapturePaymentIntent(paymentID string, amount int) (*types.CaptureIntentRes, error)
	GetAllPaymentIntents() ([]*types.PaymentIntent, error)
	CreateRefund(paymentID string, amount int) (*types.CreateRefundRes, error)
}

func New() StripeService {
	stripeCfg := config.GetGlobalConfig().GetStripeConfig()

	return &stripeClient{
		client: client.New(stripeCfg.SecretKey, nil),
	}
}

func NewMock() StripeService {
	httpClient := &http.Client{
		Timeout: 10 * time.Second,
	}
	backend := stripe.GetBackendWithConfig(stripe.APIBackend, &stripe.BackendConfig{
		HTTPClient: httpClient,
		URL:        stripe.String("http://localhost:12111"),
	})

	return &stripeClient{
		client: client.New("sk_test_123", &stripe.Backends{API: backend}),
	}
}

// CreatePaymentIntent creates payment intent on stripe and sends back the response.
func (sc *stripeClient) CreatePaymentIntent(req *types.CreateIntentReq) (*types.CreateIntentRes, error) {
	intent := &stripe.PaymentIntentParams{
		Amount:       &req.Amount,
		Currency:     stripe.String(string(stripe.CurrencyINR)),
		Description:  &req.Description,
		ReceiptEmail: &req.Email,
		PaymentMethodTypes: []*string{
			stripe.String("card"),
		},
		CaptureMethod: stripe.String("manual"),
	}

	stripeIntent, err := sc.client.PaymentIntents.New(intent)
	if err != nil {
		msg := "source:stripe, message:error creating payment intent"
		logger.Logger.Error(context.TODO(), msg, err)

		if stripeErr, ok := err.(*stripe.Error); ok {
			if stripeErr.Code == stripe.ErrorCodePaymentIntentInvalidParameter {
				return nil, fault.New(http.StatusBadRequest, "stripeclient", "error creating payment intent", "check the request params", "ERR_INVALID_PARAMS", err)
			} else if stripeErr.Code == stripe.ErrorCodeAmountTooSmall {
				return nil, fault.New(http.StatusBadRequest, "stripeclient", "error creating payment intent", "amount too small", "ERR_AMOUNT_TOO_SMALL", err)
			}
		}

		return nil, err
	}

	res := new(types.CreateIntentRes)

	err = copier.Copy(res, stripeIntent)
	if err != nil {
		msg := "source:copier, message: error copying payment intent"
		logger.Logger.Error(context.TODO(), msg, err)

		return nil, fault.New(http.StatusInternalServerError, "stripeclient", "error creating copying data", "something went wrong", "ERR_INTERNAL_SERVER_ERROR", err)
	}

	return res, nil
}

// CapturePaymentIntent captures a payment intent.
func (sc *stripeClient) CapturePaymentIntent(paymentID string, amount int) (*types.CaptureIntentRes, error) {
	// update the payment intent with payment method first
	_, err := sc.client.PaymentIntents.Confirm(paymentID, &stripe.PaymentIntentConfirmParams{
		PaymentMethod: stripe.String("pm_card_visa"),
	})
	if err != nil {
		msg := "source:stripe, message:error updating payment intent"
		logger.Logger.Error(context.TODO(), msg, err)
		if stripeErr, ok := err.(*stripe.Error); ok {
			if stripeErr.Code == stripe.ErrorCodePaymentIntentInvalidParameter {
				return nil, fault.New(http.StatusBadRequest, "stripeclient", "error updating payment intent", "check the request params", "ERR_INVALID_PARAMS", err)
			} else if stripeErr.Code == stripe.ErrorCodeChargeAlreadyCaptured {
				return nil, fault.New(http.StatusBadRequest, "stripeclient", "error updating payment intent", "payment intent already captured", "ERR_ALREADY_CAPTURED", err)
			} else if stripeErr.Code == stripe.ErrorCodePaymentIntentUnexpectedState {
				return nil, fault.New(http.StatusBadRequest, "stripeclient", "error updating payment intent", "payment intent unexpected state", "ERR_UNEXPECTED_STATE", err)
			}
		}

		return nil, err
	}

	// capture the payment intent using amount
	paymentIntent, err := sc.client.PaymentIntents.Capture(paymentID, &stripe.PaymentIntentCaptureParams{
		AmountToCapture: stripe.Int64(int64(amount)),
	})
	if err != nil {
		msg := "source:stripe, message:error capturing payment intent"
		logger.Logger.Error(context.TODO(), msg, err)

		return nil, err
	}

	res := new(types.CaptureIntentRes)

	err = copier.Copy(res, paymentIntent)
	if err != nil {
		msg := "source:copier, message: error copying payment intent"
		logger.Logger.Error(context.TODO(), msg, err)

		return nil, err
	}

	return res, nil
}

// GetAllPaymentIntents returns all payment intents.
func (sc *stripeClient) GetAllPaymentIntents() ([]*types.PaymentIntent, error) {
	resp := make([]*types.PaymentIntent, 0)

	params := &stripe.PaymentIntentListParams{}
	params.Filters.AddFilter("limit", "", "10")
	i := sc.client.PaymentIntents.List(params)

	for i.Next() {
		paymentIntent := new(types.PaymentIntent)
		intent := i.PaymentIntent()
		_ = copier.Copy(paymentIntent, intent)
		resp = append(resp, paymentIntent)
	}

	return resp, nil
}

// CreateRefund creates a refund on stripe and sends back the response.
func (sc *stripeClient) CreateRefund(paymentID string, amount int) (*types.CreateRefundRes, error) {
	refund, err := sc.client.Refunds.New(&stripe.RefundParams{
		Amount:        stripe.Int64(int64(amount)),
		PaymentIntent: stripe.String(paymentID),
	})

	if err != nil {
		msg := "source:stripe, message:error creating refund"
		logger.Logger.Error(context.TODO(), msg, err)

		if stripeErr, ok := err.(*stripe.Error); ok {
			if stripeErr.Code == stripe.ErrorCodeChargeAlreadyRefunded {
				return nil, fault.New(http.StatusBadRequest, "stripeclient", "error creating refund", "payment intent already refunded", "ERR_ALREADY_REFUNDED", err)
			}
		}

		return nil, err
	}

	resp := new(types.CreateRefundRes)

	err = copier.Copy(resp, refund)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
