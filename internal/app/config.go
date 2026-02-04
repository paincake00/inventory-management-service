package app

import (
	"fmt"
	"time"

	"github.com/paincake00/inventory-management-service/internal/utils/env"
)

type Config struct {
	addr string
	db   DBConfig
}

type DBConfig struct {
	address         string
	driver          string
	maxOpenConn     int
	maxIdleConn     int
	maxConnLifetime time.Duration
}

func InitConfig() Config {
	return Config{
		addr: env.GetString("APP_ADDR", ":8080"),
		db: DBConfig{
			address:         getDBUri(),
			driver:          env.GetString("DB_DRIVER", "postgres"),
			maxOpenConn:     env.GetInt("DB_MAX_OPEN_CONS", 30),
			maxIdleConn:     env.GetInt("DB_MAX_IDLE_CONS", 30),
			maxConnLifetime: env.GetDuration("DB_MAX_CONN_LIFETIME", 30*time.Minute),
		},
	}
}

func getDBUri() string {
	schema := env.GetString("DB_DRIVER", "postgres")
	user := env.GetString("POSTGRES_USER", "postgres")
	password := env.GetString("POSTGRES_PASSWORD", "postgres")
	host := env.GetString("POSTGRES_HOST", "localhost")
	port := env.GetInt("POSTGRES_PORT", 5432)
	db := env.GetString("POSTGRES_DB", "postgres")

	return fmt.Sprintf("%s://%s:%s@%s:%d/%s?sslmode=disable", schema, user, password, host, port, db)
}
