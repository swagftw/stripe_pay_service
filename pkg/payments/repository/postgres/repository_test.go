package postgres

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/swagftw/stripe_pay_service/pkg/payments"
	"github.com/swagftw/stripe_pay_service/utl/config"
	"github.com/swagftw/stripe_pay_service/utl/constant"
	"github.com/swagftw/stripe_pay_service/utl/logger"
	"github.com/swagftw/stripe_pay_service/utl/storage"
)

func TestCreatePayment(t *testing.T) {
	logger.InitLogger()

	err := config.InitConfig("../../../../utl/config/config.local.yaml", "../../../../.env")
	if err != nil {
		t.Fatal(err)
	}

	db, err := storage.NewPostgresDB()
	if err != nil {
		t.Error(err)
	}

	cases := []struct {
		name    string
		wantErr bool
		data    *payments.PaymentIntent
	}{
		{
			name:    "success",
			wantErr: false,
			data: &payments.PaymentIntent{
				ID:         "",
				ProviderID: "pi_test",
				Email:      "asd@gmail.com",
				Status:     "succeeded",
				Payload:    "null",
			},
		},
		{
			name:    "invalid payload",
			wantErr: true,
			data: &payments.PaymentIntent{
				ID:         "",
				ProviderID: "pi_test",
				Email:      "",
				Status:     "succeeded",
			},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			repo := NewPaymentsRepo(db)
			err = repo.CreatePayment(context.TODO(), tt.data)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestGetPayment(t *testing.T) {
	logger.InitLogger()

	err := config.InitConfig("../../../../utl/config/config.local.yaml", "../../../../.env")
	if err != nil {
		t.Fatal(err)
	}

	db, err := storage.NewPostgresDB()
	if err != nil {
		t.Error(err)
	}

	cases := []struct {
		name       string
		wantErr    bool
		providerId string
	}{
		{
			name:    "success",
			wantErr: false,
		},
		{
			name:       "invalid id",
			wantErr:    true,
			providerId: "123",
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			repo := NewPaymentsRepo(db)

			if tt.name == "success" {
				intent := &payments.PaymentIntent{
					ID:         "",
					ProviderID: "pi_test",
					Email:      "asd@asd.com",
					Status:     "succeeded",
					Payload:    "null",
					Amount:     100,
				}

				err = repo.CreatePayment(context.TODO(), intent)
				if err != nil {
					t.Error(err)
				}

				intent, err = repo.GetPayment(context.TODO(), intent.ProviderID)
				assert.Equal(t, tt.wantErr, err != nil)
				assert.Equal(t, tt.providerId, intent.ProviderID)
				return
			}

			_, err = repo.GetPayment(context.TODO(), tt.providerId)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestUpdatePayment(t *testing.T) {
	logger.InitLogger()

	err := config.InitConfig("../../../../utl/config/config.local.yaml", "../../../../.env")
	if err != nil {
		t.Fatal(err)
	}

	db, err := storage.NewPostgresDB()
	if err != nil {
		t.Error(err)
	}

	cases := []struct {
		name    string
		wantErr bool
		data    *payments.PaymentIntent
	}{
		{
			name:    "success",
			wantErr: false,
		},
		{
			name:    "invalid id",
			wantErr: false,
			data: &payments.PaymentIntent{
				ID: "asd",
			},
		}, {
			name:    "invalid payload",
			wantErr: true,
			data: &payments.PaymentIntent{
				ID: "",
			},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			payS := NewPaymentsRepo(db)
			intent := &payments.PaymentIntent{
				Amount:     100,
				ProviderID: "pi_test",
				Email:      "asd@asd.com",
				Status:     "new",
				Payload:    "null",
			}

			err = payS.CreatePayment(context.TODO(), intent)
			intent.Status = "succeeded"
			if tt.data != nil {
				err := payS.UpdatePayment(context.TODO(), tt.data)
				assert.Equal(t, tt.data.ProviderID, "")
				assert.Equal(t, tt.wantErr, err != nil)
				return
			}
			err := payS.UpdatePayment(context.TODO(), intent)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestCreateRefund(t *testing.T) {
	logger.InitLogger()

	err := config.InitConfig("../../../../utl/config/config.local.yaml", "../../../../.env")
	if err != nil {
		t.Fatal(err)
	}

	db, err := storage.NewPostgresDB()
	if err != nil {
		t.Error(err)
	}

	cases := []struct {
		name    string
		wantErr bool
		data    *payments.Refund
	}{
		{
			name:    "success",
			wantErr: false,
			data: &payments.Refund{
				ProviderID:      constant.StringToPtr("rf_test"),
				PaymentIntentID: constant.StringToPtr("pi_test"),
				Amount:          100,
				Status:          constant.StringToPtr("succeeded"),
			},
		},
		{
			name:    "invalid data",
			wantErr: true,
			data: &payments.Refund{
				ProviderID:      nil,
				PaymentIntentID: constant.StringToPtr("pi_test"),
				Amount:          100,
				Status:          constant.StringToPtr("succeeded"),
			},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			payS := NewPaymentsRepo(db)
			_, err = payS.CreateRefund(context.TODO(), tt.data)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}
