package utils

import (
	"encoding/json"
)


type JsonParser struct {}

func (j *JsonParser) ParseMatch(filepath string) (MatchSchema /*despues se devuelve el struct bien*/, error) {

	var match MatchSchema
	fileReader := FileReader{}
	object, fileReaderError := fileReader.ReadFile(filepath)
	
	if fileReaderError != nil {
		return match, fileReaderError
	}

	jsonParseError := json.Unmarshal(object, &match)

	return match, jsonParseError
}