package main

import (
	"fmt"
	"os"
	"strings"
)

var Character char

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

func main() {
	Character.name = stringInput("name")

	fmt.Printf("\nWhat level is your Character?\n")
	fmt.Scan(&Character.level)
	if Character.level > 20 {
		fmt.Println("This sheet does not support legendary Characters.")
		os.Exit(1)
	}

	speciesInfo := speciesMGMT()
	Character.species = speciesInfo.species
	Character.speed = speciesInfo.speed
	Character.languages = speciesInfo.languages

	//Need to support multiclassing.
	jobInfo := jobMGMT(Character.level)
	Character.job = jobInfo.job

	//combining attributes that pull from species and job
	for i := 0; i < len(speciesInfo.other); i++ {
		Character.features = append(Character.features, speciesInfo.other[i])
	}

	for i := 0; i < len(jobInfo.features); i++ {
		Character.features = append(Character.features, jobInfo.features[i])
	}

	fmt.Printf("\n%v is a level %v %v %v.\n\n", Character.name, Character.level, speciesInfo.species, jobInfo.job)
	sleep()

	var generate string

	fmt.Println("Yes or No: do you want your stats to be randomly generated?")

	fmt.Scan(&generate)
	generate = strings.ToLower(generate)
	if generate == "yes" || generate == "y" {
		Character.abilityScores, Character.modifiers = randomizedStats(Character.level, speciesInfo.species, speciesInfo.subspecies, speciesInfo.statmods, jobInfo.job)
	} else {
		Character.abilityScores, Character.modifiers = manualStats(Character.level, speciesInfo.species, speciesInfo.subspecies, speciesInfo.statmods, jobInfo.job)
	}

	//Need to support job/species HP adjustment
	Character.hp = (Character.level * Character.modifiers["CON"]) + jobInfo.hp

	sleep()

	fmt.Printf("\n%+v\n Level: %+v\n Ability Scores: %+v\n Modifiers: %+v\n HP: %+v\n Species: %+v\n Job: %+v\n Speed: %+v\n Languages: %+v\n Feats: %+v\n",
		Character.name, Character.level, Character.abilityScores, Character.modifiers, Character.hp, Character.species, Character.job, Character.speed,
		Character.languages, Character.features)
}
