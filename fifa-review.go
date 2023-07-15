package main

import (
	"fifa-review/entities"
	"fifa-review/utils"
	"flag"
	"fmt"
)

func prepareData() ([]entities.Match, []entities.Rule, error) {
	var match_file_paths utils.FlagsArray
	flag.Var(&match_file_paths, "match", "matches files")
	rules_file_path := flag.String("rules", "", "rules file")
	flag.Parse()

	if len(match_file_paths) == 0 {
		fmt.Println("Please provide at least one match file path")
		return nil, nil, nil
	}

	parser := utils.JsonParser{}

	var rules []entities.Rule

	if len(*rules_file_path) != 0 {
		var err error
		rules, err = parser.ParseRules(*rules_file_path)

		if err != nil {
			return nil, nil, err
		}
	}

	matches := []entities.Match{}

	for _, matchFilePath := range match_file_paths {

		match, err := parser.ParseMatch(matchFilePath)

		if err != nil {
			return nil, nil, err
		}

		matches = append(matches, match)
	}

	return matches, rules, nil

}

func main() {

	matches, rules, err := prepareData()

	if err != nil {
		return
	}

	fmt.Println(matches)
	fmt.Println(rules)
}
