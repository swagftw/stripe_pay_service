package payments

import (
	"context"

	"github.com/swagftw/stripe_pay_service/utl/storage"
)

type (
	// Repository is the interface for the payment repository.
	Repository interface {
		CreatePayment(ctx context.Context, payment *PaymentIntent) error
		UpdatePayment(ctx context.Context, payment *PaymentIntent) error
		GetPayment(ctx context.Context, id string) (*PaymentIntent, error)
		CreateRefund(ctx context.Context, refund *Refund) (*Refund, error)
	}

	// PaymentIntent is the db model for the payment intent.
	PaymentIntent struct {
		ID         string `gorm:"primaryKey;default:('pi_' || generate_uid(12));not null"`
		Amount     int
		ProviderID string
		Email      string
		Status     string
		Payload    string `gorm:"type:jsonb"`
		storage.GormBase
	}

	// Refund is the db model for the refund.
	Refund struct {
		ID              string  `gorm:"primaryKey; default:('rf_' || generate_uid(12))"`
		ProviderID      *string `gorm:"not null"`
		PaymentIntentID *string `gorm:"not null"`
		Amount          int     `gorm:"not null"`
		Status          *string `gorm:"not null"`
		storage.GormBase
	}
)

func (*PaymentIntent) TableName() string {
	return "payment.payment_intents"
}

func (*Refund) TableName() string {
	return "payment.refunds"
}
