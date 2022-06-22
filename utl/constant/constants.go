package constant

const PostgresTxKey string = "postgres.tx"
const RequestIDKey string = "request.id"

type TxKey string

const TryAgainLater = "try again later"

func StringToPtr(val string) *string {
	return &val
}
