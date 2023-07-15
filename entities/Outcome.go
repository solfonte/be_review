package entities

type Outcome struct {
	result Result
	events map[string][]Event
}

func NewOutcome(events []Event) *Outcome {

	eventsPerType := make(map[string][]Event)
	for _, event := range events {

		eventType := event.GetType()
		_, hasKey := eventsPerType[eventType]

		if !hasKey {
			eventsPerType[eventType] = []Event{event}
		} else {
			eventsPerType[eventType] = append(eventsPerType[eventType], event)
		}
	}

	return &Outcome{result: Result{}, events: eventsPerType}
}
