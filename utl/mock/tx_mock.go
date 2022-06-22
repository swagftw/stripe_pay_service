package mock

import (
	"context"

	"github.com/swagftw/stripe_pay_service/transaction"
)

type txMock struct{}

func (t txMock) Run(ctx context.Context, fn func(ctx context.Context) error) error {
	return fn(ctx)
}

func NewTxMock() transaction.Transaction {
	return &txMock{}
}
