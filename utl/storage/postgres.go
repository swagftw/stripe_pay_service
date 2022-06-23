package storage

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/swagftw/stripe_pay_service/utl/config"
	"github.com/swagftw/stripe_pay_service/utl/constant"
)

// NewPostgresDB creates a new postgres db connection.
func NewPostgresDB() (*gorm.DB, error) {
	cfg := config.GetGlobalConfig()

	dbCfg := cfg.GetDBConfig()
	user := dbCfg.User
	password := dbCfg.Password
	host := net.JoinHostPort(dbCfg.Host, dbCfg.Port)
	dbname := dbCfg.Database

	sqlDB, err := sql.Open("pgx", fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", user, password, host, dbname))
	if err != nil {
		return nil, err
	}

	dbLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Millisecond * 200, // Slow SQL threshold
			LogLevel:                  logger.Error,           // Log level
			IgnoreRecordNotFoundError: false,                  // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,                   // Disable color
		},
	)

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{
		Logger: dbLogger,
	})
	if err != nil {
		return nil, err
	}

	return gormDB, nil
}

// GetGormDBFromContext returns gorm db from context.
func GetGormDBFromContext(ctx context.Context, db *gorm.DB) *gorm.DB {
	if ctx.Value(constant.PostgresTxKey) != nil {
		return ctx.Value(constant.PostgresTxKey).(*gorm.DB)
	}

	return db.WithContext(ctx)
}

type GormBase struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
