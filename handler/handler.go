package handler

import (
	"encoding/json"
	"net/http"

	"notification-service/domain"
	"notification-service/storage"
)

type Handler struct {
	storage *storage.Storage
	events  chan domain.Event
}

func New(storage *storage.Storage, events chan domain.Event) *Handler {
	return &Handler{
		storage: storage,
		events:  events,
	}
}

func (h *Handler) InitRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("/api/v1/", http.HandlerFunc(h.AddEvent))

	return mux
}

func (h *Handler) AddEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.sendErrorResponse(w, http.StatusBadRequest, "Method not allowed")
		return
	}

	var event domain.Event
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		h.sendErrorResponse(w, http.StatusBadRequest, "Invalid payload")
		return
	}

	h.storage.AddEvent(event)
	h.events <- event
	h.writeJSON(w, http.StatusCreated, map[string]string{"message": "New event added"})
}

func (h *Handler) sendErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := map[string]string{"error": message}
	_ = json.NewEncoder(w).Encode(response)
}

func (h *Handler) writeJSON(w http.ResponseWriter, statusCode int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(v)
}
