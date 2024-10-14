package storage

import (
	"sync"

	"notification-service/domain"
)

type Storage struct {
	mu     sync.RWMutex
	events []domain.Event
}

func New() *Storage {
	return &Storage{
		events: make([]domain.Event, 0),
	}
}

func (s *Storage) AddEvent(event domain.Event) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.events = append(s.events, event)
}

func (s *Storage) GetEvents() []domain.Event {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.events
}
