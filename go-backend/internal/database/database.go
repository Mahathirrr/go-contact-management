package database

import (
	"database/sql"
	"fmt"
	"go-backend/internal/config"
	"go-backend/internal/logger"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDatabase(cfg *config.DatabaseConfig) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Name,
	)

	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}

	if err = DB.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	logger.Info("Database connected successfully")
	return nil
}

func CloseDatabase() {
	if DB != nil {
		DB.Close()
	}
}