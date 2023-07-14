package main

import (
	"fmt"
	"flag"
	"fifa-review/utils"
)


func main() {
	
	var match_file_paths utils.FlagsArray
	flag.Var(&match_file_paths, "match", "matches files")
	rules_file_path := flag.String("rules", "", "rules file")
    flag.Parse()

	fmt.Println("READ ", *rules_file_path)

	if len(match_file_paths) == 0 {
		fmt.Println("Please provide at least one match file path")
		return
	}

	parser:= utils.JsonParser{}

	if len(*rules_file_path) != 0 {
		rules, err :=parser.ParseRules(*rules_file_path)

		if err != nil {
			return
		}
	}


	for _, matchFilePath := range match_file_paths {

		match, err := parser.ParseMatch(matchFilePath)
		
		if err != nil {
			return
		}
		
		fmt.Println(match)
	}

	

}