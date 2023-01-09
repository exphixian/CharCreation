package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
type char struct {
}

*/

func stringInput(req string) string {
	//reads string input from the user and returns it.
	stringInput := bufio.NewScanner(os.Stdin)
	fmt.Printf("\nWhat is your character's %s?\n", req)
	stringInput.Scan()
	variable := stringInput.Text()
	return variable
}

func main() {
	name := stringInput("name")

	level := 0
	fmt.Printf("What level is your character?\n")
	fmt.Scan(&level)
	if level > 20 {
		fmt.Println("This sheet does not support legendary characters at this time.")
		os.Exit(1)
	}

	speciesInfo := speciesMGMT()
	//Need to support multiclassing.
	jobInfo := jobMGMT(level)

	fmt.Printf("%v is a level %v %v %v.\n\n", name, level, speciesInfo.species, jobInfo.job)

	sleep()

	var generate string
	var stats map[string]int
	var mods map[string]int

	fmt.Println("Yes or No: do you want your stats to be randomly generated?")

	fmt.Scan(&generate)
	if generate == "yes" || generate == "Yes" {
		stats, mods = randomizedStats(level, speciesInfo.species, speciesInfo.subspecies, speciesInfo.statmods, jobInfo.job)
	} else {
		stats, mods = manualStats(level, speciesInfo.species, speciesInfo.subspecies, speciesInfo.statmods, jobInfo.job)
	}

	fmt.Println(stats, mods)

	sleep()

	fmt.Println(speciesInfo)

}
