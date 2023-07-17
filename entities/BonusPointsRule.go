package entities
type BonusPointsRule struct {
	ruleType		string
	event             string
	minimunOcurrences int
	bonusPoints       int
	player            string
	afterTime         []string
	distance 		string
}

func NewBonusPointsRule(event string, minimunOcurrences int, bonusPoints int, player string, afterTime []string, distance string) BonusPointsRule {
	return BonusPointsRule{
		ruleType:			"bonusPoints",
		event:             event,
		minimunOcurrences: minimunOcurrences,
		bonusPoints:       bonusPoints,
		player:            player,
		afterTime:         afterTime,
		distance: distance,
	}
}


func (r *BonusPointsRule) AppliesToEvent(event Event) bool {

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


func (r *BonusPointsRule) Apply (eventsMap map[string][]*Event) int {

	bonusPoints := 0
	 events, hasEvent := eventsMap[r.event]
	if !hasEvent {
		return 0
	}

	if r.minimunOcurrences > 0{ 
		if len(events) < r.minimunOcurrences {
			return 0
		} 
		return r.bonusPoints
	}

	for _, event := range events {

		if r.AppliesToEvent(*event) {
			bonusPoints += r.bonusPoints
		}
	} 

	return bonusPoints

}


