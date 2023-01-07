package main

import "fmt"

type jobDetails struct {
	hp               int
	savingThrows     []string
	proficiencies    []string
	proficiencyBonus int
	features         []string
	specialty        map[string]int
	spellslots       map[int]int
}

func spells(level int) map[int]int {
	spells := map[int]map[int]int{1: {1: 2}, 2: {1: 3}, 3: {1: 4, 2: 2}, 4: {1: 4, 2: 3}, 5: {1: 4, 2: 3, 3: 2}, 6: {1: 4, 2: 3, 3: 3},
		7: {1: 4, 2: 3, 3: 3, 4: 1}, 8: {1: 4, 2: 3, 3: 3, 4: 2}, 9: {1: 4, 2: 3, 3: 3, 4: 3, 5: 1}, 10: {1: 4, 2: 3, 3: 3, 4: 3, 5: 2},
		11: {1: 4, 2: 3, 3: 3, 4: 3, 5: 2, 6: 1}, 12: {1: 4, 2: 3, 3: 3, 4: 3, 5: 2, 6: 1}, 13: {1: 4, 2: 3, 3: 3, 4: 3, 5: 2, 6: 1, 7: 1},
		14: {1: 4, 2: 3, 3: 3, 4: 3, 5: 2, 6: 1, 7: 1}, 15: {1: 4, 2: 3, 3: 3, 4: 3, 5: 2, 6: 1, 7: 1, 8: 1},
		16: {1: 4, 2: 3, 3: 3, 4: 3, 5: 2, 6: 1, 7: 1, 8: 1}, 17: {1: 4, 2: 3, 3: 3, 4: 3, 5: 2, 6: 1, 7: 1, 8: 1, 9: 1},
		18: {1: 4, 2: 3, 3: 3, 4: 3, 5: 3, 6: 1, 7: 1, 8: 1, 9: 1}, 19: {1: 4, 2: 3, 3: 3, 4: 3, 5: 3, 6: 2, 7: 1, 8: 1, 9: 1},
		20: {1: 4, 2: 3, 3: 3, 4: 3, 5: 3, 6: 2, 7: 2, 8: 1, 9: 1}}

	s := map[int]int{}
	for k, v := range spells {
		if k == level {
			s = v
			return s
			break
		}
	}

	return s
}

