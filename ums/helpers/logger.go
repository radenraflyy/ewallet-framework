package helpers

import (
	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

func SetUpLogger() {
	log := logrus.New()

	log.SetFormatter(&logrus.JSONFormatter{
		PrettyPrint: true,
	})

	log.Info("Logger initialized")
	Logger = log
}
