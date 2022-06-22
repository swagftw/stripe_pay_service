package postgres

import (
	"context"
	"net/http"

	"gorm.io/gorm"

	"github.com/swagftw/stripe_pay_service/pkg/payments"
	"github.com/swagftw/stripe_pay_service/utl/fault"
	"github.com/swagftw/stripe_pay_service/utl/storage"
)

type repository struct {
	db *gorm.DB
}

// CreatePayment creates a payment intent.
func (r repository) CreatePayment(ctx context.Context, payment *payments.PaymentIntent) error {
	db := storage.GetGormDBFromContext(ctx, r.db)
	err := db.Model(payment).Create(payment).Error

	return err
}

// GetPayment gets the payment intent.
func (r repository) GetPayment(ctx context.Context, id string) (*payments.PaymentIntent, error) {
	db := storage.GetGormDBFromContext(ctx, r.db)

	payment := new(payments.PaymentIntent)
	err := db.Where("provider_id = ?", id).First(payment).Error

	if err == gorm.ErrRecordNotFound {
		return nil, fault.New(http.StatusNotFound, "payment_repo", "payment intent not found", "provide valid intent id", "INVALID_PAYMENT_INTENT_ID", err)
	}

	return payment, err
}

// UpdatePayment updates the payment intent.
func (r repository) UpdatePayment(ctx context.Context, payment *payments.PaymentIntent) error {
	db := storage.GetGormDBFromContext(ctx, r.db)
	err := db.Updates(payment).Error

	return err
}

// CreateRefund creates a refund.
func (r repository) CreateRefund(ctx context.Context, refund *payments.Refund) (*payments.Refund, error) {
	db := storage.GetGormDBFromContext(ctx, r.db)
	err := db.Model(&payments.Refund{}).Create(refund).Error

	return refund, err
}

// NewPaymentsRepo returns a new payment repository.
func NewPaymentsRepo(db *gorm.DB) payments.Repository {
	return &repository{
		db: db,
	}
}
