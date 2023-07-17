package main

import (
	"fifa-review/entities"
	"fifa-review/utils"
	"flag"
	"fmt"
	"sort"
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

func defineTableOrder(results map[string]*entities.Result) map[string]entities.Result {
	points := make(map[string]int)
	orderedResults := make(map[string]entities.Result)

	for team, result := range results {
		totalPoints := result.Total_points + result.Bonus_points
		points[team] = totalPoints
	}

	teams := make([]string, 0, len(points))
	for key := range points {
		teams = append(teams, key)
	}

	sort.Strings(teams)

	for _, team := range teams {
		orderedResults[team] = *results[team]
	}

	return orderedResults
}

func prepareRules(rulesFilePath string) ([]entities.MatchRule, []entities.BonusPointsRule, []entities.ParticularRule, error) {
	parser := utils.JsonParser{}

	matchRules, bonusPointsRules, particularRules, err := parser.ParseRules(rulesFilePath)

	if err != nil {
		return nil, nil, nil, err
	}

	return matchRules, bonusPointsRules, particularRules, nil
}

func SaveResults(orderedCountriesByPoints map[string]entities.Result) {
	parser := utils.JsonParser{}
	parser.SaveResults(orderedCountriesByPoints)
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
	if len(rulesFile) > 0 {
		matchRules, bonusPointsRules, particularRules, err = prepareRules(rulesFile)

	}

	if err != nil {
		return
	}

	resultsPerCountry := make(map[string]*entities.Result)

	for _, match := range matches {

		for _, rule := range particularRules {
			match.ApplySpecialRule(rule)
		}
		match.DefineFinalResult()

		for _, rule := range matchRules {
			match.ApplyRuleToWinner(rule)
		}

		match.AssignPointsAccordingFinalResult()

		for _, rule := range bonusPointsRules {
			match.ApplyBonusPointsRule(rule)
		}

		for team, result := range match.GetResults() {

			r, hasResult := resultsPerCountry[team]
			if hasResult {
				r.Total_points += result.Total_points
				r.Bonus_points += result.Bonus_points
				r.Played_matches_amount += 1
				r.Scores_in_favor_amount += result.Scores_in_favor_amount
			} else {
				resultsPerCountry[team] = &entities.Result{
					Total_points:           result.Total_points,
					Bonus_points:           result.Bonus_points,
					Played_matches_amount:  result.Played_matches_amount,
					Scores_in_favor_amount: result.Scores_in_favor_amount,
				}
			}

		}

	}

	orderedCountriesByPoints := defineTableOrder(resultsPerCountry)
	SaveResults(orderedCountriesByPoints)

}
