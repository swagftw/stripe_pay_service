package mock

import (
	"context"

	"github.com/swagftw/stripe_pay_service/pkg/payments"
)

type PaymentMockRepository struct {
	CreatePaymentFn func(ctx context.Context, payment *payments.PaymentIntent) error
	UpdatePaymentFn func(ctx context.Context, payment *payments.PaymentIntent) error
	GetPaymentFn    func(ctx context.Context, id string) (*payments.PaymentIntent, error)
	CreateRefundFn  func(ctx context.Context, refund *payments.Refund) (*payments.Refund, error)
}

func (p PaymentMockRepository) CreatePayment(ctx context.Context, payment *payments.PaymentIntent) error {
	return p.CreatePaymentFn(ctx, payment)
}

func (p PaymentMockRepository) UpdatePayment(ctx context.Context, payment *payments.PaymentIntent) error {
	return p.UpdatePaymentFn(ctx, payment)
}

func (p PaymentMockRepository) GetPayment(ctx context.Context, id string) (*payments.PaymentIntent, error) {
	return p.GetPaymentFn(ctx, id)
}

func (p PaymentMockRepository) CreateRefund(ctx context.Context, refund *payments.Refund) (*payments.Refund, error) {
	return p.CreateRefundFn(ctx, refund)
}
