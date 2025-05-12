package repository

import (
	"sync"

	"github.com/krakatoa/go-project-start-template/internal/model"
)

type MessageRepository struct {
	mu       sync.RWMutex
	messages []model.Message
}

func NewMessageRepository() *MessageRepository {
	return &MessageRepository{
		messages: make([]model.Message, 0),
	}
}

func (r *MessageRepository) Add(msg model.Message) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.messages = append(r.messages, msg)
}

func (r *MessageRepository) GetAll() []model.Message {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.messages
}
