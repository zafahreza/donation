package main

import (
	"donation/app"
	"donation/handler"
	"donation/repository"
	"donation/service"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func main() {
	db := app.NewSetupDB()
	validate := validator.New()
	chc := app.NewChacheDB()

	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, chc, db, validate)
	userHandler := handler.NewUserHanlder(userService)

	router := app.NewRouter(userHandler)

	server := http.Server{
		Addr:    "localhost:3333",
		Handler: router,
	}

	fmt.Println("server is running.....")

	server.ListenAndServe()
}
