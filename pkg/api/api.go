package api

import (
	"github.com/swagftw/stripe_pay_service/pkg/payments"
	postgres2 "github.com/swagftw/stripe_pay_service/pkg/payments/repository/postgres"
	"github.com/swagftw/stripe_pay_service/transaction/postgres"
	paymentsHTTP "github.com/swagftw/stripe_pay_service/transport/payments"
	"github.com/swagftw/stripe_pay_service/utl/server"
	"github.com/swagftw/stripe_pay_service/utl/storage"
	"github.com/swagftw/stripe_pay_service/utl/stripeclient"
)

func Start() {
	echoServer := server.NewServer()

	// get new postgres database connection
	db, err := storage.NewPostgresDB()
	if err != nil {
		return
	}

	// init postgres transaction
	postgresTx := postgres.NewPostgresTx(db)

	v1Group := echoServer.Group("/api/v1")

	// initialize services
	// init payments service
	payService := payments.NewService(postgresTx, postgres2.NewPaymentsRepo(db), stripeclient.New())

	// init http handlers
	paymentsHTTP.InitHTTPHandlers(payService, v1Group)

	server.StartServer(echoServer)
}