func jobMGMT(level int) jobDetails {
	/*
		need to:
			*support multiclassing
			*support paths
		 	*support skill choices
			*test spells
	*/

	character := stringInput("job")

	switch {
	case character == "Barbarian":
		hitPoints := 12

		for i := 1; i < level; i++ {
			hpRolls := roll(12, 1)
			for i := 0; i < len(hpRolls); i++ {
				hitPoints += hpRolls[i]
			}
		}

		possiblefeatures := map[string]int{"Rage": 1, "Unarmored Defense": 1, "Reckless Attack": 2, "Danger Sense": 2, "Primal Path": 3,
			"Extra Attack": 5, "Fast Movement": 5, "Feral Instinct": 7, "Brutal Critical (1)": 9, "Relentless Rage": 11, "Brutal Critical (2)": 13,
			"Persistent Rage": 15, "Brutal Critical (3)": 17, "Indomitable Might": 18, "Primal Champion": 20}
		feats := []string{}

		for k, v := range possiblefeatures {
			if v <= level {
				feats = append(feats, k)
			}
		}

		intAdj := map[int][]int{1: {2, 2, 2}, 2: {2, 2, 2}, 3: {2, 3, 2}, 4: {2, 3, 2}, 5: {3, 3, 2}, 6: {3, 4, 2}, 7: {3, 4, 2}, 8: {3, 4, 2}, 9: {4, 4, 3},
			10: {4, 4, 3}, 11: {4, 4, 3}, 12: {4, 5, 3}, 13: {5, 5, 3}, 14: {5, 5, 3}, 15: {5, 5, 3}, 16: {5, 5, 4}, 17: {6, 6, 4}, 18: {6, 6, 4}, 19: {6, 6, 4},
			20: {6, 00, 4}}

		jobInfo := jobDetails{
			hp:               hitPoints,
			savingThrows:     []string{"strength", "constitution"},
			proficiencies:    []string{"Light Armor", "Medium Armor", "Shields", "Simple Weapons", "Martial Weapons"},
			proficiencyBonus: intAdj[level][0],
			features:         feats,
			specialty:        map[string]int{"Rages": intAdj[level][1], "Rage Damage": intAdj[level][2]},
		}

		return jobInfo

	case character == "Bard":
		hitPoints := 8

		for i := 1; i < level; i++ {
			hpRolls := roll(8, 1)
			for i := 0; i < len(hpRolls); i++ {
				hitPoints += hpRolls[i]
			}
		}

		possiblefeatures := map[string]int{"Spellcasting": 1, "Bardic Inspiration (d6)": 1, "Jack of All Trades": 2, "Song of Rest (d6)": 2,
			"Bard College": 3, "Expertise #1": 3, "Bardic Inspiration (d8)": 5, "Font of Inspiration": 5, "Counter Charm": 6, "Bard College Feature #1": 6,
			"Song of Rest (d8)": 9, "Bardic Inspiration (d10)": 10, "Expertise #2": 10, "Magical Secrets #1": 10, "Song of Rest (d10)": 13,
			"Magical Secrets #2": 14, "Bardic Inspiration (d12)": 15, "Song of Rest (d12)": 17, "Magical Secrets #3": 18, "Superior Inspiration": 20}

		feats := []string{}

		for k, v := range possiblefeatures {
			if v <= level {
				feats = append(feats, k)
			}
		}

		intAdj := map[int][]int{1: {2, 2}, 2: {2, 2}, 3: {2, 2}, 4: {2, 3}, 5: {3, 3}, 6: {3, 3}, 7: {3, 3}, 8: {3, 3}, 9: {4, 3}, 10: {4, 4},
			11: {4, 4}, 12: {4, 4}, 13: {5, 4}, 14: {5, 4}, 15: {5, 4}, 16: {5, 4}, 17: {6, 4}, 18: {6, 4}, 19: {6, 4}, 20: {6, 4}}

		jobInfo := jobDetails{
			hp:               hitPoints,
			savingThrows:     []string{"dexterity", "charisma"},
			proficiencies:    []string{"Light Armor", "Hand Crossbows", "Longswords", "Simple Weapons", "Rapiers", "Shortswords"},
			proficiencyBonus: intAdj[level][0],
			features:         feats,
			specialty:        map[string]int{"Cantrips": intAdj[level][1]},
			spellslots:       spells(level),
		}

		return jobInfo

	case character == "Cleric":
		hitPoints := 8

		for i := 1; i < level; i++ {
			hpRolls := roll(8, 1)
			for i := 0; i < len(hpRolls); i++ {
				hitPoints += hpRolls[i]
			}
		}

		possiblefeatures := map[string]int{"Spellcasting": 1, "Divine Domain": 1, "Channel Divinity (1/rest)": 2, "Divine Domain Feature #1": 2,
			"Destroy Undead (CR 1/2)": 5, "Channel Divinity (2/rest)": 6, "Divine Domain Feature #2": 6, "Destroy Undead (CR 1)": 8,
			"Divine Domain Feature #3": 8, "Divine Intervention": 10, "Destroy Undead (CR 2)": 11, "Destroy Undead (CR 3)": 14,
			"Destroy Undead (CR 4)": 17, "Divine Domain Feature #4": 17, "Channel Divinity (3/rest)": 18, "Divine Intervention Improvement": 20}

		feats := []string{}

		for k, v := range possiblefeatures {
			if v <= level {
				feats = append(feats, k)
			}
		}

		intAdj := map[int][]int{1: {2, 3}, 2: {2, 3}, 3: {2, 3}, 4: {2, 4}, 5: {3, 4}, 6: {3, 4}, 7: {3, 4}, 8: {3, 4}, 9: {4, 4}, 10: {4, 5}, 11: {4, 5},
			12: {4, 5}, 13: {5, 5}, 14: {5, 5}, 15: {5, 5}, 16: {5, 5}, 17: {6, 5}, 18: {6, 5}, 19: {6, 5}, 20: {6, 5}}

		jobInfo := jobDetails{
			hp:               hitPoints,
			savingThrows:     []string{"wisdom", "charisma"},
			proficiencies:    []string{"Light Armor", "Medium Armor", "Shields", "Simple Weapons"},
			proficiencyBonus: intAdj[level][0],
			features:         feats,
			specialty:        map[string]int{"Cantrips": intAdj[level][1]},
			spellslots:       spells(level),
		}

		return jobInfo

	case character == "Druid":
		hitPoints := 8

		for i := 1; i < level; i++ {
			hpRolls := roll(8, 1)
			for i := 0; i < len(hpRolls); i++ {
				hitPoints += hpRolls[i]
			}
		}

		possiblefeatures := map[string]int{"Spellcasting": 1, "Drudic": 1, "Wild Shape": 2, "Druid Circle": 2, "Wild Shape Improvement #1": 4,
			"Druid Circle feature #1": 6, "Wild Shape Improvement #2": 8, "Druid Circle feature #2": 10, "Druid Circle feature #3": 14, "Timeless Body": 18,
			"Beast Spells": 18, "Archdruid": 20}

		feats := []string{}

		for k, v := range possiblefeatures {
			if v <= level {
				feats = append(feats, k)
			}
		}

		intAdj := map[int][]int{1: {2, 2}, 2: {2, 2}, 3: {2, 2}, 4: {2, 3}, 5: {3, 3}, 6: {3, 3}, 7: {3, 3}, 8: {3, 3}, 9: {4, 3}, 10: {4, 4}, 11: {4, 4},
			12: {4, 4}, 13: {5, 4}, 14: {5, 4}, 15: {5, 4}, 16: {5, 4}, 17: {6, 4}, 18: {6, 4}, 19: {6, 4}, 20: {6, 4}}

		jobInfo := jobDetails{
			hp:               hitPoints,
			savingThrows:     []string{"wisdom", "intelligence"},
			proficiencies:    []string{"Light Armor", "Medium Armor", "Shields", "Weapons: clubs, daggers, darts, javelins, maces, quarterstaffs, sickles, slings, spears", "Herbalism Kits"},
			proficiencyBonus: intAdj[level][0],
			features:         feats,
			specialty:        map[string]int{"Cantrips": intAdj[level][1]},
			spellslots:       spells(level),
		}

		return jobInfo

	case character == "Fighter":
	case character == "Monk":
	case character == "Paladin":
	case character == "Ranger":
	case character == "Rogue":
	case character == "Sorcerer":
	case character == "Warlock":
	case character == "Wizard":
	default:
		fmt.Println("That job is not supported by this script. Job related character adjustments will not be added.\n")
	}

	var jobInfo jobDetails
	return jobInfo
}
