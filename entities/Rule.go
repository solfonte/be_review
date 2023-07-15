package entities

type Rule interface{}

type BonusPointsRule struct {
	event             string
	minimunOcurrences int
	bonusPoints       int
	player            string
	afterTime         string
}

type MatchRule struct {
	event  string
	points int
}

type ParticularRule struct {
	event             string
	minimunOcurrences int
	valueFactor       int
	player            string
	afterTime         string
}

func RuleFactory(ruleType string, event string, points int, distance string, player string, minimunOcurrences int, valueFactor int, afterTime string) Rule {
	if ruleType == "score" || ruleType == "single" {
		return BonusPointsRule{
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
			event:  event,
			points: points,
		}
	}
}

func NewParticularRule(event string, minimunOcurrences int, valueFactor int, player string, afterTime string) ParticularRule {
	return ParticularRule {
		event:             event,
		minimunOcurrences: minimunOcurrences,
		valueFactor:       valueFactor,
		player:            player,
		afterTime:         afterTime,
	}
}


func (particularRule *ParticularRule) Apply (event map[string][]*Event) {
	
	events, hasKey := event["score"]
	if !hasKey {
		return
	}

	for _, event := range events {
		//TODO: tiene que cumplir condicion
		event.SetValueFactor(particularRule.valueFactor)
	}
}