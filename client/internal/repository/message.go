package repository

import (
	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"

	"tratnik.net/client/internal/model"
)

type IMessage interface {
	Subscribe(callback func(msg model.Message)) error
}

var _ IMessage = (*Message)(nil)

type Message struct {
	connection *nats.EncodedConn
	channel    string
}

func NewMessage(connection *nats.EncodedConn, channel string) *Message {
	return &Message{
		connection: connection,
		channel:    channel,
	}
}

func (r *Message) Subscribe(callback func(msg model.Message)) error {
	_, err := r.connection.Subscribe(r.channel, callback)
	if err != nil {
		logrus.WithError(err).WithField("channel", r.channel).Error("Unable to subscribe to msgs")
		return ErrUnknown
	}

	return nil
}
