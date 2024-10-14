package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"notification-service/domain"
	"notification-service/handler"
	"notification-service/storage"
)

func TestProcessorHandler(t *testing.T) {
	// Инициализация хранения событий и обработчика
	storage := storage.New()
	eventsChan := make(chan domain.Event, 1)
	h := handler.New(storage, eventsChan)

	// Пример события
	event := domain.Event{
		OrderType:  "Purchase",
		SessionId:  "29827525-06c9-4b1e-9d9b-7c4584e82f56",
		Card:       "4433**1409",
		EventDate:  "2023-01-04 13:44:52.835626 +00:00",
		WebsiteURL: "https://amazon.com",
	}

	// Преобразование события в JSON
	eventJSON, err := json.Marshal(event)
	if err != nil {
		t.Fatalf("Failed to marshal event: %v", err)
	}

	// Создание запроса
	req, err := http.NewRequest(http.MethodPost, "/api/v1/", bytes.NewBuffer(eventJSON))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Создание записи для тестирования
	rr := httptest.NewRecorder()
	handlerFunc := http.HandlerFunc(h.AddEvent)

	// Выполнение запроса
	handlerFunc.ServeHTTP(rr, req)

	// Проверка статуса ответа
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	// Проверка, что событие было добавлено в хранилище
	events := storage.GetEvents()
	if len(events) == 0 {
		t.Errorf("Expected event to be stored, got none")
	}

	// Проверка, что событие в канале
	select {
	case receivedEvent := <-eventsChan:
		if receivedEvent != event {
			t.Errorf("Expected event %v, got %v", event, receivedEvent)
		}
	default:
		t.Errorf("Expected event to be sent to channel, got none")
	}
}
