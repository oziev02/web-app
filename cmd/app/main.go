package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"web-app/internal/database"
	"web-app/internal/handlers"
	"web-app/internal/messagesService"
	"web-app/internal/userService"
	"web-app/internal/web/messages"
	"web-app/internal/web/users"
)

func main() {
	database.InitDB()
	if err := database.DB.AutoMigrate(&messagesService.Message{}, &userService.User{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	messagesRepo := messagesService.NewMessageRepository(database.DB)
	messagesService := messagesService.NewService(messagesRepo)
	messagesHandler := handlers.NewHandler(messagesService)

	userRepo := userService.NewUserRepository(database.DB)
	userService := userService.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover()) //не дает приложению крашнутся

	messagesStrictHandler := messages.NewStrictHandler(messagesHandler, nil)
	messages.RegisterHandlers(e, messagesStrictHandler)

	usersStrictHandler := users.NewStrictHandler(userHandler, nil)
	users.RegisterHandlers(e, usersStrictHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}
