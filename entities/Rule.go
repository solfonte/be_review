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

type Particular struct {
	event             string
	minimunOcurrences int
	points            int
	player            string
	afterTime         string
}

func NewRule(ruleType string, event string, points int, distance string, player string, minimunOcurrences int, valueFactor string, afterTime string) Rule {
	if ruleType == "score" || ruleType == "single" {
		return BonusPointsRule{
			event:             event,
			minimunOcurrences: minimunOcurrences,
			bonusPoints:       points,
			player:            player,
			afterTime:         afterTime,
		}
	} else if ruleType == "particular" {
		return Particular{
			event:             event,
			minimunOcurrences: minimunOcurrences,
			points:            points,
			player:            player,
			afterTime:         afterTime,
		}
	} else {
		return MatchRule{
			event:  event,
			points: points,
		}
	}
}
