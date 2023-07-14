package utils

import (
	"encoding/json"
	"fifa-review/schemas"
	"fmt"
)


type JsonParser struct {}

func (j *JsonParser) ParseMatch(filepath string) (schemas.MatchSchema /*despues se devuelve el struct bien*/, error) {

	var match schemas.MatchSchema
	fileReader := FileReader{}
	object, fileReaderError := fileReader.ReadFile(filepath)
	
	if fileReaderError != nil {
		return match, fileReaderError
	}

	jsonParseError := json.Unmarshal(object, &match)

	if jsonParseError != nil {
		fmt.Println("An error occured while parsing file ", filepath)
	}

	return match, jsonParseError
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