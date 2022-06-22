package stripeclient_test_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/swagftw/stripe_pay_service/types"
	"github.com/swagftw/stripe_pay_service/utl/config"
	"github.com/swagftw/stripe_pay_service/utl/logger"
	"github.com/swagftw/stripe_pay_service/utl/stripeclient"
)

func TestGetAllPaymentIntents(t *testing.T) {
	logger.InitLogger()

	err := config.InitConfig("../../config/config.local.yaml", "../../../.env")
	if err != nil {
		panic(err)
	}

	cases := []struct {
		name          string
		wantErr       bool
		stripeService stripeclient.StripeService
	}{
		{
			name:          "success",
			wantErr:       false,
			stripeService: stripeclient.NewMock(),
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.stripeService.GetAllPaymentIntents()
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestCreatePaymentIntent(t *testing.T) {
	logger.InitLogger()

	err := config.InitConfig("../../config/config.local.yaml", "../../../.env")
	if err != nil {
		panic(err)
	}

	cases := []struct {
		name          string
		wantErr       bool
		stripeService stripeclient.StripeService
		data          *types.CreateIntentReq
	}{
		{
			name:          "success",
			wantErr:       false,
			stripeService: stripeclient.NewMock(),
			data: &types.CreateIntentReq{
				Amount:      100,
				Email:       "asd@asd.com",
				Phone:       "1231231231",
				Description: "test",
			},
		},
		{
			name: "invalid amount",
			data: &types.CreateIntentReq{
				Amount: 0,
			},
			wantErr:       true,
			stripeService: stripeclient.New(),
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.stripeService.CreatePaymentIntent(tt.data)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestCapturePaymentIntent(t *testing.T) {
	logger.InitLogger()

	err := config.InitConfig("../../config/config.local.yaml", "../../../.env")
	if err != nil {
		panic(err)
	}

	cases := []struct {
		name          string
		wantErr       bool
		paymentID     string
		amount        int
		stripeService stripeclient.StripeService
	}{
		{
			name:          "success",
			wantErr:       false,
			paymentID:     "pi_1GqXqXqXqXqXqXqXqXqXqXqXqXqXqXqX",
			amount:        100,
			stripeService: stripeclient.New(),
		},
		{
			name:          "invalid payment id",
			wantErr:       true,
			paymentID:     "",
			amount:        100,
			stripeService: stripeclient.NewMock(),
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "success" {
				intent, err := tt.stripeService.CreatePaymentIntent(&types.CreateIntentReq{
					Amount:      100,
					Email:       "asd@asd.com",
					Phone:       "",
					Description: "test",
				})
				if err != nil {
					panic(err)
				}

				tt.paymentID = intent.ID

				_, err = tt.stripeService.CapturePaymentIntent(tt.paymentID, tt.amount)
				assert.Equal(t, tt.wantErr, err != nil)

				return
			}
			_, err := stripeclient.New().CapturePaymentIntent(tt.paymentID, tt.amount)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestCreateRefund(t *testing.T) {
	logger.InitLogger()

	err := config.InitConfig("../../config/config.local.yaml", "../../../.env")
	if err != nil {
		panic(err)
	}

	cases := []struct {
		name          string
		wantErr       bool
		paymentID     string
		amount        int
		stripeService stripeclient.StripeService
	}{
		{
			name:          "success",
			wantErr:       false,
			stripeService: stripeclient.New(),
		},
		{
			name:          "invalid payment id",
			wantErr:       true,
			paymentID:     "test",
			stripeService: stripeclient.New(),
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "success" {
				intent, err := tt.stripeService.CreatePaymentIntent(&types.CreateIntentReq{
					Amount:      100,
					Email:       "asd@asd.com",
					Phone:       "",
					Description: "test",
				})
				if err != nil {
					panic(err)
				}

				tt.paymentID = intent.ID
				tt.amount = intent.Amount

				_, err = tt.stripeService.CapturePaymentIntent(tt.paymentID, tt.amount)
				assert.Equal(t, tt.wantErr, err != nil)

				_, err = tt.stripeService.CreateRefund(tt.paymentID, tt.amount)
				assert.Equal(t, tt.wantErr, err != nil)

				return
			}

			_, err := tt.stripeService.CreateRefund(tt.paymentID, tt.amount)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}
