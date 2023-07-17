package entities

type ParticularRule struct {
	ruleType		string
	event             string
	minimunOcurrences int
	valueFactor       int
	player            string
	afterTime         string
	distance 			string
}

func NewParticularRule(event string, minimunOcurrences int, valueFactor int, player string, afterTime string, distance string) ParticularRule {
	
	return ParticularRule {
		ruleType:			"particular",
		event:             event,
		minimunOcurrences: minimunOcurrences,
		valueFactor:       valueFactor,
		player:            player,
		afterTime:         afterTime,
		distance: 			distance,
	}
}


func (r *ParticularRule) AppliesToEvent(event *Event) bool {
	isPlayer := false
	isDistance := false
	isAfterTime := false

	if len(r.player) == 0 || event.GetPlayer() == r.player {
		isPlayer = true
	}

	if len(r.afterTime) > 0 {
		eventTime := event.GetTime()
		if len(r.afterTime) == 2 {
			isAfterTime = eventTime[:2] >= r.afterTime[:2]
		} else {
			isAfterTime = eventTime[:2] >= r.afterTime[:2]
			if len(eventTime) > 2 {
				isAfterTime = isAfterTime && eventTime[4:5] >= r.afterTime[4:5]
			}
		}
	}

	if len(r.distance) > 0 {
		eventDistance := event.GetDistance()

		if r.distance[0:1] == "+" {
			isDistance = eventDistance >= r.distance[1:]
		} else {
			isDistance = eventDistance <= r.distance[1:]
		}
	}

	return isAfterTime && isPlayer && isDistance
}


func (particularRule ParticularRule) Apply (eventsMap map[string][]*Event) {
	
	events, hasKey := eventsMap["score"]
	if !hasKey {
		return
	}

	for _, event := range events {
		amountScores := len(events) 
		if particularRule.AppliesToEvent(event) || amountScores >= particularRule.minimunOcurrences{
			event.SetValueFactor(particularRule.valueFactor)
		}
	}
	
}