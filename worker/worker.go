package worker

import (
	"fmt"
	"notification-service/domain"
)

func ProcessEvents(events chan domain.Event) {
	for event := range events {
		go func(e domain.Event) {
			// Имитация обработки ивента
			fmt.Printf("Processed: %+v", event)
		}(event)
	}
}
