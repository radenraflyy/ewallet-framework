package helpers

import (
	"github.com/sirupsen/logrus"
)

func SetUpLogger() *logrus.Logger {
	log := logrus.New()

	log.SetFormatter(&logrus.JSONFormatter{
		PrettyPrint: true,
	})

	log.Info("Logger initialized")
	return log
}
