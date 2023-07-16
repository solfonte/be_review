package entities


type Outcome struct {
	result *Result
	events map[string][]*Event
}

func NewOutcome(events []*Event) *Outcome {

	eventsPerType := make(map[string][]*Event)
	for _, event := range events {

		eventType := event.GetType()
		_, hasKey := eventsPerType[eventType]
		
		if !hasKey {
			eventsPerType[eventType] = []*Event{event}
			} else {
				eventsPerType[eventType] = append(eventsPerType[eventType], event)
			}
		}
		
	return &Outcome{result: &Result{}, events: eventsPerType}
}

func (o *Outcome) GetEvents() map[string][]*Event{
	return o.events
}

func (o *Outcome) SetTotalPoints(points int) {
	o.result.AddTotalPoints(points)
}


func (o *Outcome) WinMatch() {
	o.events["win"] = []*Event{NewEvent("win", "90", make(map[string]string))}
}

func (o *Outcome) GetResults() *Result {
	return o.result
}

func (o *Outcome) AssignPointsIfWinner() {

	event, isWinner := o.events["win"]

	if isWinner {
		o.result.AddTotalPoints(event[0].GetFinalPoints())
	}

}