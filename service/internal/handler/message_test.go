package handler

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"tratnik.net/service/internal/model"
	"tratnik.net/service/internal/service"
)

type MockedMessageService struct {
	mock.Mock
}

func (m *MockedMessageService) Create(ctx context.Context, msg model.Message) error {
	args := m.Called(msg.AccountID)
	return args.Error(0)
}

func TestCreate(t *testing.T) {
	mms := &MockedMessageService{}
	mms.On("Create", int64(1)).Return(nil)
	mms.On("Create", int64(2)).Return(service.ErrAccountValidation)
	mms.On("Create", int64(3)).Return(service.ErrAccountRetrieve)
	mms.On("Create", int64(4)).Return(service.ErrMessagePublish)

	messageSrvc := NewMessage(mux.NewRouter(), mms)
	handler := http.HandlerFunc(messageSrvc.create)

	testCases := []struct {
		Body           string
		ExpectedStatus int
	}{
		{`{"account_id":"1","data":"test"}`, http.StatusBadRequest},
		{`{"account_id":1,"data":1}`, http.StatusCreated},
		{`{"account_id":1,"data":"test"}`, http.StatusCreated},
		{`{"account_id":1,"data":{"test":true}}`, http.StatusCreated},
		{`{"account_id":2,"data":"test"}`, http.StatusBadRequest},
		{`{"account_id":3,"data":"test"}`, http.StatusInternalServerError},
		{`{"account_id":4,"data":"test"}`, http.StatusInternalServerError},
	}

	for _, testCase := range testCases {
		body := ioutil.NopCloser(bytes.NewBuffer([]byte(testCase.Body)))
		req, err := http.NewRequest("POST", "/", body)
		if err != nil {
			require.Nil(t, err)
		}
		resp := httptest.NewRecorder()

		handler.ServeHTTP(resp, req)

		assert.Equal(t, testCase.ExpectedStatus, resp.Code)
	}
}
