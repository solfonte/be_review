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

	if len(match_file_paths) == 0 {
		fmt.Println("Please provide at least one match file path")
	}

	if len(*rules_file_path) != 0 {
		//do something
	}
}