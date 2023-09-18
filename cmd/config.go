package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/rl404/fairy/monitoring/prometheus/database"
	"github.com/ronyelkahfi/todos/internal/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type config struct {
	App      appConfig      `envconfig:"APP"`
	DB       dbConfig       `envconfig:"DB"`

}

type appConfig struct {
	Env             string        `envconfig:"ENV" default:"dev" validate:"required,oneof=dev prod" mod:"no_space,lcase"` // dev/prod
	Port            string        `envconfig:"PORT" default:"8006" validate:"required" mod:"no_space"`
	ReadTimeout     time.Duration `envconfig:"READ_TIMEOUT" default:"5s" validate:"required,gt=0"`
	WriteTimeout    time.Duration `envconfig:"WRITE_TIMEOUT" default:"5s" validate:"required,gt=0"`
	GracefulTimeout time.Duration `envconfig:"GRACEFUL_TIMEOUT" default:"10s" validate:"required,gt=0"`
	InternalKey     string        `envconfig:"INTERNAL_KEY"`
	Host            string        `envconfig:"HOST" validate:"required,url"`
}


type dbConfig struct {
	Address         string        `envconfig:"ADDRESS" default:"localhost:3306" validate:"required"`
	Name            string        `envconfig:"NAME" default:"payment" validate:"required"`
	User            string        `envconfig:"USER" default:"root" validate:"required"`
	Password        string        `envconfig:"PASSWORD"`
	MaxConnOpen     int           `envconfig:"MAX_CONN_OPEN" default:"10" validate:"required,gt=0"`
	MaxConnIdle     int           `envconfig:"MAX_CONN_IDLE" default:"10" validate:"required,gt=0"`
	MaxConnLifetime time.Duration `envconfig:"MAX_CONN_LIFETIME" default:"60s" validate:"required,gt=0"`
}


const envPath = "../../.env"
const envPrefix = "PAYMENT"
const pubsubTopic = "payment-pubsub"
const esIndex = "logs-payment"

func getConfig() (*config, error) {
	var cfg config

	// Load .env file.
	_ = godotenv.Load(envPath)

	// Convert env to struct.
	if err := envconfig.Process(envPrefix, &cfg); err != nil {
		return nil, err
	}

	// Validate.
	// if err := utils.Validate(&cfg); err != nil {
	// 	return nil, err
	// }

	return &cfg, nil
}

func newDB(cfg dbConfig) (*gorm.DB, error) {
	// Split host and port.
	split := strings.Split(cfg.Address, ":")
	if len(split) != 2 {
		return nil, errors.ErrInvalidDBFormat
	}

	// Prepare dsn and open connection.
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true&loc=Local", cfg.User, cfg.Password, cfg.Address, cfg.Name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		AllowGlobalUpdate: true,
		Logger:            logger.Default.LogMode(logger.Silent),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return nil, err
	}

	tmp, err := db.DB()
	if err != nil {
		return nil, err
	}

	// Set basic config.
	tmp.SetMaxIdleConns(cfg.MaxConnIdle)
	tmp.SetMaxOpenConns(cfg.MaxConnOpen)
	tmp.SetConnMaxLifetime(time.Duration(cfg.MaxConnLifetime) * time.Second)
	
	database.NewGORM(cfg.Name)
	return db, nil
}

func getESIndex(keys ...string) string {
	i := []string{esIndex}
	i = append(i, keys...)
	return strings.Join(i, "-")
}
