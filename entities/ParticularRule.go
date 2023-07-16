package entities

import ("fmt")

type ParticularRule struct {
	ruleType		string
	event             string
	minimunOcurrences int
	valueFactor       int
	player            string
	afterTime         string
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
		fmt.Println("vvvv", particularRule.valueFactor)
		event.SetValueFactor(particularRule.valueFactor)
	}
}
