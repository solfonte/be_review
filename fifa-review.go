package main

import (
	"fifa-review/entities"
	"fifa-review/utils"
	"flag"
	"fmt"
)

func getFlagFiles() (string, utils.FlagsArray) {
	rulesFilePath := flag.String("rules", "", "rules file")
	var matchFilePaths utils.FlagsArray
	flag.Var(&matchFilePaths, "match", "matches files")
	flag.Parse()

	return *rulesFilePath, matchFilePaths
}

func prepareMatches(matchFilePaths utils.FlagsArray) ([]*entities.Match, error) {
	

	if len(matchFilePaths) == 0 {
		fmt.Println("Please provide at least one match file path")
		return nil, nil
	}

	parser := utils.JsonParser{}
	matches := []*entities.Match{}

	for _, matchFilePath := range matchFilePaths {

		match, err := parser.ParseMatch(matchFilePath)

		if err != nil {
			return nil, err
		}

		matches = append(matches, match)
	}

	return matches, nil
}


func prepareRules(rulesFilePath string) (map[string][]entities.Rule, error){
	parser := utils.JsonParser{}

	rules,  err := parser.ParseRules(rulesFilePath)

	if err != nil {
		return nil, err
	}

	return rules, nil
}


func main() {

	rulesFile, matchesFiles := getFlagFiles()
	matches, err := prepareMatches(matchesFiles)
	if err != nil {
		return
	}

	var rules map[string][]entities.Rule
	if len(rulesFile) > 0{
		rules, err = prepareRules(rulesFile)
		
	}

	if err != nil {
		return
	}

	fmt.Println(rules)
	fmt.Println(matches)

	var resultsPerCountry map[string]entities.Result

	for _, match := range matches {
		
		for _, rule := range rules["particular"] {
			match.ApplyRules(rule)
		}
		match.DefineWinner()

		 for _, rule := range rules["bonusPoints"] {
			match.ApplyRules(rule)
		}

		_, hasWinnerRule := rules["match"]
		if hasWinnerRule {
			match.ApplyRules(rules["match"][0])
		}

		fmt.Println(resultsPerCountry)
		fmt.Println(match.GetResults())
	}
}
