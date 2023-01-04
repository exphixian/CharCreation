package main

import(
	"fmt"
	"math/rand"
)

func characterAdj(level int, species string, subspecies string, job string, stats map[string]int) map[string]int{
	//add species support
	///*
	switch {
	case species == "Dwarf":
		stats["constitution"] += 2
		switch {
		case subspecies == "Hill":
			stats["wisdom"] += 1
		case subspecies == "Mountain":
			stats["strength"] += 2
		}

	case species == "Elf":
		stats["dexterity"] += 2
		switch {
		case subspecies == "High":
			stats["intelligence"] += 1
		case subspecies == "Wood":
			stats["wisdom"] += 1
		case subspecies == "Dark":
			stats["charisma"] += 1
		}

	case species == "Halfling":
		stats["dexterity"] += 2
		switch {
		case subspecies == "Lightfoot":
			stats["charisma"] += 1
		case subspecies == "Stout":
			stats["constitution"] += 1
		}

	case species == "Human":
		for k, v := range stats {
			v += 1
		}

	case species == "Dragonborn":
		stats["strength"] += 2
		stats["charisma"] += 1

	case species == "Gnome":
		stats["intelligence"] += 2
		switch {
		case subspecies == "Forest":
			stats["dexterity"] += 1
		case subspecies == "Rock":
			stats["constitution"] += 1
		}

	case species == "Half-Elf":
		stats["charisma"] += 2

	case species == "Half-Orc":
		stats["strength"] += 2
		stats["constitution"] += 1

	case species == "Tiefling":
		stats["charisma"] += 2
		stats["intelligence"] += 1
	}
	//*/

	abilityPoints := (Level / 4) * 2

	if species == "Half-Elf" {
		abilityPoints += 2
	}

	if level > 3 {
		if Level >= 19 {
			abilityPoints += 2
		}

		if job == "fighter" {
			if Level >= 14 {
				abilityPoints += 4
			} else if Level >= 6 {
				abilityPoints += 2
			}

		} else if job == "rogue" {
			if Level >= 10 {
				abilityPoints += 2
			}
		}

	}

	fmt.Printf("Where do you want to allocate your %d ability points?\n", abilityPoints)
		for i := 0; i < abilityPoints; i++ {
			var input string
			fmt.Scanln(&input)
			stats[input] += 1
		}
	
}

func randomizedStats(level int, species string, subspecies string, job string) map[string]int, map[string]int {
	rand.Seed(time.Now().UnixNano())
	stats := map[string]int{
		"strength" : rand.Intn(15) + 4,
		"constitution" : rand.Intn(15) + 4,
		"dexterity" : rand.Intn(15) + 4,
		"intelligence" :rand.Intn(15) + 4,
		"wisdom": rand.Intn(15) + 4,
		"charisma": rand.Intn(15) + 4,
	}

	//need to add in a confirm and reroll option

	stats = characterAdj(level, species, stats)
	

	mods := map[string]int{
		"STR" : (stats["strength"]-10)/2,
		"CON" : (stats["constitution"]-10)/2,
		"DEX" : (stats["dexterity"]-10)/2,
		"INT" : (stats["intelligence"]-10)/2,
		"WIS": (stats["wisdom"]-10)/2,
		"CHA": (stats["charisma"]-10)/2,
	}

	//Need to figure out the negative mod mismatch
	/*
	for k, v := range mods {
		if v 
	}
	*/

	return stats, mods
}

func manualStats() map[string]int, map[string]int {
	
	catagories := []string{"strength", "constitution", "dexterity", "intelligence", "wisdom", "charisma"}
	stats := map[string]int{}

	fmt.Println("Please insert your stats.")
	for i := 0; i < len(catagories); i++ {
		var input int
		fmt.Printf("\n%v:", catagories[i])
		fmt.Scan(&input)

		stats = append(stats, catagories[i]: input)
	}

		//need to add in a confirm and reroll option

	stats = characterAdj(level, species, stats)
	

	mods := map[string]int{
		"STR" : (stats["strength"]-10)/2,
		"CON" : (stats["constitution"]-10)/2,
		"DEX" : (stats["dexterity"]-10)/2,
		"INT" : (stats["intelligence"]-10)/2,
		"WIS": (stats["wisdom"]-10)/2,
		"CHA": (stats["charisma"]-10)/2,
	}

	//Need to figure out the negative mod mismatch
	/*
	for k, v := range mods {
		if v 
	}
	*/

	return stats, mods
}