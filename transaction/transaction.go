package transaction

import "context"

// Transaction interface represents any db transaction
type Transaction interface {
	Run(ctx context.Context, fn func(ctx context.Context) error) error
}
