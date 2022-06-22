package main

import (
	"flag"

	"github.com/swagftw/stripe_pay_service/pkg/api"
	"github.com/swagftw/stripe_pay_service/utl/config"
	"github.com/swagftw/stripe_pay_service/utl/logger"
	"github.com/swagftw/stripe_pay_service/utl/migration"
)

func main() {
	cfgPath := flag.String("config", "./utl/config/config.local.yaml", "config file")
	flag.Parse()

	// init log
	logger.InitLogger()

	// init config
	err := config.InitConfig(*cfgPath, "./.env")
	if err != nil {
		return
	}

	// run migration before starting server
	migration.Migrate()

	api.Start()
}
