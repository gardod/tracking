package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"tratnik.net/service/internal/model"
	"tratnik.net/service/internal/service"
	"tratnik.net/service/pkg/http/response"
)

type Message struct {
	messageService service.IMessage
}

func NewMessage(router *mux.Router, messageService service.IMessage) *Message {
	h := &Message{
		messageService: messageService,
	}
	h.registerRoutes(router)
	return h
}

func (h *Message) registerRoutes(r *mux.Router) {
	r.Path("/").Methods("POST").HandlerFunc(h.create)
}

func (h *Message) create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	msg := model.Message{Timestamp: time.Now()}

	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		err = fmt.Errorf("malformed body: %w", err)
		response.JSON(w, err, http.StatusBadRequest)
		return
	}

	switch err := h.messageService.Create(ctx, msg); err {
	case nil:
	case service.ErrAccountValidation:
		response.JSON(w, err, http.StatusBadRequest)
		return
	default:
		response.JSON(w, nil, http.StatusInternalServerError)
		return
	}

	response.JSON(w, nil, http.StatusCreated)
}
