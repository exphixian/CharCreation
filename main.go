package main

import (
	"bufio"
	"fmt"
	"os"
)

// placing all character sheet information in one place to allow for easier export and further development
type char struct {
	name          string
	level         int
	abilityScores map[string]int
	modifiers     map[string]int
	hp            int
	species       string
	job           string
	speed         int
	languages     string
	features      []string
}

// reads string input from the user and returns it.
func stringInput(req string) string {
	stringInput := bufio.NewScanner(os.Stdin)
	fmt.Printf("\nWhat is your character's %s?\n", req)
	stringInput.Scan()
	variable := stringInput.Text()
	return variable
}

func main() {
	var character char
	character.name = stringInput("name")

	character.level = 0
	fmt.Printf("\nWhat level is your character?\n")
	fmt.Scan(&character.level)
	if character.level > 20 {
		fmt.Println("This sheet does not support legendary characters.")
		os.Exit(1)
	}

	speciesInfo := speciesMGMT()
	character.species = speciesInfo.species
	character.speed = speciesInfo.speed
	character.languages = speciesInfo.languages

	//Need to support multiclassing.
	jobInfo := jobMGMT(character.level)
	character.job = jobInfo.job

	//combining attributes that pull from species and job
	for i := 0; i < len(speciesInfo.other); i++ {
		character.features = append(character.features, speciesInfo.other[i])
	}

	for i := 0; i < len(jobInfo.features); i++ {
		character.features = append(character.features, jobInfo.features[i])
	}

	fmt.Printf("\n%v is a level %v %v %v.\n\n", character.name, character.level, speciesInfo.species, jobInfo.job)
	sleep()

	var generate string

	fmt.Println("Yes or No: do you want your stats to be randomly generated?")

	fmt.Scan(&generate)
	if generate == "yes" || generate == "Yes" {
		character.abilityScores, character.modifiers = randomizedStats(character.level, speciesInfo.species, speciesInfo.subspecies, speciesInfo.statmods, jobInfo.job)
	} else {
		character.abilityScores, character.modifiers = manualStats(character.level, speciesInfo.species, speciesInfo.subspecies, speciesInfo.statmods, jobInfo.job)
	}

	//Need to support job/species HP adjustment
	character.hp = (character.level * character.modifiers["CON"]) + jobInfo.hp

	sleep()

	fmt.Printf("\n%+v\n Level: %+v\n Ability Scores: %+v\n Modifiers: %+v\n HP: %+v\n Species: %+v\n Job: %+v\n Speed: %+v\n Languages: %+v\n Feats: %+v\n",
		character.name, character.level, character.abilityScores, character.modifiers, character.hp, character.species, character.job, character.speed,
		character.languages, character.features)
}
