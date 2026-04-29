package main

import (
	"os"

	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()
	log.Out = os.Stdout

	file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.Out = file
	} else {
		log.Info("Failed to log to file, using default stderr")
	}

	log.SetFormatter(&logrus.JSONFormatter{})
	loggerContext := log.WithFields(logrus.Fields{"foo": "foo"})
	loggerContext.WithFields(logrus.Fields{"bar": "bar"}).Info("bar logged along with foo")
}
