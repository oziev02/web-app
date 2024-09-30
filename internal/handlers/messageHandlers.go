package handlers

import (
	"context"
	"web-app/internal/messagesService"
	"web-app/internal/web/messages"
)

type Handler struct {
	Service *messagesService.MessageService
}

func (h *Handler) GetMessages(_ context.Context, _ messages.GetMessagesRequestObject) (messages.GetMessagesResponseObject, error) {

	allMessages, err := h.Service.GetAllMessages()
	if err != nil {
		return nil, err
	}

	response := messages.GetMessages200JSONResponse{}

	// Заполняем слайс response всеми сообщениями из БД
	for _, msg := range allMessages {
		message := messages.Message{
			Id:      &msg.ID,
			Message: &msg.Text,
		}
		response = append(response, message)
	}

	return response, nil
}

func (h *Handler) PostMessages(_ context.Context, request messages.PostMessagesRequestObject) (messages.PostMessagesResponseObject, error) {
	// Распаковываем тело запроса напрямую, без декодера!
	messageRequest := request.Body
	// Обращаемся к сервису и создаем сообщение
	messageToCreate := messagesService.Message{
		Text: *messageRequest.Message,
	}

	// Передаем сообщение в сервис, для дальнейшого сохранения
	createdMessage, err := h.Service.CreateMessage(messageToCreate)
	if err != nil {
		return nil, err
	}
	// создаем структуру респонс
	response := messages.PostMessages201JSONResponse{
		Id:      &createdMessage.ID,
		Message: &createdMessage.Text,
	}
	return response, nil
}

func (h *Handler) PatchMessagesId(_ context.Context, request messages.PatchMessagesIdRequestObject) (messages.PatchMessagesIdResponseObject, error) {
	messageRequest := request.Body

	messageToUpdate := messagesService.Message{Text: *messageRequest.Message}

	updatedMessage, err := h.Service.UpdateMessageByID(request.Id, messageToUpdate)
	if err != nil {
		return nil, err
	}

	response := messages.PatchMessagesId200JSONResponse{
		Id:      &updatedMessage.ID,
		Message: &updatedMessage.Text,
	}

	return response, nil
}

func (h *Handler) DeleteMessagesId(_ context.Context, request messages.DeleteMessagesIdRequestObject) (messages.DeleteMessagesIdResponseObject, error) {
	id := request.Id

	err := h.Service.DeleteMessageByID(id)
	if err != nil {
		return nil, err
	}

	return messages.DeleteMessagesId204Response{}, nil
}

func NewHandler(service *messagesService.MessageService) *Handler {
	return &Handler{Service: service}
}
