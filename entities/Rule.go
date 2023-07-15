package entities

type Rule interface{
	GetRuleType() string
	Apply(eventsMap map[string][]*Event)
}

type BonusPointsRule struct {
	ruleType		string
	event             string
	minimunOcurrences int
	bonusPoints       int
	player            string
	afterTime         string
}

type MatchRule struct {
	ruleType		string
	event  string
	points int
}

type ParticularRule struct {
	ruleType		string
	event             string
	minimunOcurrences int
	valueFactor       int
	player            string
	afterTime         string
}

func (particularRule ParticularRule) GetRuleType() string {
	return particularRule.ruleType
}

func (bonusPointsRule BonusPointsRule) GetRuleType() string {
	return bonusPointsRule.ruleType
}

func (matchRule MatchRule) GetRuleType() string {
	return matchRule.ruleType
}

func RuleFactory(ruleType string, event string, points int, distance string, player string, minimunOcurrences int, valueFactor int, afterTime string) Rule {
	if ruleType == "score" || ruleType == "single" {
		return BonusPointsRule{
			ruleType:			"bonusPoints",
			event:             event,
			minimunOcurrences: minimunOcurrences,
			bonusPoints:       points,
			player:            player,
			afterTime:         afterTime,
		}
	} else if ruleType == "particular" {
		return NewParticularRule(event, minimunOcurrences, valueFactor, player, afterTime)
	} else {
		return MatchRule{
			ruleType:	"match",
			event:  	event,
			points: 	points,
		}
	}
}

func NewParticularRule(event string, minimunOcurrences int, valueFactor int, player string, afterTime string) ParticularRule {
	return ParticularRule {
		ruleType:			"particular",
		event:             event,
		minimunOcurrences: minimunOcurrences,
		valueFactor:       valueFactor,
		player:            player,
		afterTime:         afterTime,
	}
}


func (particularRule ParticularRule) Apply (eventsMap map[string][]*Event) {
	
	events, hasKey := eventsMap["score"]
	if !hasKey {
		return
	}

	for _, event := range events {
		//TODO: tiene que cumplir condicion
		event.SetValueFactor(particularRule.valueFactor)
	}
}

func (bonuPointsRule BonusPointsRule) Apply (eventsMap map[string][]*Event) {

	
	events, hasKey := eventsMap[bonuPointsRule.event]
	if !hasKey {
		return
	}

	for _, event := range events {
		//TODO: tiene que cumplir condicion
		event.SetBonusPoints(bonuPointsRule.bonusPoints)
	}
}

func (matchRule MatchRule) Apply (eventsMap map[string][]*Event) {

	
	events, hasKey := eventsMap["win"]
	if !hasKey {
		return
	}

	events[0].SetPoints(matchRule.points)
}