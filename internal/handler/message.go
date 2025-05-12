package handler

import (
	"encoding/json"
	"net/http"

	"github.com/krakatoa/go-project-start-template/internal/model"
)

type MessageHandler struct {
	repo MessageGateway
}

//go:generate mockery --with-expecter=true --name=MessageGateway
type MessageGateway interface {
	Add(msg model.Message)
	GetAll() []model.Message
}

func NewMessageHandler(repo MessageGateway) *MessageHandler {
	return &MessageHandler{
		repo: repo,
	}
}

func (h *MessageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.handlePost(w, r)
	case http.MethodGet:
		h.handleGet(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *MessageHandler) handlePost(w http.ResponseWriter, r *http.Request) {
	var msg model.Message
	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		http.Error(w, "Invalid message", http.StatusBadRequest)
		return
	}

	h.repo.Add(msg)
	w.WriteHeader(http.StatusCreated)
}

func (h *MessageHandler) handleGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(h.repo.GetAll())
}
