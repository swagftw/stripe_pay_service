package postgres

import (
	"context"

	"gorm.io/gorm"

	"github.com/swagftw/stripe_pay_service/transaction"
	"github.com/swagftw/stripe_pay_service/utl/constant"
)

type txn struct {
	db *gorm.DB
}

// Run runs the given function in a postgres transaction.
func (t *txn) Run(ctx context.Context, fn func(ctx context.Context) error) error {
	db := t.db

	// check if the transaction is already in progress
	if ctx.Value(constant.PostgresTxKey) != nil {
		if postgresTx, ok := ctx.Value(constant.PostgresTxKey).(*gorm.DB); ok {
			db = postgresTx
		}
	}

	// return new transaction from the given db (which may have another transaction in progress)
	return db.Transaction(func(tx *gorm.DB) error {
		ctx = context.WithValue(ctx, constant.TxKey(constant.PostgresTxKey), tx)

		return fn(ctx)
	})
}

func NewPostgresTx(db *gorm.DB) transaction.Transaction {
	return &txn{db: db}
}
