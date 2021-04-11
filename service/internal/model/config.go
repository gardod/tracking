package model

import (
	"tratnik.net/service/pkg/http/server"
	"tratnik.net/service/pkg/nats"
	"tratnik.net/service/pkg/postgres"
)

type Config struct {
	Server        server.Config   `mapstructure:"server"`
	Database      postgres.Config `mapstructure:"database"`
	MessageBroker nats.Config     `mapstructure:"message_broker"`
}
