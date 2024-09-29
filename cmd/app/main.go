package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"web-app/internal/database"
	"web-app/internal/handlers"
	"web-app/internal/messagesService"
	"web-app/internal/web/messages"
)

func main() {
	database.InitDB()
	if err := database.DB.AutoMigrate(&messagesService.Message{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	repo := messagesService.NewMessageRepository(database.DB)
	service := messagesService.NewService(repo)
	handler := handlers.NewHandler(service)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover()) //не дает приложению крашнутся

	strictHandler := messages.NewStrictHandler(handler, nil)
	messages.RegisterHandlers(e, strictHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}
