package log

import (
	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

func init() {
	InitLogger("info", "text")
}

func InitLogger(logLevel, logFormat string) {
	initLogger := logrus.New()
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

	logger = initLogger
}

func Info(args ...interface{}) {
	logger.Info(args)
}

func Error(args ...interface{}) {
	logger.Error(args)
}
