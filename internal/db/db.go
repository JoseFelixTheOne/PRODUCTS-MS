package db

import (
	"log"
	"time"

	"github.com/JoseFelixTheOne/products-ms/internal/config"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func MustConnect(cfg *config.Config) *gorm.DB {
	dial := sqlserver.Open(cfg.DBDsn) //el dialector que se abre aqui, es el traductor entre el orm y transactsql(puede ser con otra sintaxis de sql tambien)
	gormCfg := &gorm.Config{Logger: logger.Default.LogMode(logger.Warn)}
	db, err := gorm.Open(dial, gormCfg)
	if err != nil {
		log.Fatalf("no se pudo conectar a SQL Server: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("sql.DB error: %v", err)
	}
	sqlDB.SetMaxOpenConns(cfg.DBMaxOpenConns)
	sqlDB.SetMaxIdleConns(cfg.DBMaxIdleConns)
	sqlDB.SetConnMaxLifetime(time.Duration(cfg.DBConnMaxLifetimeMin) * time.Minute)
	return db
}
