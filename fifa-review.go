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

func prepareRules(rulesFilePath string) ([]entities.MatchRule,[]entities.BonusPointsRule,[]entities.ParticularRule, error) {
	parser := utils.JsonParser{}

	matchRules, bonusPointsRules, particularRules, err := parser.ParseRules(rulesFilePath)

	if err != nil {
		return nil, nil, nil, err
	}

	return matchRules, bonusPointsRules, particularRules, nil
}


func main() {

	rulesFile, matchesFiles := getFlagFiles()
	matches, err := prepareMatches(matchesFiles)
	if err != nil {
		return
	}

	var matchRules []entities.MatchRule
	var bonusPointsRules []entities.BonusPointsRule
	var particularRules []entities.ParticularRule
	if len(rulesFile) > 0{
		matchRules, bonusPointsRules, particularRules, err = prepareRules(rulesFile)
		
	}

	if err != nil {
		return
	}

	fmt.Println("a", matchRules)
	fmt.Println("b", bonusPointsRules)
	fmt.Println("c", particularRules)

	var resultsPerCountry map[string]entities.Result

	for _, match := range matches {
		
		for _, rule := range particularRules {
			match.ApplySpecialRule(rule)
		}
		match.DefineWinner()

		fmt.Println(match)

		for _, rule := range bonusPointsRules {
			match.ApplyBonusPointsRule(rule)
		}

		for _, rule := range matchRules {
			match.ApplyRuleToWinner(rule)
		}

		match.AssignPointsToWinner()

		fmt.Println(resultsPerCountry)
		fmt.Println(match.GetResults())
	}
}
