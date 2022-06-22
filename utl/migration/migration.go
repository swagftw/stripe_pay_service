package migration

import (
	"context"

	"gorm.io/gorm"

	"github.com/swagftw/stripe_pay_service/pkg/payments"
	"github.com/swagftw/stripe_pay_service/utl/config"
	"github.com/swagftw/stripe_pay_service/utl/logger"
	"github.com/swagftw/stripe_pay_service/utl/storage"
)

func Migrate() {
	err := config.InitConfig("./utl/config/config.local.yaml", "./.env")
	if err != nil {
		return
	}

	logger.Logger.Debug(context.TODO(), "Config loaded and getting db connection", nil)

	db, err := storage.NewPostgresDB()
	if err != nil {
		return
	}

	err = db.Transaction(func(tx *gorm.DB) error {
		err = db.Exec("CREATE SCHEMA IF NOT EXISTS payment;").Error
		if err != nil {
			return err
		}

		// create payments related table
		err = db.AutoMigrate(&payments.PaymentIntent{}, &payments.Refund{})
		if err != nil {
			return err
		}

		return err
	})

	if err != nil {
		logger.Logger.Error(context.TODO(), "Error executing migration", err)
		panic(err)
	}
}
