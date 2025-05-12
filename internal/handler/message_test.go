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
	t.Run("Group: Post", func(t *testing.T) {
		t.Run("Should create message When sent correct post", func(t *testing.T) {
			msg := model.Message{
				Text: "message",
			}
			body, _ := json.Marshal(msg)
			messageGatewayMock := mocks.NewMessageGateway(t)
			handler := NewMessageHandler(messageGatewayMock)

			messageGatewayMock.EXPECT().Add(msg)

			r := httptest.NewRequest(http.MethodPost, "/message", bytes.NewReader(body))
			w := httptest.NewRecorder()
			handler.ServeHTTP(w, r)

			assert.Equal(t, http.StatusCreated, w.Code)
		})
		t.Run("Should return bad request When payload is invalid", func(t *testing.T) {
			messageGatewayMock := mocks.NewMessageGateway(t)
			handler := NewMessageHandler(messageGatewayMock)
			r := httptest.NewRequest(http.MethodPost, "/message", bytes.NewBufferString("{invalid json}"))
			w := httptest.NewRecorder()

			handler.ServeHTTP(w, r)

			assert.Equal(t, http.StatusBadRequest, w.Code)
			assert.Equal(t, w.Body.String(), "Invalid message\n")
		})
	})

	t.Run("Group: Get", func(t *testing.T) {
		t.Run("Should return all messages When get request is successful", func(t *testing.T) {
			expectedMessage := []model.Message{
				{Text: "Hello"},
				{Text: "World"},
			}
			messageGatewayMock := mocks.NewMessageGateway(t)
			handler := NewMessageHandler(messageGatewayMock)
			r := httptest.NewRequest(http.MethodGet, "/message", nil)
			w := httptest.NewRecorder()

			messageGatewayMock.EXPECT().GetAll().Return(expectedMessage).Times(1)

			handler.ServeHTTP(w, r)

			var result []model.Message
			err := json.Unmarshal(w.Body.Bytes(), &result)
			assert.NoError(t, err)
			assert.Equal(t, expectedMessage, result)
		})
	})

	t.Run("Should return invalid method When not using implemented methods", func(t *testing.T) {
		handler := NewMessageHandler(nil)
		r := httptest.NewRequest(http.MethodPut, "/message", nil)
		w := httptest.NewRecorder()

		handler.ServeHTTP(w, r)

		assert.Equal(t, http.StatusMethodNotAllowed, w.Code)
	})
}
