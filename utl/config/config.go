package config

import (
	"context"
	"io/ioutil"
	"sync"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"

	"github.com/swagftw/stripe_pay_service/utl/logger"
)

var config *GlobalConfig

type GlobalConfig struct {
	Server Server `yaml:"server"`
	DB     DB     `yaml:"database"`
	Stripe Stripe `yaml:"stripeclient"`
	mutex  sync.Mutex
}

type Server struct {
	Port     string `yaml:"port"`
	Debug    bool   `yaml:"debug"`
	LogLevel string `yaml:"logLevel"`
	Timeout  int    `yaml:"timeout"`
}

type DB struct {
	User     string `yaml:"user"`
	Password string `yaml:"pass"`
	Host     string `yaml:"host"`
	Database string `yaml:"db"`
	Timeout  int    `yaml:"timeout"`
}

type Stripe struct {
	SecretKey      string `yaml:"secretKey"`
	PublishableKey string `yaml:"publishableKey"`
}

// InitConfig initializes the config.
func InitConfig(path string, envPath string) error {
	configFile, err := ioutil.ReadFile(path)
	if err != nil {
		logger.Logger.Error(context.TODO(), "config file not found", err)

		return err
	}

	// load env variables from file to environment
	err = godotenv.Load(envPath)
	if err != nil {
		logger.Logger.Error(context.TODO(), "failed to load env file", err)

		return err
	}

	viper.AutomaticEnv()

	config = new(GlobalConfig)
	config.mutex = sync.Mutex{}

	unmarshallErr := yaml.Unmarshal(configFile, config)

	loadEnv()

	if unmarshallErr != nil {
		logger.Logger.Error(context.TODO(), "failed to unmarshal config file", err)
	}

	return nil
}

func loadEnv() {
	port := viper.GetString("PORT")
	if port != "" {
		config.Server.Port = port
	}

	debug := viper.GetBool("DEBUG")
	config.Server.Debug = debug

	logLevel := viper.GetString("LOG_LEVEL")
	if logLevel != "" {
		config.Server.LogLevel = logLevel
	}

	timeout := viper.GetInt("TIMEOUT")
	if timeout != 0 {
		config.Server.Timeout = timeout
	}

	dbUser := viper.GetString("POSTGRES_USER")
	if dbUser != "" {
		config.DB.User = dbUser
	}

	dbPassword := viper.GetString("POSTGRES_PASSWORD")
	if dbPassword != "" {
		config.DB.Password = dbPassword
	}

	dbHost := viper.GetString("POSTGRES_HOST")
	if dbHost != "" {
		config.DB.Host = dbHost
	}

	dbDatabase := viper.GetString("POSTGRES_DB")
	if dbDatabase != "" {
		config.DB.Database = dbDatabase
	}

	dbTimout := viper.GetInt("DB_TIMEOUT")
	if dbTimout != 0 {
		config.DB.Timeout = dbTimout
	}

	secretKey := viper.GetString("STRIPE_SECRET_KEY")
	if secretKey != "" {
		config.Stripe.SecretKey = secretKey
	}

	publishableKey := viper.GetString("STRIPE_PUBLISHABLE_KEY")
	if publishableKey != "" {
		config.Stripe.PublishableKey = publishableKey
	}
}

// GetGlobalConfig returns the global config.
func GetGlobalConfig() *GlobalConfig {
	return config
}

// GetServerConfig returns the Server config.
func (c *GlobalConfig) GetServerConfig() *Server {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	return &c.Server
}

// GetDBConfig returns the database config.
func (c *GlobalConfig) GetDBConfig() *DB {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	return &c.DB
}

// GetStripeConfig returns the Stripe config.
func (c *GlobalConfig) GetStripeConfig() *Stripe {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	return &c.Stripe
}
