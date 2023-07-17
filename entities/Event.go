package entities

type Event struct {
	eventType    string
	time         string
	eventDetails *map[string]string
	valueFactor  int
	points       int
}

func NewEvent(eventType string, time string, eventDetails *map[string]string) *Event {
	points := 0
	if eventType == "score" {
		points = 1
	} else if eventType == "win" {
		points = 3
	}
	return &Event{eventType: eventType, time: time, eventDetails: eventDetails, valueFactor: 1, points: points}
}

func (e *Event) GetType() string {
	return e.eventType
}

func (e *Event) SetValueFactor(valueFactor int) {
	e.valueFactor = valueFactor
}

func (e *Event) GetFinalPoints() int {
	return e.valueFactor * e.points
}

func (e *Event) SetPoints(points int) {
	e.points = points
}

func (e *Event) GetPlayer() string {

	p, hasPlayer := (*e.eventDetails)["player"]

	if hasPlayer {
		return p
	}
	return ""
}

func (e *Event) GetTime() string {
	return e.time
}

func (e *Event) GetDistance() string {

	d, hasDistance := (*e.eventDetails)["distance"]

	if hasDistance {
		return d
	}
	return ""
}
