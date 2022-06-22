package types

import (
	"context"
	"errors"
)

var (
	ErrCreatingPaymentIntent = errors.New("error creating payment intent")
)

type (
	// PaymentService is the interface that wraps basic payment service methods.
	PaymentService interface {
		CreatePaymentIntent(ctx context.Context, intent *CreateIntentReq) (*CreateIntentRes, error)
		CapturePaymentIntent(ctx context.Context, id string) (*CaptureIntentRes, error)
		GetPaymentIntents(ctx context.Context) (*GetIntentsRes, error)
		CreateRefund(ctx context.Context, id string) (*CreateRefundRes, error)
	}

	PaymentIntent struct {
		ID               string `json:"id"`
		Object           string `json:"object"`
		Amount           int    `json:"amount"`
		AmountCapturable int    `json:"amount_capturable"`
		AmountDetails    struct {
			Tip struct {
			} `json:"tip"`
		} `json:"amount_details"`
		AmountReceived          int         `json:"amount_received"`
		Application             interface{} `json:"application"`
		ApplicationFeeAmount    interface{} `json:"application_fee_amount"`
		AutomaticPaymentMethods interface{} `json:"automatic_payment_methods"`
		CanceledAt              interface{} `json:"canceled_at"`
		CancellationReason      interface{} `json:"cancellation_reason"`
		CaptureMethod           string      `json:"capture_method"`
		Charges                 struct {
			Object  string        `json:"object"`
			Data    []interface{} `json:"data"`
			HasMore bool          `json:"has_more"`
			Url     string        `json:"url"`
		} `json:"charges"`
		ClientSecret       string      `json:"client_secret"`
		ConfirmationMethod string      `json:"confirmation_method"`
		Created            int         `json:"created"`
		Currency           string      `json:"currency"`
		Customer           interface{} `json:"customer"`
		Description        interface{} `json:"description"`
		Invoice            interface{} `json:"invoice"`
		LastPaymentError   interface{} `json:"last_payment_error"`
		Livemode           bool        `json:"livemode"`
		Metadata           struct {
		} `json:"metadata"`
		NextAction           interface{} `json:"next_action"`
		OnBehalfOf           interface{} `json:"on_behalf_of"`
		PaymentMethod        interface{} `json:"payment_method"`
		PaymentMethodOptions struct {
		} `json:"payment_method_options"`
		PaymentMethodTypes        []string    `json:"payment_method_types"`
		Processing                interface{} `json:"processing"`
		ReceiptEmail              interface{} `json:"receipt_email"`
		Redaction                 interface{} `json:"redaction"`
		Review                    interface{} `json:"review"`
		SetupFutureUsage          interface{} `json:"setup_future_usage"`
		Shipping                  interface{} `json:"shipping"`
		StatementDescriptor       interface{} `json:"statement_descriptor"`
		StatementDescriptorSuffix interface{} `json:"statement_descriptor_suffix"`
		Status                    string      `json:"status"`
		TransferData              interface{} `json:"transfer_data"`
		TransferGroup             interface{} `json:"transfer_group"`
	}

	CreateIntentReq struct {
		Amount      int64  `json:"amount" validate:"required"`
		Email       string `json:"email" validate:"required"`
		Phone       string `json:"phone" `
		Description string `json:"description"`
	}

	CreateIntentRes struct {
		ID               string `json:"id"`
		Object           string `json:"object"`
		Amount           int    `json:"amount"`
		AmountCapturable int    `json:"amount_capturable"`
		AmountDetails    struct {
			Tip struct {
			} `json:"tip"`
		} `json:"amount_details"`
		AmountReceived          int         `json:"amount_received"`
		Application             interface{} `json:"application"`
		ApplicationFeeAmount    interface{} `json:"application_fee_amount"`
		AutomaticPaymentMethods interface{} `json:"automatic_payment_methods"`
		CanceledAt              interface{} `json:"canceled_at"`
		CancellationReason      interface{} `json:"cancellation_reason"`
		CaptureMethod           string      `json:"capture_method"`
		Charges                 struct {
			Object  string        `json:"object"`
			Data    []interface{} `json:"data"`
			HasMore bool          `json:"has_more"`
			Url     string        `json:"url"`
		} `json:"charges"`
		ClientSecret       interface{} `json:"client_secret"`
		ConfirmationMethod string      `json:"confirmation_method"`
		Created            int         `json:"created"`
		Currency           string      `json:"currency"`
		Customer           interface{} `json:"customer"`
		Description        interface{} `json:"description"`
		Invoice            interface{} `json:"invoice"`
		LastPaymentError   interface{} `json:"last_payment_error"`
		Livemode           bool        `json:"livemode"`
		Metadata           struct {
		} `json:"metadata"`
		NextAction           interface{} `json:"next_action"`
		OnBehalfOf           interface{} `json:"on_behalf_of"`
		PaymentMethod        interface{} `json:"payment_method"`
		PaymentMethodOptions struct {
		} `json:"payment_method_options"`
		PaymentMethodTypes        []string    `json:"payment_method_types"`
		Processing                interface{} `json:"processing"`
		ReceiptEmail              interface{} `json:"receipt_email"`
		Redaction                 interface{} `json:"redaction"`
		Review                    interface{} `json:"review"`
		SetupFutureUsage          interface{} `json:"setup_future_usage"`
		Shipping                  interface{} `json:"shipping"`
		StatementDescriptor       interface{} `json:"statement_descriptor"`
		StatementDescriptorSuffix interface{} `json:"statement_descriptor_suffix"`
		Status                    string      `json:"status"`
		TransferData              interface{} `json:"transfer_data"`
		TransferGroup             interface{} `json:"transfer_group"`
	}

	CaptureIntentRes struct {
		ID               string `json:"id"`
		Object           string `json:"object"`
		Amount           int    `json:"amount"`
		AmountCapturable int    `json:"amount_capturable"`
		AmountDetails    struct {
			Tip struct {
			} `json:"tip"`
		} `json:"amount_details"`
		AmountReceived          int         `json:"amount_received"`
		Application             interface{} `json:"application"`
		ApplicationFeeAmount    interface{} `json:"application_fee_amount"`
		AutomaticPaymentMethods interface{} `json:"automatic_payment_methods"`
		CanceledAt              interface{} `json:"canceled_at"`
		CancellationReason      interface{} `json:"cancellation_reason"`
		CaptureMethod           string      `json:"capture_method"`
		Charges                 struct {
			Object string `json:"object"`
			Data   []struct {
				Id                   string      `json:"id"`
				Object               string      `json:"object"`
				Amount               int         `json:"amount"`
				AmountCaptured       int         `json:"amount_captured"`
				AmountRefunded       int         `json:"amount_refunded"`
				Application          interface{} `json:"application"`
				ApplicationFee       interface{} `json:"application_fee"`
				ApplicationFeeAmount interface{} `json:"application_fee_amount"`
				BalanceTransaction   string      `json:"balance_transaction"`
				BillingDetails       struct {
					Address struct {
						City       interface{} `json:"city"`
						Country    interface{} `json:"country"`
						Line1      interface{} `json:"line1"`
						Line2      interface{} `json:"line2"`
						PostalCode interface{} `json:"postal_code"`
						State      interface{} `json:"state"`
					} `json:"address"`
					Email interface{} `json:"email"`
					Name  interface{} `json:"name"`
					Phone interface{} `json:"phone"`
				} `json:"billing_details"`
				CalculatedStatementDescriptor string      `json:"calculated_statement_descriptor"`
				Captured                      bool        `json:"captured"`
				Created                       int         `json:"created"`
				Currency                      string      `json:"currency"`
				Customer                      interface{} `json:"customer"`
				Description                   string      `json:"description"`
				Disputed                      bool        `json:"disputed"`
				FailureBalanceTransaction     interface{} `json:"failure_balance_transaction"`
				FailureCode                   interface{} `json:"failure_code"`
				FailureMessage                interface{} `json:"failure_message"`
				FraudDetails                  struct {
				} `json:"fraud_details"`
				Invoice  interface{} `json:"invoice"`
				Livemode bool        `json:"livemode"`
				Metadata struct {
				} `json:"metadata"`
				OnBehalfOf interface{} `json:"on_behalf_of"`
				Outcome    struct {
					NetworkStatus string      `json:"network_status"`
					Reason        interface{} `json:"reason"`
					RiskLevel     string      `json:"risk_level"`
					RiskScore     int         `json:"risk_score"`
					SellerMessage string      `json:"seller_message"`
					Type          string      `json:"type"`
				} `json:"outcome"`
				Paid                 bool   `json:"paid"`
				PaymentIntent        string `json:"payment_intent"`
				PaymentMethod        string `json:"payment_method"`
				PaymentMethodDetails struct {
					Card struct {
						Brand  string `json:"brand"`
						Checks struct {
							AddressLine1Check      interface{} `json:"address_line1_check"`
							AddressPostalCodeCheck interface{} `json:"address_postal_code_check"`
							CvcCheck               interface{} `json:"cvc_check"`
						} `json:"checks"`
						Country      string      `json:"country"`
						ExpMonth     int         `json:"exp_month"`
						ExpYear      int         `json:"exp_year"`
						Fingerprint  string      `json:"fingerprint"`
						Funding      string      `json:"funding"`
						Installments interface{} `json:"installments"`
						Last4        string      `json:"last4"`
						Mandate      interface{} `json:"mandate"`
						Moto         interface{} `json:"moto"`
						Network      string      `json:"network"`
						ThreeDSecure interface{} `json:"three_d_secure"`
						Wallet       interface{} `json:"wallet"`
					} `json:"card"`
					Type string `json:"type"`
				} `json:"payment_method_details"`
				ReceiptEmail  interface{} `json:"receipt_email"`
				ReceiptNumber string      `json:"receipt_number"`
				ReceiptUrl    string      `json:"receipt_url"`
				Redaction     interface{} `json:"redaction"`
				Refunded      bool        `json:"refunded"`
				Refunds       struct {
					Object  string        `json:"object"`
					Data    []interface{} `json:"data"`
					HasMore bool          `json:"has_more"`
					Url     string        `json:"url"`
				} `json:"refunds"`
				Review                    interface{} `json:"review"`
				Shipping                  interface{} `json:"shipping"`
				SourceTransfer            interface{} `json:"source_transfer"`
				StatementDescriptor       interface{} `json:"statement_descriptor"`
				StatementDescriptorSuffix interface{} `json:"statement_descriptor_suffix"`
				Status                    string      `json:"status"`
				TransferData              interface{} `json:"transfer_data"`
				TransferGroup             interface{} `json:"transfer_group"`
			} `json:"data"`
			HasMore bool   `json:"has_more"`
			Url     string `json:"url"`
		} `json:"charges"`
		ClientSecret       string      `json:"client_secret"`
		ConfirmationMethod string      `json:"confirmation_method"`
		Created            int         `json:"created"`
		Currency           string      `json:"currency"`
		Customer           interface{} `json:"customer"`
		Description        string      `json:"description"`
		Invoice            interface{} `json:"invoice"`
		LastPaymentError   interface{} `json:"last_payment_error"`
		Livemode           bool        `json:"livemode"`
		Metadata           struct {
		} `json:"metadata"`
		NextAction           interface{} `json:"next_action"`
		OnBehalfOf           interface{} `json:"on_behalf_of"`
		PaymentMethod        string      `json:"payment_method"`
		PaymentMethodOptions struct {
		} `json:"payment_method_options"`
		PaymentMethodTypes        []string    `json:"payment_method_types"`
		Processing                interface{} `json:"processing"`
		ReceiptEmail              interface{} `json:"receipt_email"`
		Redaction                 interface{} `json:"redaction"`
		Review                    interface{} `json:"review"`
		SetupFutureUsage          interface{} `json:"setup_future_usage"`
		Shipping                  interface{} `json:"shipping"`
		StatementDescriptor       interface{} `json:"statement_descriptor"`
		StatementDescriptorSuffix interface{} `json:"statement_descriptor_suffix"`
		Status                    string      `json:"status"`
		TransferData              interface{} `json:"transfer_data"`
		TransferGroup             interface{} `json:"transfer_group"`
	}

	GetIntentRes struct {
	}

	GetIntentsRes struct {
		Intents []*PaymentIntent `json:"intents"`
	}

	CreateRefundRes struct {
		ID                 string      `json:"id"`
		Object             string      `json:"object"`
		Amount             int         `json:"amount"`
		BalanceTransaction interface{} `json:"balance_transaction"`
		Charge             string      `json:"charge"`
		Created            int         `json:"created"`
		Currency           string      `json:"currency"`
		Metadata           struct {
		} `json:"metadata"`
		PaymentIntent          interface{} `json:"payment_intent"`
		Reason                 interface{} `json:"reason"`
		ReceiptNumber          interface{} `json:"receipt_number"`
		SourceTransferReversal interface{} `json:"source_transfer_reversal"`
		Status                 string      `json:"status"`
		TransferReversal       interface{} `json:"transfer_reversal"`
	}
)
