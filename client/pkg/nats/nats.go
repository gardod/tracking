package nats

import (
	"fmt"

	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

func New(config Config) *nats.EncodedConn {
	logrus.Info("Setting up message broker")

	c, err := nats.Connect(getURL(config))
	if err != nil {
		logrus.WithError(err).Fatal("Unable to connect to message broker")
	}

	ec, err := nats.NewEncodedConn(c, nats.JSON_ENCODER)
	if err != nil {
		logrus.WithError(err).Fatal("Unable to connect to message broker")
	}

	return ec
}

func getURL(config Config) string {
	return fmt.Sprintf(
		"nats://%s:%d",
		config.Host,
		config.Port,
	)
}
