package payment_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/swagftw/stripe_pay_service/pkg/payments"
	"github.com/swagftw/stripe_pay_service/transaction"
	"github.com/swagftw/stripe_pay_service/types"
	"github.com/swagftw/stripe_pay_service/utl/config"
	"github.com/swagftw/stripe_pay_service/utl/logger"
	"github.com/swagftw/stripe_pay_service/utl/mock"
	"github.com/swagftw/stripe_pay_service/utl/stripeclient"
)

func TestCreatePaymentIntent(t *testing.T) {
	logger.InitLogger()

	err := config.InitConfig("../../../utl/config/config.local.yaml", "../../../.env")
	if err != nil {
		t.Fatal(err)
	}

	cases := []struct {
		name          string
		data          *types.CreateIntentReq
		wantData      *types.CreateIntentRes
		wantErr       bool
		tx            transaction.Transaction
		repo          payments.Repository
		stripeService stripeclient.StripeService
	}{
		{
			name:    "success",
			wantErr: false,
			data: &types.CreateIntentReq{
				Amount:      100,
				Email:       "asd@y.com",
				Phone:       "1234567890",
				Description: "test",
			},
			tx:            mock.NewTxMock(),
			stripeService: stripeclient.New(),
			repo: mock.PaymentMockRepository{
				CreatePaymentFn: func(ctx context.Context, payment *payments.PaymentIntent) error {
					return nil
				},
			},
		},
		{
			name:          "invalid amount",
			data:          &types.CreateIntentReq{Amount: 0},
			wantErr:       true,
			tx:            mock.NewTxMock(),
			stripeService: stripeclient.New(),
			repo: mock.PaymentMockRepository{
				CreatePaymentFn: func(ctx context.Context, payment *payments.PaymentIntent) error {
					return nil
				},
			},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			payS := payments.NewService(tt.tx, tt.repo, tt.stripeService)
			_, err := payS.CreatePaymentIntent(context.TODO(), tt.data)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestCapturePaymentIntent(t *testing.T) {
	logger.InitLogger()

	err := config.InitConfig("../../../utl/config/config.local.yaml", "../../../.env")
	if err != nil {
		t.Fatal(err)
	}

	cases := []struct {
		name          string
		id            string
		wantErr       bool
		tx            transaction.Transaction
		repo          payments.Repository
		stripeService stripeclient.StripeService
	}{
		{
			name:          "success",
			id:            "123",
			wantErr:       false,
			stripeService: stripeclient.NewMock(),
			tx:            mock.NewTxMock(),
			repo: mock.PaymentMockRepository{
				GetPaymentFn: func(ctx context.Context, id string) (*payments.PaymentIntent, error) {
					return &payments.PaymentIntent{
						ID: "",
					}, nil
				},
				UpdatePaymentFn: func(ctx context.Context, payment *payments.PaymentIntent) error {
					return nil
				},
			},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			payS := payments.NewService(tt.tx, tt.repo, tt.stripeService)
			_, err := payS.CapturePaymentIntent(context.TODO(), tt.id)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestGetPaymentIntents(t *testing.T) {
	logger.InitLogger()

	err := config.InitConfig("../../../utl/config/config.local.yaml", "../../../.env")
	if err != nil {
		t.Fatal(err)
	}

	cases := []struct {
		name          string
		stripeService stripeclient.StripeService
		wantErr       bool
	}{
		{
			name:          "success",
			stripeService: stripeclient.NewMock(),
			wantErr:       false,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			payS := payments.NewService(mock.NewTxMock(), mock.PaymentMockRepository{}, tt.stripeService)
			_, err := payS.GetPaymentIntents(context.TODO())
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestCreateRefund(t *testing.T) {
	logger.InitLogger()

	err := config.InitConfig("../../../utl/config/config.local.yaml", "../../../.env")
	if err != nil {
		t.Fatal(err)
	}

	cases := []struct {
		name          string
		paymentID     string
		wantErr       bool
		stripeService stripeclient.StripeService
		tx            transaction.Transaction
		repo          payments.Repository
	}{
		{
			name:          "success",
			paymentID:     "123",
			wantErr:       false,
			stripeService: stripeclient.NewMock(),
			tx:            mock.NewTxMock(),
			repo: mock.PaymentMockRepository{
				UpdatePaymentFn: func(ctx context.Context, payment *payments.PaymentIntent) error {
					return nil
				},
				CreateRefundFn: func(ctx context.Context, refund *payments.Refund) (*payments.Refund, error) {
					return &payments.Refund{}, nil
				},
				GetPaymentFn: func(ctx context.Context, id string) (*payments.PaymentIntent, error) {
					return &payments.PaymentIntent{}, nil
				},
			},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			payS := payments.NewService(tt.tx, tt.repo, tt.stripeService)
			_, err := payS.CreateRefund(context.TODO(), tt.paymentID)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}
