package repository

import (
	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"

	"tratnik.net/service/internal/model"
)

type IMessage interface {
	Publish(msg model.Message) error
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

func (r *Message) Publish(msg model.Message) error {
	err := r.connection.Publish(r.channel, msg)
	if err != nil {
		logrus.WithError(err).WithField("msg", msg).Error("Unable to publish msg")
		return ErrUnknown
	}

	return nil
}
