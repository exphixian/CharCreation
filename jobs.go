package main

import "fmt"

type jobDetails struct {
	job              string
	hp               int
	savingThrows     []string
	proficiencies    []string
	proficiencyBonus int
	features         []string
	specialty        map[string]int
	spellslots       map[int]int
}

func hpRolls(level int, hitPoints int, sides int) int {
	for i := 1; i < level; i++ {
		hpRolls := roll(sides, 1)
		for i := 0; i < len(hpRolls); i++ {
			hitPoints += hpRolls[i]
		}
	}
	return hitPoints
}

//level map for spellcasters supporting 5 spell levels
func spells5(level int) map[int]int {
	spells := map[int]map[int]int{1: {1: 2}, 2: {1: 3}, 3: {1: 3}, 4: {1: 4}, 5: {1: 4, 2: 2}, 6: {1: 4, 2: 2}, 7: {1: 4, 2: 3}, 8: {1: 4, 2: 3},
		9: {1: 4, 2: 3, 3: 2}, 10: {1: 4, 2: 3, 3: 2}, 11: {1: 4, 2: 3, 3: 3}, 12: {1: 4, 2: 3, 3: 3}, 13: {1: 4, 2: 3, 3: 3, 4: 1},
		14: {1: 4, 2: 3, 3: 3, 4: 1}, 15: {1: 4, 2: 3, 3: 3, 4: 2}, 16: {1: 4, 2: 3, 3: 3, 4: 2}, 17: {1: 4, 2: 3, 3: 3, 4: 3, 5: 1},
		18: {1: 4, 2: 3, 3: 3, 4: 3, 5: 1}, 19: {1: 4, 2: 3, 3: 3, 4: 3, 5: 2}, 20: {1: 4, 2: 3, 3: 3, 4: 3, 5: 2}}

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

//level map for spellcasters supporting 9 spell levels
func spells9(level int) map[int]int {
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

//identifying which mapped features are in scope for the character
func featParse(possiblefeatures map[string]int, level int) []string {
	feats := []string{}
	for k, v := range possiblefeatures {
		if v <= level {
			feats = append(feats, k)
		}
	}
	return feats
}

/*
	need to:
		*support multiclassing
		*support paths
	 	*support skill choices
*/

func jobMGMT(level int) jobDetails {

	fmt.Println("\n\nSupported jobs: Barbarian, Bard, Cleric, Druid, Fighter, Monk, Paladin, Ranger, Rogue, Sorcerer, Warlock, Wizard")
	character := stringInput("job")

	//level map for proficiency bonus
	profBonus := map[int]int{1: 2, 2: 2, 3: 2, 4: 2, 5: 3, 6: 3, 7: 3, 8: 3, 9: 4, 10: 4, 11: 4, 12: 4, 13: 5, 14: 5, 15: 5, 16: 5, 17: 6,
		18: 6, 19: 6, 20: 6}

	//level map for standard cantrips
	cantrips := map[int]int{1: 2, 2: 2, 3: 2, 4: 3, 5: 3, 6: 3, 7: 3, 8: 3, 9: 3, 10: 4, 11: 4, 12: 4, 13: 4, 14: 4, 15: 4, 16: 4, 17: 4,
		18: 4, 19: 4, 20: 4}

	//class identification & build information
	switch {
	case character == "Barbarian":
		hitPoints := 12

		possiblefeatures := map[string]int{"Rage": 1, "Unarmored Defense": 1, "Reckless Attack": 2, "Danger Sense": 2, "Primal Path": 3,
			"Extra Attack": 5, "Fast Movement": 5, "Feral Instinct": 7, "Brutal Critical (1)": 9, "Relentless Rage": 11, "Brutal Critical (2)": 13,
			"Persistent Rage": 15, "Brutal Critical (3)": 17, "Indomitable Might": 18, "Primal Champion": 20}

		feats := featParse(possiblefeatures, level)

		specialtyInts := map[int][]int{1: {2, 2}, 2: {2, 2}, 3: {3, 2}, 4: {3, 2}, 5: {3, 2}, 6: {4, 2}, 7: {4, 2}, 8: {4, 2}, 9: {4, 3},
			10: {4, 3}, 11: {4, 3}, 12: {5, 3}, 13: {5, 3}, 14: {5, 3}, 15: {5, 3}, 16: {5, 4}, 17: {6, 4}, 18: {6, 4}, 19: {6, 4},
			20: {00, 4}}

		jobInfo := jobDetails{
			job:              character,
			hp:               hpRolls(level, hitPoints, hitPoints),
			savingThrows:     []string{"strength", "constitution"},
			proficiencies:    []string{"Light Armor", "Medium Armor", "Shields", "Simple Weapons", "Martial Weapons"},
			proficiencyBonus: profBonus[level],
			features:         feats,
			specialty:        map[string]int{"Rages: ": specialtyInts[level][0], "Rage Damage: ": specialtyInts[level][1]},
		}

		return jobInfo

	case character == "Bard":
		hitPoints := 8

		possiblefeatures := map[string]int{"Spellcasting": 1, "Bardic Inspiration (d6)": 1, "Jack of All Trades": 2, "Song of Rest (d6)": 2,
			"Bard College": 3, "Expertise #1": 3, "Bardic Inspiration (d8)": 5, "Font of Inspiration": 5, "Counter Charm": 6, "Bard College Feature #1": 6,
			"Song of Rest (d8)": 9, "Bardic Inspiration (d10)": 10, "Expertise #2": 10, "Magical Secrets #1": 10, "Song of Rest (d10)": 13,
			"Magical Secrets #2": 14, "Bardic Inspiration (d12)": 15, "Song of Rest (d12)": 17, "Magical Secrets #3": 18, "Superior Inspiration": 20}

		feats := featParse(possiblefeatures, level)

		jobInfo := jobDetails{
			job:              character,
			hp:               hpRolls(level, hitPoints, hitPoints),
			savingThrows:     []string{"dexterity", "charisma"},
			proficiencies:    []string{"Light Armor", "Hand Crossbows", "Longswords", "Simple Weapons", "Rapiers", "Shortswords", "3 musical instruments"},
			proficiencyBonus: profBonus[level],
			features:         feats,
			specialty:        map[string]int{"Cantrips: ": cantrips[level]},
			spellslots:       spells9(level),
		}

		return jobInfo

	case character == "Cleric":
		hitPoints := 8

		possiblefeatures := map[string]int{"Spellcasting": 1, "Divine Domain": 1, "Channel Divinity (1/rest)": 2, "Divine Domain Feature #1": 2,
			"Destroy Undead (CR 1/2)": 5, "Channel Divinity (2/rest)": 6, "Divine Domain Feature #2": 6, "Destroy Undead (CR 1)": 8,
			"Divine Domain Feature #3": 8, "Divine Intervention": 10, "Destroy Undead (CR 2)": 11, "Destroy Undead (CR 3)": 14,
			"Destroy Undead (CR 4)": 17, "Divine Domain Feature #4": 17, "Channel Divinity (3/rest)": 18, "Divine Intervention Improvement": 20}

		feats := featParse(possiblefeatures, level)

		specialtyInts := map[int]int{1: 3, 2: 3, 3: 3, 4: 4, 5: 4, 6: 4, 7: 4, 8: 4, 9: 4, 10: 5, 11: 5, 12: 5, 13: 5, 14: 5, 15: 5, 16: 5,
			17: 5, 18: 5, 19: 5, 20: 5}

		jobInfo := jobDetails{
			job:              character,
			hp:               hpRolls(level, hitPoints, hitPoints),
			savingThrows:     []string{"wisdom", "charisma"},
			proficiencies:    []string{"Light Armor", "Medium Armor", "Shields", "Simple Weapons"},
			proficiencyBonus: profBonus[level],
			features:         feats,
			specialty:        map[string]int{"Cantrips": specialtyInts[level]},
			spellslots:       spells9(level),
		}

		return jobInfo

	case character == "Druid":
		hitPoints := 8

		possiblefeatures := map[string]int{"Spellcasting": 1, "Drudic": 1, "Wild Shape": 2, "Druid Circle": 2, "Wild Shape Improvement #1": 4,
			"Druid Circle feature #1": 6, "Wild Shape Improvement #2": 8, "Druid Circle feature #2": 10, "Druid Circle feature #3": 14, "Timeless Body": 18,
			"Beast Spells": 18, "Archdruid": 20}

		feats := featParse(possiblefeatures, level)

		jobInfo := jobDetails{
			job:          character,
			hp:           hpRolls(level, hitPoints, hitPoints),
			savingThrows: []string{"wisdom", "intelligence"},
			proficiencies: []string{"Light Armor (non-metal)", "Medium Armor (non-metal)", "Shields (non-metal)", "Herbalism Kits",
				"Weapons: clubs, daggers, darts, javelins, maces, quarterstaffs, sickles, slings, spears"},
			proficiencyBonus: profBonus[level],
			features:         feats,
			specialty:        map[string]int{"Cantrips: ": cantrips[level]},
			spellslots:       spells9(level),
		}

		return jobInfo

	case character == "Fighter":
		hitPoints := 10

		possiblefeatures := map[string]int{"Fighting Style": 1, "Second Wind": 1, "Action Surge (1 use)": 2, "Martial Archetype": 3, "Extra Attack (1)": 5,
			"Martial Archetype feat #1": 7, "Indomitable (1 use)": 9, "Martial Archetype feat #2": 10, "Extra Attack (2)": 11, "Indomitable (2 uses)": 13,
			"Martial Archetype feat #3": 15, "Action Surge (2 uses)": 17, "Indomitable (3 uses)": 17, "Martial Archetype feat #4": 18, "Extra Attack (3)": 20}

		feats := featParse(possiblefeatures, level)

		jobInfo := jobDetails{
			job:              character,
			hp:               hpRolls(level, hitPoints, hitPoints),
			savingThrows:     []string{"strength", "constitution"},
			proficiencies:    []string{"All Armor", "Shields", "Simple Weapons", "Martial Weapons"},
			proficiencyBonus: profBonus[level],
			features:         feats,
		}

		return jobInfo

	case character == "Monk":
		hitPoints := 8

		possiblefeatures := map[string]int{"Martial Arts": 1, "Unarmored Defense": 1, "Ki": 2, "Unarmored Movement": 2, "Monastic Tradition": 3,
			"Deflect Missiles": 3, "Slow Fall": 4, "Extra Attack (1)": 5, "Stunning Strike": 5, "Ki-Empowered Strikes": 6, "Monastic Tradition Feat #1": 6,
			"Evasion": 7, "Stillness of Mind": 7, "Unarmored Movement Improvement": 9, "Purity of Body": 10, "Monastic Tradition Feat #2": 11,
			"Tongue of the Sun and Moon": 13, "Diamond Soul": 14, "Timeless Body": 15, "Monastic Tradition Feat #3": 17, "Empty Body": 18, "Perfect Self": 20}

		feats := featParse(possiblefeatures, level)

		specialtyInts := map[int][]int{1: {4, 0, 0}, 2: {4, 2, 10}, 3: {4, 3, 10}, 4: {4, 4, 10}, 5: {6, 5, 10}, 6: {6, 6, 15},
			7: {6, 7, 15}, 8: {6, 8, 15}, 9: {6, 9, 15}, 10: {6, 10, 20}, 11: {8, 11, 20}, 12: {8, 12, 20}, 13: {8, 13, 20},
			14: {8, 14, 25}, 15: {8, 15, 25}, 16: {8, 16, 25}, 17: {10, 17, 25}, 18: {10, 18, 30}, 19: {10, 19, 30}, 20: {10, 20, 30}}

		jobInfo := jobDetails{
			job:              character,
			hp:               hpRolls(level, hitPoints, hitPoints),
			savingThrows:     []string{"strength", "dexterity"},
			proficiencies:    []string{"Simple Weapons", "Shortswords"},
			proficiencyBonus: profBonus[level],
			features:         feats,
			specialty:        map[string]int{"Martial Arts: 1d": specialtyInts[level][0], "Ki Points:": specialtyInts[level][1], "Unarmored Movement: +": specialtyInts[level][2]},
		}
		return jobInfo

	case character == "Paladin":
		hitPoints := 10

		possiblefeatures := map[string]int{"Divine Sense": 1, "Lay on Hands": 1, "Fighting Style": 2, "Spellcasting": 2, "Divine Smite": 2,
			"Divine Health": 3, "Sacred Oath": 3, "Extra Attack": 5, "Aura of Protection": 6, "Sacred Oath feature #1": 7, "Aura of Courage": 10,
			"Improved Divine Smite": 11, "Cleansing Touch": 14, "Sacred Oath feature #2": 15, "Aura Improvements": 18, "Sacred Oath feature #3": 20}

		feats := featParse(possiblefeatures, level)

		jobInfo := jobDetails{
			job:              character,
			hp:               hpRolls(level, hitPoints, hitPoints),
			savingThrows:     []string{"wisdom", "charisma"},
			proficiencies:    []string{"All Armor", "Shields", "Simple Weapons", "Martial Weapons"},
			proficiencyBonus: profBonus[level],
			features:         feats,
			spellslots:       spells5(level),
		}
		return jobInfo

	case character == "Ranger":
		hitPoints := 10

		possiblefeatures := map[string]int{"Favored Enemy": 1, "Natural Explorer": 1, "Fighting Style": 2, "Spellcasting": 2, "Ranger Archetype": 3,
			"Primeval Awareness": 3, "Extra Attack": 5, "Favored Enemy Improvement #1": 6, "Natural Explorer Improvement #1": 6,
			"Ranger Archetype feature #1": 7, "Land's Stride": 8, "Natural Explorer Improvement #2": 10, "Hide in Plain Sight": 10,
			"Ranger Archetype feature #2": 11, "Vanish": 14, "Favored Enemy Improvement #2": 14, "Ranger Archetype feature #3": 15,
			"Feral Senses": 18, "Foe Slayer": 20}

		feats := featParse(possiblefeatures, level)

		jobInfo := jobDetails{
			job:              character,
			hp:               hpRolls(level, hitPoints, hitPoints),
			savingThrows:     []string{"strength", "dexterity"},
			proficiencies:    []string{"Light Armor", "Medium Armor", "Simple Weapons", "Martial Weapons"},
			proficiencyBonus: profBonus[level],
			features:         feats,
			spellslots:       spells5(level),
		}
		return jobInfo

	case character == "Rogue":
		hitPoints := 8

		possiblefeatures := map[string]int{"Expertise #1": 1, "Sneak Attack": 1, "Thieves' Cant": 1, "Cunning Action": 2, "Roguish Archetype": 3,
			"Uncanny Dodge": 5, "Expertise #2": 6, "Evasion": 7, "Roguish Archetype feature #1": 9, "Reliable Talent": 11, "Roguish Archetype feature #2": 13,
			"Blindsense": 14, "Slippery Mind": 15, "Roguish Archetype feature #3": 17, "Elusive": 18, "Stroke of Luck": 20}

		feats := featParse(possiblefeatures, level)

		jobInfo := jobDetails{
			job:              character,
			hp:               hpRolls(level, hitPoints, hitPoints),
			savingThrows:     []string{"dexterity", "intelligence"},
			proficiencies:    []string{"Light Armor", "Simple Weapons", "Hand Crossbows", "Longswords", "Rapiers", "Shortswords", "Thieves' Tools"},
			proficiencyBonus: profBonus[level],
			features:         feats,
		}
		return jobInfo

	case character == "Sorcerer":
		hitPoints := 6

		specialtyInts := map[int]int{1: 4, 2: 4, 3: 4, 4: 5, 5: 5, 6: 5, 7: 5, 8: 5, 9: 5, 10: 6, 11: 6, 12: 6, 13: 6, 14: 6, 15: 6, 16: 6, 17: 6,
			18: 6, 19: 6, 20: 6}

		possiblefeatures := map[string]int{"Spellcasting": 1, "Sorcerous Origin": 1, "Font of Magic": 2, "Metamagic #1": 3, "Sorcerous Origin feat #1": 6,
			"Metamagic #2": 10, "Sorcerous Origin feat #2": 14, "Metamagic #3": 17, "Sorcerous Origin feat #3": 18, "Sorcerous Restoration": 20}

		feats := featParse(possiblefeatures, level)

		jobInfo := jobDetails{
			job:              character,
			hp:               hpRolls(level, hitPoints, hitPoints),
			savingThrows:     []string{"constitution", "charisma"},
			proficiencies:    []string{"Daggers", "Darts", " Slings", "Quarterstaffs", "Light Crossbows"},
			proficiencyBonus: profBonus[level],
			features:         feats,
			specialty:        map[string]int{"Cantrips": specialtyInts[level]},
			spellslots:       spells9(level),
		}
		return jobInfo

	case character == "Warlock":
		hitPoints := 8

		specialtyInts := map[int][]int{1: {2, 1, 1, 0}, 2: {3, 2, 1, 2}, 3: {4, 2, 2, 2}, 4: {5, 2, 2, 2}, 5: {6, 2, 3, 3}, 6: {7, 2, 3, 3},
			7: {8, 2, 4, 4}, 8: {9, 2, 4, 4}, 9: {10, 2, 5, 5}, 10: {10, 2, 5, 5}, 11: {11, 3, 5, 5}, 12: {11, 3, 5, 6}, 13: {12, 3, 5, 6},
			14: {12, 3, 5, 6}, 15: {13, 3, 5, 7}, 16: {13, 3, 5, 7}, 17: {14, 4, 5, 7}, 18: {14, 4, 5, 8}, 19: {15, 4, 5, 8}, 20: {15, 4, 5, 8}}

		possiblefeatures := map[string]int{"Overworldly Patron": 1, "Pact Magic": 1, "Eldritch Invocations": 2, "Pact Boon": 3,
			"Overworldly Patron feat #1": 6, "Overworldly Patron feat #2": 10, "Mystic Arcanum (6th level)": 11, "Mystic Arcanum (7th level)": 13,
			"Overworldly Patron feat #3": 14, "Mystic Arcanum (8th level)": 15, "Mystic Arcanum (9th level)": 17, "Eldritch Master": 20}

		feats := featParse(possiblefeatures, level)

		jobInfo := jobDetails{
			job:              character,
			hp:               hpRolls(level, hitPoints, hitPoints),
			savingThrows:     []string{"wisdom", "charisma"},
			proficiencies:    []string{"Light Armor", "Simple Weapons"},
			proficiencyBonus: profBonus[level],
			features:         feats,
			specialty: map[string]int{"Cantrips: ": cantrips[level], "Spells Known: ": specialtyInts[level][0],
				"Spell Slots: ": specialtyInts[level][1], "Slot Level: ": specialtyInts[level][2], "Invocations Known: ": specialtyInts[level][3]},
		}
		return jobInfo

	case character == "Wizard":
		hitPoints := 6

		possiblefeatures := map[string]int{"Spellcasting": 1, "Arcane Recovery": 1, "Arcane Tradition": 2, "Arcane Tradition feat #1": 6,
			"Arcane Tradition feat #2": 10, "Arcane Tradition feat #3": 14, "Spell Mastery": 18, "Signature Spell": 20}

		feats := featParse(possiblefeatures, level)

		jobInfo := jobDetails{
			job:              character,
			hp:               hpRolls(level, hitPoints, hitPoints),
			savingThrows:     []string{"wisdom", "intelligence"},
			proficiencies:    []string{"Daggers", "Darts", "Quarterstaffs", "Slings", "Light Crossbows"},
			proficiencyBonus: profBonus[level],
			features:         feats,
			specialty:        map[string]int{"Cantrips": cantrips[level]},
			spellslots:       spells9(level),
		}

		return jobInfo

	default:
		fmt.Println("That job is not supported by this script. Job related character adjustments will not be added.\n")
	}

	jobInfo := jobDetails{
		job: character,
	}
	return jobInfo
}
