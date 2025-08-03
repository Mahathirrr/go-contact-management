package logger

import (
	"go-backend/internal/config"

	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func InitLogger(cfg *config.LoggingConfig) {
	Log = logrus.New()

	// Set log level
	level, err := logrus.ParseLevel(cfg.Level)
	if err != nil {
		level = logrus.InfoLevel
	}
	Log.SetLevel(level)

	// Set log format
	if cfg.Format == "json" {
		Log.SetFormatter(&logrus.JSONFormatter{})
	} else {
		Log.SetFormatter(&logrus.TextFormatter{})
	}
}

func Info(args ...interface{}) {
	Log.Info(args...)
}

func Error(args ...interface{}) {
	Log.Error(args...)
}

func Debug(args ...interface{}) {
	Log.Debug(args...)
}

func Warn(args ...interface{}) {
	Log.Warn(args...)
}

func Fatal(args ...interface{}) {
	Log.Fatal(args...)
}