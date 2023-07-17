package entities
type Outcome struct {
	totalPoints int
	bonusPoints int
	amountScores int
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
}

func (o *Outcome) AddBonusPoints(points int) {
	o.bonusPoints += points
}


func (o *Outcome) WinMatch() {
	m := make(map[string]string)
	event := NewEvent("win", "90", &m)
	o.events["win"] = []*Event{event}
}

func (o *Outcome) DrawMatch() {
	m := make(map[string]string)
	e := NewEvent("draw", "90", &m)
	e.SetPoints(1)
	o.events["draw"] = []*Event{e}

}

func (o *Outcome) GetResults() (int, int, int) {
	return o.totalPoints, o.bonusPoints, o.amountScores
}

func (o *Outcome) AssignPointsIfWinner() {

	event, isWinner := o.events["win"]

	if isWinner {
		o.totalPoints += event[0].GetFinalPoints()
	}

}

func (o *Outcome) AssignPointsIfDraw() {

	event, isDraw := o.events["draw"]

	if isDraw {
		o.totalPoints += event[0].GetFinalPoints()
	}

}

func (o *Outcome) SetAmountScores(amountScores int) {
	o.amountScores = amountScores
}