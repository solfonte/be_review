package entities

type ParticularRule struct {
	ruleType          string
	event             string
	minimunOcurrences int
	valueFactor       int
	player            string
	afterTime         []string
	distance          string
}

func NewParticularRule(event string, minimunOcurrences int, valueFactor int, player string, afterTime []string, distance string) ParticularRule {
	return ParticularRule{
		ruleType:          "particular",
		event:             event,
		minimunOcurrences: minimunOcurrences,
		valueFactor:       valueFactor,
		player:            player,
		afterTime:         afterTime,
		distance:          distance,
	}
}

func (r *ParticularRule) AppliesToEvent(event *Event) bool {
	isPlayer := true
	isDistance := true
	isAfterTime := true

	if len(r.player) > 0 && event.GetPlayer() != r.player {
		isPlayer = false
	}
	if len(r.afterTime) > 0 {
		eventTime := event.GetTime()
		for _, time := range r.afterTime {
			if len(time) == 2 {
				isAfterTime = eventTime[:2] >= time[:2]
			} else {
				isAfterTime = eventTime[:2] >= time[:2]
				if len(eventTime) > 2 {
					isAfterTime = isAfterTime && eventTime[4:5] >= time[4:5]
				}
			}
			if isAfterTime {
				break
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

func (particularRule ParticularRule) Apply(eventsMap map[string][]*Event) {

	events, hasKey := eventsMap["score"]
	if !hasKey {
		return
	}

	for _, event := range events {
		amountScores := len(events)
		if particularRule.AppliesToEvent(event) || (particularRule.minimunOcurrences > 0 && amountScores >= particularRule.minimunOcurrences) {
			event.SetValueFactor(particularRule.valueFactor)
		}
	}

}
