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
		// create generate uid extension
		err = db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";\n\nCREATE EXTENSION IF NOT EXISTS pgcrypto;\n\nCREATE OR REPLACE FUNCTION generate_uid(size INT) RETURNS TEXT AS $$\nDECLARE\ncharacters TEXT := 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789';\n  bytes BYTEA := gen_random_bytes(size);\n  l INT := length(characters);\n  i INT := 0;\noutput TEXT := '';\nBEGIN\n  WHILE i < size LOOP\n    output := output || substr(characters, get_byte(bytes, i) % l + 1, 1);\n    i := i + 1;\nEND LOOP;\n  RETURN output;\nEND;\n$$ LANGUAGE plpgsql VOLATILE;").Error
		if err != nil {
			return err
		}

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
