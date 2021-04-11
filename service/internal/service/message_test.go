package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"tratnik.net/service/internal/model"
	"tratnik.net/service/internal/repository"
)

type MockedAccountRepo struct {
	mock.Mock
}

func (m *MockedAccountRepo) GetByID(ctx context.Context, accountID int64) (*model.Account, error) {
	args := m.Called(accountID)
	return args.Get(0).(*model.Account), args.Error(1)
}

type MockedMessageRepo struct {
	mock.Mock
}

func (m *MockedMessageRepo) Publish(msg model.Message) error {
	args := m.Called(msg.AccountID)
	return args.Error(0)
}

func TestCreate(t *testing.T) {
	mar := &MockedAccountRepo{}
	mar.On("GetByID", int64(1)).Return(&model.Account{ID: 1, Name: "one", IsActive: true}, nil)
	mar.On("GetByID", int64(2)).Return(&model.Account{ID: 2, Name: "two", IsActive: true}, nil)
	mar.On("GetByID", int64(3)).Return(&model.Account{ID: 3, Name: "three", IsActive: false}, nil)
	mar.On("GetByID", int64(4)).Return((*model.Account)(nil), repository.ErrNoResults)
	mar.On("GetByID", int64(5)).Return((*model.Account)(nil), repository.ErrUnknown)

	mmr := &MockedMessageRepo{}
	mmr.On("Publish", int64(1)).Return(nil)
	mmr.On("Publish", int64(2)).Return(repository.ErrUnknown)

	messageSrvc := NewMessage(mar, mmr)

	testCases := []struct {
		AccountID     int64
		ExpectedError error
	}{
		{1, nil},
		{2, ErrMessagePublish},
		{3, ErrAccountValidation},
		{4, ErrAccountValidation},
		{5, ErrAccountRetrieve},
	}

	for _, testCase := range testCases {
		err := messageSrvc.Create(context.Background(), model.Message{AccountID: testCase.AccountID})
		assert.Equal(t, testCase.ExpectedError, err)
	}
}
