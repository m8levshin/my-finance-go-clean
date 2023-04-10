package logger

import (
	"github.com/mlevshin/my-finance-go-clean/internal/domain"
	"github.com/sirupsen/logrus"
)

type logrusLogger struct {
	logger *logrus.Logger
}

func NewLogger(logLevel, logFormat string) domain.Logger {
	logger := logrus.New()
	l, err := logrus.ParseLevel(logLevel)
	if err != nil {
		logrus.WithField("level", logLevel).Warn("Invalid log level, fallback to 'info'")
	} else {
		logrus.SetLevel(l)
	}

	switch logFormat {
	case "json":
		logrus.SetFormatter(&logrus.JSONFormatter{})
	default:
	case "text":
		logrus.SetFormatter(&logrus.TextFormatter{})
	}

	return &logrusLogger{logger: logger}
}

func (l logrusLogger) Info(args ...interface{}) {
	l.logger.Info(args)
}

func (l logrusLogger) Error(args ...interface{}) {
	l.logger.Error(args)
}
