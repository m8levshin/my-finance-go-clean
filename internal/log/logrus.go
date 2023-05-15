package log

import (
	"github.com/mlevshin/my-finance-go-clean/internal/domain"
	"github.com/sirupsen/logrus"
)

var loggerInstance *logrus.Logger

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
	loggerInstance = initLogger
}

func NewDomainLogger() *domain.Logger {
	var l domain.Logger = &logger{}
	return &l
}

type logger struct {
}

func (*logger) Info(args ...interface{}) {
	loggerInstance.Info(args)
}

func (*logger) Error(args ...interface{}) {
	loggerInstance.Error(args)
}

func (*logger) Fatal(args ...interface{}) {
	loggerInstance.Fatal(args)
}

func Info(args ...interface{}) {
	loggerInstance.Info(args)
}

func Error(args ...interface{}) {
	loggerInstance.Error(args)
}

func Fatal(args ...interface{}) {
	loggerInstance.Fatal(args)
}
