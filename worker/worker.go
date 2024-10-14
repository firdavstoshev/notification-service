package worker

import (
	"fmt"

	"notification-service/domain"
)

func ProcessEvents(events chan domain.Event) {
	for event := range events {
		// Имитация обработки ивента
		fmt.Printf("Processed: %+v", event)
	}
}
