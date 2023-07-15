package utils

import (
	"encoding/json"
	"fifa-review/entities"
	"fifa-review/schemas"
	"fmt"
	"strconv"
)

type JsonParser struct{}

func CreateEventsList(events []schemas.EventSchema) []*entities.Event {

	parsedEvents := []*entities.Event{}

	for _, event := range events {

		eventDetails := make(map[string]string)

		if len(event.Distance) > 0 {
			eventDetails["distance"] = event.Distance
		}

		if len(event.Player) > 0 {
			eventDetails["player"] = event.Player
		}

		if len(event.Player) > 0 {
			eventDetails["obs"] = event.Obs
		}

		parsedEvents = append(parsedEvents, entities.NewEvent(event.Event, event.Time, eventDetails))
	}

	return parsedEvents

}

func (j *JsonParser) ParseMatch(filepath string) (entities.Match, error) {

	var match schemas.MatchSchema
	fileReader := FileReader{}
	object, fileReaderError := fileReader.ReadFile(filepath)

	if fileReaderError != nil {
		return entities.Match{}, fileReaderError
	}

	jsonParseError := json.Unmarshal(object, &match)

	if jsonParseError != nil {
		fmt.Println("An error occured while parsing file ", filepath)
	}

	parsedAwayEvents := CreateEventsList(match.Away_events)
	parsedHomeEvents := CreateEventsList(match.Home_events)

	parsedMatch := *entities.NewMatch(match.Teams.Away, parsedAwayEvents, match.Teams.Home, parsedHomeEvents)

	return parsedMatch, jsonParseError
}

func (j *JsonParser) ParseRules(filepath string) (map[string][]entities.Rule, error) {

	var rules []schemas.RuleSchema
	parsedRules := make(map[string][]entities.Rule)
	fileReader := FileReader{}
	object, fileReaderError := fileReader.ReadFile(filepath)

	if fileReaderError != nil {
		return parsedRules, fileReaderError
	}

	jsonParseError := json.Unmarshal(object, &rules)

	if jsonParseError != nil {
		fmt.Println("An error occured while parsing file ", filepath)
	}

	for _, rule := range rules {

		var points int
		if rule.Points > 0 {
			points = rule.Points
		} else {
			points = rule.Bonus_points
		}

		var valueFactor int
		if len(rule.Value_factor) > 0 {
			fmt.Println(rule.Value_factor[1:2])
			valueFactor, _ = strconv.Atoi(rule.Value_factor[1:2])

		} 
		newRule := entities.RuleFactory(rule.Type, rule.Event, points, rule.Condition.Distance, rule.Condition.Player, rule.Condition.At_least, valueFactor, rule.Condition.After_time)

		_, hasKey := parsedRules[rule.Type]

		if hasKey {
			parsedRules[rule.Type] = append(parsedRules[rule.Type], newRule)
		} else {
			parsedRules[rule.Type] = []entities.Rule {newRule}
		}
	}

	return parsedRules, jsonParseError
}
