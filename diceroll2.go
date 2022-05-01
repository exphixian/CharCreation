package main

import (
	"fmt"
	"math/rand"
	"time"
)

func roll() int {
	rand.Seed(time.Now().UnixNano())
	stat := rand.Intn(15) + 3
	return stat
}

func modifier(modifier int) int {
	if modifier < 10 {
		modifier -= 1
	}
	modifier = (modifier - 10) / 2
	return modifier
}

func main() {
	var generate string
	var random bool

	fmt.Println("Yes or No: do you want your stats to be randomly generated?")
	fmt.Scan(&generate)
	if generate == "yes" {
		random = true
	} else if generate == "Yes" {
		random = true
	} else {
		random = false
	}

	Stats := make(map[string]int)
	Mods := make(map[string]int)
	Cats := [6]string{"strength", "constitution", "dexterity", "intelligence", "wisdom", "charisma"}
	Modifiers := [6]string{"STR", "CON", "DEX", "INT", "WIS", "CHA"}

	if random {
		for i := 0; i < len(Cats); i++ {
			Stats[Cats[i]] = roll()
		}

	} else {
		fmt.Println("Please insert your stats.")
		var input int
		for i := 0; i < len(Cats); i++ {
			fmt.Printf("\n%v:", Cats[i])
			fmt.Scan(&input)
			Stats[Cats[i]] = input
		}
	}

	fmt.Println(Stats)
	//fmt.Println(Mods)

	var level int
	fmt.Println("What level is your character?")
	fmt.Scan(&level)
	level = level / 4
	//need to adjust for fighter and monk
	if level > 0 {
		level *= 2
		fmt.Printf("Where do you want to allocate your %d leveling points?\n", level)
		for i := 0; i < level; i++ {
			//fmt.Println(i)
			var input string
			fmt.Scan(&input)
			Stats[input] += 1
		}
	}

	for i := 0; i < len(Modifiers); i++ {
		Mods[Modifiers[i]] = modifier(Stats[Cats[i]])
	}

	time.Sleep(time.Second)

	for i := 0; i < len(Cats); i++ {
		fmt.Printf("\n%v: %v", Cats[i], Stats[Cats[i]])
		fmt.Printf("\n%v: %v", Modifiers[i], Mods[Modifiers[i]])
		fmt.Println()
	}
}
