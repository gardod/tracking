package service

import (
	"context"

	"tratnik.net/service/internal/model"
	"tratnik.net/service/internal/repository"
)

type IMessage interface {
	Create(ctx context.Context, msg model.Message) error
}

var _ IMessage = (*Message)(nil)

type Message struct {
	accountRepo repository.IAccount
	messageRepo repository.IMessage
}

func NewMessage(accountRepo repository.IAccount, messageRepo repository.IMessage) *Message {
	return &Message{
		accountRepo: accountRepo,
		messageRepo: messageRepo,
	}
}

func (s *Message) Create(ctx context.Context, msg model.Message) error {
	account, err := s.accountRepo.GetByID(ctx, msg.AccountID)
	switch err {
	case nil:
	case repository.ErrNoResults:
		return ErrAccountValidation
	default:
		return ErrAccountRetrieve
	}

	if !account.IsActive {
		return ErrAccountValidation
	}

	if err := s.messageRepo.Publish(msg); err != nil {
		return ErrMessagePublish
	}

	return nil
}
