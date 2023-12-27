package main

import (
	"fmt"
	"math/rand"
)

func characterAdj(level int, species string, subspecies string, speciesmods map[string]int, job string, stats map[string]int) map[string]int {
	for k, v := range speciesmods {
		stats[k] += v
	}

	abilityPoints := (level / 4) * 2

	if species == "Half-Elf" {
		abilityPoints += 2
	}

	if level >= 19 {
		abilityPoints += 2
	}

	if job == "fighter" {
		if level >= 14 {
			abilityPoints += 4
		} else if level >= 6 {
			abilityPoints += 2
		}

	} else if job == "rogue" {
		if level >= 10 {
			abilityPoints += 2
		}
	}

	if abilityPoints != 0 {
		fmt.Printf("Where do you want to allocate your %d ability points?\n", abilityPoints)
		for i := 0; i < abilityPoints; i++ {
			var input string
			fmt.Scanln(&input)
			stats[input] += 1
		}
	}

	return stats

}

// for randomly generating and adjusting stats
// Need to add in a confirm and reroll option
func randomizedStats(level int, species string, subspecies string, speciesmods map[string]int, job string) (map[string]int, map[string]int) {

	//randomizing stats - currently set to roll between 4 & 18
	stats := map[string]int{
		"strength":     rand.Intn(15) + 4,
		"constitution": rand.Intn(15) + 4,
		"dexterity":    rand.Intn(15) + 4,
		"intelligence": rand.Intn(15) + 4,
		"wisdom":       rand.Intn(15) + 4,
		"charisma":     rand.Intn(15) + 4,
	}

	fmt.Printf("\n\nYour base stats are: \n%+v\n", stats)

	//Stat adjustment based on character features (level, species, job, etc)
	stats = characterAdj(level, species, subspecies, speciesmods, job, stats)

	// defining modifiers for skill checks.
	// Need to figure out the negative mod mismatch
	mods := map[string]int{
		"STR": (stats["strength"] - 10) / 2,
		"CON": (stats["constitution"] - 10) / 2,
		"DEX": (stats["dexterity"] - 10) / 2,
		"INT": (stats["intelligence"] - 10) / 2,
		"WIS": (stats["wisdom"] - 10) / 2,
		"CHA": (stats["charisma"] - 10) / 2,
	}
	return stats, mods
}

// for randomly generating and adjusting stats
// Need to add in confirmation query
func manualStats(level int, species string, subspecies string, speciesmods map[string]int, job string) (map[string]int, map[string]int) {

	catagories := []string{"strength", "constitution", "dexterity", "intelligence", "wisdom", "charisma"}
	stats := map[string]int{}

	fmt.Println("Please insert your stats.")
	for i := 0; i < len(catagories); i++ {
		var input int
		fmt.Printf("\n%v:", catagories[i])
		fmt.Scan(&input)

		stats[catagories[i]] = input
	}

	//Stat adjustment based on character features (level, species, job, etc)
	stats = characterAdj(level, species, subspecies, speciesmods, job, stats)

	//defining modifiers for skill checks.
	//Need to figure out the negative mod mismatch
	mods := map[string]int{
		"STR": (stats["strength"] - 10) / 2,
		"CON": (stats["constitution"] - 10) / 2,
		"DEX": (stats["dexterity"] - 10) / 2,
		"INT": (stats["intelligence"] - 10) / 2,
		"WIS": (stats["wisdom"] - 10) / 2,
		"CHA": (stats["charisma"] - 10) / 2,
	}

	return stats, mods
}
