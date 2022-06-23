package main

import (
	"os"
	"testing"

	"github.com/swagftw/stripe_pay_service/utl/config"
)

func TestMain(m *testing.M) {
	err := config.InitConfig("../../utl/config/config.local.yaml", "../../.env")
	if err != nil {
		os.Exit(1)
	}

	exitCode := m.Run()

	// call with result of m.Run()
	os.Exit(exitCode)
}
