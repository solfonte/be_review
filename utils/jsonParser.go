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

		parsedEvents = append(parsedEvents, entities.NewEvent(event.Event, event.Time, &eventDetails))
	}

	return parsedEvents

}

func (j *JsonParser) ParseMatch(filepath string) (*entities.Match, error) {

	var match schemas.MatchSchema
	fileReader := FileReader{}
	object, fileReaderError := fileReader.ReadFile(filepath)

	if fileReaderError != nil {
		return nil, fileReaderError
	}

	jsonParseError := json.Unmarshal(object, &match)

	if jsonParseError != nil {
		fmt.Println("An error occured while parsing file ", filepath)
	}

	parsedAwayEvents := CreateEventsList(match.Away_events)
	parsedHomeEvents := CreateEventsList(match.Home_events)

	parsedMatch := entities.NewMatch(match.Teams.Away, parsedAwayEvents, match.Teams.Home, parsedHomeEvents)

	return parsedMatch, jsonParseError
}

func (j *JsonParser) ParseRules(filepath string) ([]entities.MatchRule, []entities.BonusPointsRule, []entities.ParticularRule, error) {

	var rules []schemas.RuleSchema
	var matchRules []entities.MatchRule
	var bonusPointsRules []entities.BonusPointsRule
	var particularRules []entities.ParticularRule

	fileReader := FileReader{}
	object, fileReaderError := fileReader.ReadFile(filepath)

	if fileReaderError != nil {
		return nil, nil, nil, fileReaderError
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
			valueFactor, _ = strconv.Atoi(rule.Value_factor[1:2])

		}

		if rule.Type == "match" {
			newRule := entities.NewMatchRule(rule.Event, points)
			matchRules = append(matchRules, newRule)

		} else if rule.Type == "single" || rule.Type == "side" {
			newRule := entities.NewBonusPointsRule(rule.Event, rule.Condition.At_least, rule.Bonus_points, rule.Condition.Player, rule.Condition.After_time, rule.Condition.Distance)
			bonusPointsRules = append(bonusPointsRules, newRule)

		} else {
			newRule := entities.NewParticularRule(rule.Event, rule.Condition.At_least, valueFactor, rule.Condition.Player, rule.Condition.After_time, rule.Condition.Distance)
			particularRules = append(particularRules, newRule)

		}
	}

	return matchRules, bonusPointsRules, particularRules, jsonParseError
}

func (j *JsonParser) SaveResults(results map[string]entities.Result) {
	jsonData, err := json.Marshal(results)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}
	fileReader := FileReader{}
	fileReader.WriteFile("results.json", jsonData)

}
