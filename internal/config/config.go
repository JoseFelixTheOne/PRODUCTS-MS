package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	AppEnv               string
	AppPort              string
	DBDsn                string
	DBMaxOpenConns       int
	DBMaxIdleConns       int
	DBConnMaxLifetimeMin int
}

func getenv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

func getenvInt(key string, def int) int {
	if v := os.Getenv(key); v != "" {
		if i, err := strconv.Atoi(v); err == nil {
			return i
		}
	}
	return def
}

func Load() *Config {
	_ = godotenv.Load()
	cfg := &Config{
		AppEnv:               getenv("APP_ENV", "dev"),
		AppPort:              getenv("APP_PORT", "8080"),
		DBDsn:                getenv("DB_DSN", ""),
		DBMaxOpenConns:       getenvInt("DB_MAX_OPEN_CONNS", 30),
		DBMaxIdleConns:       getenvInt("DB_MAX_IDLE_CONNS", 10),
		DBConnMaxLifetimeMin: getenvInt("DB_CONN_MAX_LIFETIME_MIN", 60),
	}
	if cfg.DBDsn == "" {
		log.Fatal("DB_DSN no configurado")
	}
	return cfg
}
