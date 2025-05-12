package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/krakatoa/go-project-start-template/internal/handler/mocks"
	"github.com/krakatoa/go-project-start-template/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestMessageHandler_AddMessage(t *testing.T) {
	t.Run("", func(t *testing.T) {
		mockRepo := mocks.NewMessageGateway(t)
		handlerUnderTest := NewMessageHandler(mockRepo)

		msg := model.Message{Text: "Hello, World!"}
		body, _ := json.Marshal(msg)

		// Espera que Add seja chamado com a mensagem correta
		mockRepo.EXPECT().Add(msg).Times(1)

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(body))
		w := httptest.NewRecorder()

		handlerUnderTest.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)
	})

	t.Run("Should create message When sent correct post", func(t *testing.T) {
		messageGatewayMock := mocks.NewMessageGateway(t)
		handler := NewMessageHandler(messageGatewayMock)

		msg := model.Message{
			Text: "message",
		}
		body, _ := json.Marshal(msg)

		r := httptest.NewRequest(http.MethodPost, "/message", bytes.NewReader(body))
		w := httptest.NewRecorder()
		handler.ServeHTTP(w, r)
	})
}
