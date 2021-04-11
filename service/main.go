package main

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"tratnik.net/service/internal/config"
	"tratnik.net/service/internal/handler"
	"tratnik.net/service/internal/middleware"
	"tratnik.net/service/internal/repository"
	"tratnik.net/service/internal/service"
	"tratnik.net/service/pkg/http/response"
	"tratnik.net/service/pkg/http/server"
	"tratnik.net/service/pkg/nats"
	"tratnik.net/service/pkg/postgres"
)

func main() {
	c := config.GetConfigFromFile("")

	// A workaround for lack of a deployment script. Migrate should not be called here.
	// Wanted everything to fit into a single clean docker-compose.
	time.Sleep(time.Second * 5)
	postgres.Migrate(c.Database)

	router := InitRouter()

	db := postgres.New(c.Database)
	msgBroker := nats.New(c.MessageBroker)

	accountRepo := repository.NewAccount(db)
	messageRepo := repository.NewMessage(msgBroker, "demo")

	messageSrvc := service.NewMessage(accountRepo, messageRepo)

	_ = handler.NewMessage(router, messageSrvc)

	server.Serve(c.Server, router)
}

func InitRouter() *mux.Router {
	r := mux.NewRouter()

	r.Use(
		middleware.Recoverer(response.JSON),
	)

	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		response.JSON(w, nil, http.StatusNotFound)
	})

	r.MethodNotAllowedHandler = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		response.JSON(w, nil, http.StatusMethodNotAllowed)
	})

	return r
}
