package utils

import (
	"encoding/json"
	"fifa-review/schemas"
	"fifa-review/entities"
	"fmt"
)


type JsonParser struct {}

func CreateEventsList(events []schemas.EventSchema) []entities.Event {

	parsedEvents := []entities.Event {}

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

func (j *JsonParser) ParseMatch(filepath string) (entities.Match /*despues se devuelve el struct bien*/, error) {

	var match schemas.MatchSchema
	fileReader := FileReader{}
	object, fileReaderError := fileReader.ReadFile(filepath)
	
	if fileReaderError != nil {
		return entities.Match {}, fileReaderError
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


func (j *JsonParser) ParseRules(filepath string) ([]schemas.RuleSchema /*despues se devuelve el struct bien*/, error) {

	var rules []schemas.RuleSchema
	fileReader := FileReader{}
	object, fileReaderError := fileReader.ReadFile(filepath)
	
	if fileReaderError != nil {
		return rules, fileReaderError
	}

	jsonParseError := json.Unmarshal(object, &rules)

	if jsonParseError != nil {
		fmt.Println("An error occured while parsing file ", filepath)
	}

	return rules, jsonParseError
}