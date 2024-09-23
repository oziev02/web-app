package messagesService

import "gorm.io/gorm"

type MessageRepository interface {
	CreateMessage(message Message) (Message, error)
	GetAllMessages() ([]Message, error)
	UpdateMessageByID(id int, message Message) (Message, error)
	DeleteMessageByID(id int) error
}

type messageRepository struct {
	db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) *messageRepository {
	return &messageRepository{db: db}
}

func (r *messageRepository) CreateMessage(message Message) (Message, error) {
	result := r.db.Create(&message)
	if result.Error != nil {
		return Message{}, result.Error
	}
	return message, nil
}

func (r *messageRepository) GetAllMessages() ([]Message, error) {
	var messages []Message
	err := r.db.Find(&messages).Error
	return messages, err
}

func (r *messageRepository) UpdateMessageByID(id int, message Message) (Message, error) {
	var existingMessage Message
	result := r.db.First(&existingMessage, id)
	if result.Error != nil {
		return Message{}, result.Error
	}

	existingMessage.Text = message.Text
	r.db.Save(&existingMessage)

	return existingMessage, nil
}

func (r *messageRepository) DeleteMessageByID(id int) error {
	result := r.db.Delete(&Message{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
