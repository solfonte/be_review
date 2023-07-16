package entities


type BonusPointsRule struct {
	ruleType		string
	event             string
	minimunOcurrences int
	bonusPoints       int
	player            string
	afterTime         string
	distance 		string
}

func NewBonusPointsRule(event string, minimunOcurrences int, bonusPoints int, player string, afterTime string, distance string) BonusPointsRule {
	return BonusPointsRule{
		ruleType:			"bonusPoints",
		event:             event,
		minimunOcurrences: minimunOcurrences,
		bonusPoints:       bonusPoints,
		player:            player,
		afterTime:         afterTime,
	}
}


func (r *BonusPointsRule) AppliesToEvent(event Event) bool {

	isPlayer := false
	//isDistance := false
	isAfterTime := false

	if len(r.player) == 0 || event.GetPlayer() == r.player {
		isPlayer = true
	}

	if len(r.afterTime) == 0 || event.GetTime() == r.afterTime {
		isAfterTime = true
	}

	return isAfterTime && isPlayer
}


func (r *BonusPointsRule) Apply (eventsMap map[string][]*Event) int {

	bonusPoints := 0
	 events, hasEvent := eventsMap[r.event]
	if !hasEvent {
		return 0
	}

	if r.minimunOcurrences > 0{
		//aplica a la cantidad 
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


