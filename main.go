package main

import (
	"donation/app"
	"donation/handler"
	"donation/repository"
	"donation/service"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func main() {
	db := app.NewSetupDB()
	validate := validator.New()

	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, db, validate)
	userHandler := handler.NewUserHanlder(userService)

	router := app.NewRouter(userHandler)

	server := http.Server{
		Addr:    "localhost:3333",
		Handler: router,
	}

	server.ListenAndServe()
}
