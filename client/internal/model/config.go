package model

import (
	"tratnik.net/client/pkg/nats"
)

type Config struct {
	Filter        Filter      `mapstructure:"filter"`
	MessageBroker nats.Config `mapstructure:"message_broker"`
}

type Filter struct {
	AccountID int64 `mapstructure:"account_id"`
}
