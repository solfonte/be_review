package entities

type Event struct {
	eventType string
	time string
	eventDetails map[string]string
}

func NewEvent(eventType string, time string, eventDetails map[string]string) Event {
	return Event {eventType: eventType, time: time, eventDetails: eventDetails}
}

func (e *Event) GetType() string {
	return e.eventType
}