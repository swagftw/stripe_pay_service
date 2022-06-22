package payments

import (
	"context"
	"encoding/json"

	"github.com/stripe/stripe-go/v72"

	"github.com/swagftw/stripe_pay_service/transaction"
	"github.com/swagftw/stripe_pay_service/types"
	"github.com/swagftw/stripe_pay_service/utl/stripeclient"
)

type service struct {
	tx            transaction.Transaction
	repo          Repository
	stripeService stripeclient.StripeService
}

// CreatePaymentIntent creates a payment intent.
func (s service) CreatePaymentIntent(ctx context.Context, intent *types.CreateIntentReq) (*types.CreateIntentRes, error) {
	stripeIntent, err := s.stripeService.CreatePaymentIntent(intent)
	if err != nil {
		return nil, err
	}

	payload, err := json.Marshal(stripeIntent)
	if err != nil {
		return nil, err
	}

	dbIntent := &PaymentIntent{
		Amount:     stripeIntent.Amount,
		ProviderID: stripeIntent.ID,
		Payload:    string(payload),
		Status:     stripeIntent.Status,
	}
	if val, ok := stripeIntent.ReceiptEmail.(string); ok {
		dbIntent.Email = val
	}

	err = s.repo.CreatePayment(ctx, dbIntent)

	if err != nil {
		return nil, err
	}

	return stripeIntent, err
}

// CapturePaymentIntent captures a payment intent.
func (s service) CapturePaymentIntent(ctx context.Context, paymentID string) (*types.CaptureIntentRes, error) {
	// get payment intent from db first
	intent, err := s.repo.GetPayment(ctx, paymentID)
	if err != nil {
		return nil, err
	}

	// capture the payment intent using amount
	capturedIntent, err := s.stripeService.CapturePaymentIntent(paymentID, intent.Amount)
	if err != nil {
		return nil, err
	}

	// update status
	intent.Status = capturedIntent.Status
	// update the payment intent in db
	err = s.repo.UpdatePayment(ctx, intent)
	if err != nil {
		return nil, err
	}

	return capturedIntent, err
}

func (s service) GetPaymentIntents(ctx context.Context) (*types.GetIntentsRes, error) {
	intents, err := s.stripeService.GetAllPaymentIntents()
	if err != nil {
		return nil, err
	}

	resp := new(types.GetIntentsRes)
	resp.Intents = intents

	return resp, nil
}

// CreateRefund initiates a refund for a payment intent.
func (s service) CreateRefund(ctx context.Context, id string) (*types.CreateRefundRes, error) {
	// get payment intent from db first
	intent, err := s.repo.GetPayment(ctx, id)
	if err != nil {
		return nil, err
	}

	// create refund
	refund, err := s.stripeService.CreateRefund(id, intent.Amount)
	if err != nil {
		return nil, err
	}

	err = s.tx.Run(ctx, func(ctx context.Context) error {
		// update status.
		// this actually does not refund immediately but for the sake of demo update the state to refunded here.
		if refund.Status == string(stripe.RefundStatusSucceeded) {
			intent.Status = "refunded"
		}

		// update the payment intent in db
		err = s.repo.UpdatePayment(ctx, intent)
		if err != nil {
			return err
		}

		// create refund entry
		refundEntry := &Refund{
			ProviderID:      &refund.ID,
			PaymentIntentID: &intent.ProviderID,
			Amount:          refund.Amount,
			Status:          &refund.Status,
		}

		_, err = s.repo.CreateRefund(ctx, refundEntry)

		return err
	})

	return refund, err
}

// NewService creates a new payments service.
func NewService(tx transaction.Transaction, repo Repository, stripeService stripeclient.StripeService) types.PaymentService {
	return &service{
		tx:            tx,
		repo:          repo,
		stripeService: stripeService,
	}
}
