package entities

import ("fmt")

type Outcome struct {
	totalPoints int
	bonusPoints int
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
		
	return &Outcome{totalPoints: 0, bonusPoints: 0, events: eventsPerType}
}

func (o *Outcome) GetEvents() map[string][]*Event{
	return o.events
}

func (o *Outcome) AddTotalPoints(points int) {
	o.totalPoints += points
	fmt.Println("tegooooo", o.totalPoints)
}

func (o *Outcome) AddBonusPoints(points int) {
	o.bonusPoints += points
}


func (o *Outcome) WinMatch() {
	o.events["win"] = []*Event{NewEvent("win", "90", make(map[string]string))}
}

func (o *Outcome) GetResults() (int, int) {
	return o.totalPoints, o.bonusPoints
}

func (o *Outcome) AssignPointsIfWinner() {

	event, isWinner := o.events["win"]

	if isWinner {
		o.totalPoints += event[0].GetFinalPoints()
	}

}