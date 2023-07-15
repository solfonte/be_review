package entities

type Event struct {
	eventType    string
	time         string
	eventDetails map[string]string
	valueFactor  int
}

func NewEvent(eventType string, time string, eventDetails map[string]string) Event {
	return Event{eventType: eventType, time: time, eventDetails: eventDetails, valueFactor: 1}
}

func (e *Event) GetType() string {
	return e.eventType
}
