package service

import (
	"fmt"
	"sync"

	"tratnik.net/client/internal/model"
	"tratnik.net/client/internal/repository"
)

type IMessage interface {
	Listen(filter model.Filter)
}

var _ IMessage = (*Message)(nil)

type Message struct {
	messageRepo repository.IMessage
}

func NewMessage(messageRepo repository.IMessage) *Message {
	return &Message{
		messageRepo: messageRepo,
	}
}

func (s *Message) Listen(filter model.Filter) {
	// Subscribe is async so use a WaitGroup for parent goroutine to stay alive
	wg := sync.WaitGroup{}
	wg.Add(1)

	if err := s.messageRepo.Subscribe(s.process(filter)); err != nil {
		wg.Done()
	}

	wg.Wait()
}

func (s *Message) process(filter model.Filter) func(msg model.Message) {
	return func(msg model.Message) {
		if filter.AccountID != 0 && filter.AccountID != msg.AccountID {
			return
		}

		fmt.Printf("AccountID: %3d    Timestamp: %s    Data: %s\n",
			msg.AccountID,
			msg.Timestamp.Format("2006-01-02 15:04:05.000 -0700"),
			string(msg.Data),
		)
	}
}
