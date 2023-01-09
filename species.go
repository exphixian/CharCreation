package main

import (
	"fmt"
)

type speciesDetails struct {
	species    string
	size       string
	speed      int
	languages  string
	subspecies string
	statmods   map[string]int
	other      []string
}

func speciesMGMT() speciesDetails {
	character := stringInput("species")

	var subMod map[string]int
	var subMisc []string

	switch {
	case character == "Dwarf":
		sub := ""
		fmt.Println("What subspecies are you playing as? (Hill or Mountain)")
		fmt.Scanln(&sub)

		switch {
		case sub == "Hill":
			subMod = map[string]int{"wisdom": 1, "hp": 1}
		case sub == "Mountain":
			subMod = map[string]int{"strength": 2, "hp": 1}
			subMisc = append(subMisc, "Light Armor Proficiency", "Medium Armor Proficiency")
		default:
			fmt.Println("That subspecies is not supported by this script. Subspecies related character adjustments will not be added.\n")
		}

		speciesInfo := speciesDetails{
			species:    character,
			size:       "Medium",
			speed:      25,
			languages:  "Common, Dwarvish",
			subspecies: sub,
			//HP adj per level
			statmods: map[string]int{"constitution": 2},
			other:    []string{"Darkvision", "Dwarven Resilience", "Dwarven Combat Training (proficiency: battleaxe, handaxe, light hammer, warhammer)", "Tool Proficiency (smith's tools || brewer's supplies || mason's tools)", "Stonecutting"},
		}

		for k, v := range subMod {
			speciesInfo.statmods[k] = v
		}

		for i := 0; i < len(subMisc); i++ {
			speciesInfo.other = append(speciesInfo.other, subMisc[i])
		}

		return speciesInfo

	case character == "Elf":
		sub := ""
		fmt.Println("What subspecies are you playing as? (High, Wood or Dark)")
		fmt.Scanln(&sub)

		switch {
		case sub == "High":
			subMod = map[string]int{"intelligence": 1}
			subMisc = append(subMisc, "Longsword Proficiency", "Shortsword Proficiency", "Longbow Proficiency", "Shortbow Proficiency", "Cantrip 1", "Extra Language")
		case sub == "Wood":
			subMod = map[string]int{"wisdom": 1}
			subMisc = append(subMisc, "Longsword Proficiency", "Shortsword Proficiency", "Longbow Proficiency", "Shortbow Proficiency", "Speed: 35", "Mask of the Wild")
		case sub == "Dark":
			subMod = map[string]int{"charisma": 1}
			subMisc = append(subMisc, "Superior Darkvision", "Sunlight Sensitivity", "Drow Magic", "Rapier Proficiency", "Shortsword Proficiency", "Hand Crossbow Proficiency")
		default:
			fmt.Println("That subspecies is not supported by this script. Subspecies related character adjustments will not be added.\n")
		}

		speciesInfo := speciesDetails{
			species:    character,
			size:       "Medium",
			speed:      30,
			languages:  "Common, Elvish",
			subspecies: sub,
			statmods:   map[string]int{"dexterity": 2},
			//additional language, speed adj
			other: []string{"Darkvision", "Keen Senses", "Fey Ancestry", "Trance"},
		}

		for k, v := range subMod {
			speciesInfo.statmods[k] = v
		}

		for i := 0; i < len(subMisc); i++ {
			speciesInfo.other = append(speciesInfo.other, subMisc[i])
		}

		return speciesInfo

	case character == "Halfling":
		sub := ""
		fmt.Println("What subspecies are you playing as? (Lightfoot or Stout)")
		fmt.Scanln(&sub)

		switch {
		case sub == "Lightfoot":
			subMod = map[string]int{"charisma": 1}
			subMisc = append(subMisc, "Naturally Stealthy")
		case sub == "Stout":
			subMod = map[string]int{"constitution": 1}
			subMisc = append(subMisc, "Stout Resilience")
		default:
			fmt.Println("That subspecies is not supported by this script. Subspecies related character adjustments will not be added.\n")
		}

		speciesInfo := speciesDetails{
			species:    character,
			size:       "Small",
			speed:      25,
			languages:  "Common, Halfling",
			subspecies: sub,
			statmods:   map[string]int{"dexterity": 2},
			other:      []string{"Brave", "Lucky", "Halfling Nimbleness"},
		}

		for k, v := range subMod {
			speciesInfo.statmods[k] = v
		}

		for i := 0; i < len(subMisc); i++ {
			speciesInfo.other = append(speciesInfo.other, subMisc[i])
		}

		return speciesInfo

	case character == "Human":
		fmt.Printf("No %v subspecies supported at this time.", character)

		speciesInfo := speciesDetails{
			species:    character,
			size:       "Medium",
			speed:      30,
			languages:  "Common",
			subspecies: "",
			statmods:   map[string]int{"strength": 1, "constitution": 1, "dexterity": 1, "intelligence": 1, "wisdom": 1, "charisma": 1},
			other:      []string{"Additional Language"},
		}

		return speciesInfo

	case character == "Dragonborn":
		fmt.Printf("No %v subspecies supported at this time.", character)

		speciesInfo := speciesDetails{
			species:    "Dragonborn",
			size:       "Medium",
			speed:      30,
			languages:  "Common, Draconic",
			subspecies: "",
			statmods:   map[string]int{"strength": 2, "charisma": 1},
			other:      []string{"Draconic Ancestry", "Breath Weapon", "Damage Resistance: Draconic"},
		}

		return speciesInfo

	case character == "Gnome":
		sub := ""
		fmt.Println("What subspecies are you playing as? (Forest or Rock)")
		fmt.Scanln(&sub)

		switch {
		case sub == "Forest":
			subMod = map[string]int{"dexterity": 1}
			subMisc = []string{"Natural Illusionist", "Speak with Small Beasts"}
		case sub == "Rock":
			subMod = map[string]int{"constitution": 1}
			subMisc = []string{"Artificer's Lore", "Tinker"}
		default:
			fmt.Println("That subspecies is not supported by this script. Subspecies related character adjustments will not be added.\n")
		}

		speciesInfo := speciesDetails{
			species:    character,
			size:       "Small",
			speed:      25,
			languages:  "Common, Gnomish",
			subspecies: sub,
			statmods:   map[string]int{"intelligence": 2},
			other:      []string{"Darkvision", "Gnome Cunning"},
		}

		for k, v := range subMod {
			speciesInfo.statmods[k] = v
		}

		for i := 0; i < len(subMisc); i++ {
			speciesInfo.other = append(speciesInfo.other, subMisc[i])
		}

		return speciesInfo

	case character == "Half-Elf":
		fmt.Printf("No %v subspecies supported at this time.", character)

		speciesInfo := speciesDetails{
			species:    character,
			size:       "Medium",
			speed:      30,
			languages:  "Common, Elvish",
			subspecies: "",
			//2 ability points
			statmods: map[string]int{"charisma": 2},
			other:    []string{"Darkvision", "Fey Ancestry", "Skill Versitility", "Extra Language"},
		}

		return speciesInfo

	case character == "Half-Orc":
		fmt.Printf("No %v subspecies supported at this time.", character)

		speciesInfo := speciesDetails{
			species:    character,
			size:       "Medium",
			speed:      30,
			languages:  "Common, Orc",
			subspecies: "",
			statmods:   map[string]int{"strength": 2, "constitution": 1},
			other:      []string{"Darkvision", "Menacing", "Relentless Endurance", "Savage Attacks"},
		}

		return speciesInfo

	case character == "Tiefling":
		fmt.Printf("No %v subspecies supported at this time.", character)

		speciesInfo := speciesDetails{
			species:    character,
			size:       "Medium",
			speed:      30,
			languages:  "Common, Infernal",
			subspecies: "",
			statmods:   map[string]int{"charisma": 2, "intelligence": 1},
			other:      []string{"Darkvision", "Hellish Resistance", "Infernal Legacy"},
		}

		return speciesInfo

	default:
		fmt.Println("That species is not supported by this script. Species related character adjustments will not be added.\n")
		break
	}

	speciesInfo := speciesDetails{
		species: character,
	}

	return speciesInfo
}
