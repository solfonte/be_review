package entities

type MatchRule struct {
	ruleType	string
	event  string
	points int
}

func NewMatchRule(event string, points int) MatchRule {
	return MatchRule{
		ruleType:	"match",
		event:  	event,
		points: 	points,
	}
}

func (matchRule MatchRule) Apply (eventsMap map[string][]*Event) {	
	events, hasKey := eventsMap["win"]
	if !hasKey {
		return
	}

	events[0].SetPoints(matchRule.points)
}

