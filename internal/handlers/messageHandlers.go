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
	// Получение всех сообщений из сервиса
	allMessages, err := h.Service.GetAllMessages()
	if err != nil {
		return nil, err
	}

	// Создаем переменную респон типа 200джейсонРеспонс
	// Которую мы потом передадим в качестве ответа
	response := messages.GetMessages200JSONResponse{}

	// Заполняем слайс response всеми сообщениями из БД
	for _, msg := range allMessages {
		message := messages.Message{
			Id:      &msg.ID,
			Message: &msg.Text,
		}
		response = append(response, message)
	}

	// САМОЕ ПРЕКРАСНОЕ. Возвращаем просто респонс и nil!
	return response, nil
}

func (h *Handler) PostMessages(_ context.Context, request messages.PostMessagesRequestObject) (messages.PostMessagesResponseObject, error) {
	// Распаковываем тело запроса напрямую, без декодера!
	messageRequest := request.Body
	// Обращаемся к сервису и создаем сообщение
	messageToCreate := messagesService.Message{Text: *messageRequest.Message}
	createdMessage, err := h.Service.CreateMessage(messageToCreate)

	if err != nil {
		return nil, err
	}
	// создаем структуру респонс
	response := messages.PostMessages201JSONResponse{
		Id:      &createdMessage.ID,
		Message: &createdMessage.Text,
	}
	// Просто возвращаем респонс!
	return response, nil
}

func (h *Handler) PatchMessages(ctx context.Context, request messages.PatchMessagesRequestObject) (messages.PatchMessagesResponseObject, error) {
	// Извлекаем тело запроса для обновления сообщения
	messageRequest := request.Body
	// Обновляем сообщение через сервис
	messageToUpdate := messagesService.Message{
		Text: *messageRequest.Message,
	}

	updatedMessage, err := h.Service.UpdateMessageByID(int(*messageRequest.Id), messageToUpdate)
	if err != nil {
		return nil, err
	}
	// Формируем ответ
	response := messages.PatchMessages200JSONResponse{
		Id:      &updatedMessage.ID,
		Message: &updatedMessage.Text,
	}
	return response, nil
}

func (h *Handler) DeleteMessages(ctx context.Context, request messages.DeleteMessagesRequestObject) (messages.DeleteMessagesResponseObject, error) {
	// Извлекаем ID сообщения для удаления
	id := request.Params.Id

	// Удаляем сообщение через сервис
	err := h.Service.DeleteMessageByID(id)
	if err != nil {
		return nil, err
	}
	// Возвращаем успешный ответ без контента
	return messages.DeleteMessages204Response{}, nil
}

func NewHandler(service *messagesService.MessageService) *Handler {
	return &Handler{Service: service}
}
