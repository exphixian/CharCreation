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

func modifier(stat int) int {
	modifier := (stat - 10) / 2
	if modifier < 0 {
		modifier--
	}
	return modifier
}

func main() {
	var random string
	fmt.Println("Do you want your stats to be randomly generated?")
	fmt.Scan(&random)

	Stats := make(map[string]int)
	Mods := make(map[string]int)
	Cats := [6]string{"Strength", "Constitution", "Dexterity", "Intelligence", "Wisdom", "Charisma"}
	Modifiers := [6]string{"STR", "CON", "DEX", "INT", "WIS", "CHA"}

	if random == "yes" {
		for i := 0; i < len(Cats); i++ {
			Stats[Cats[i]] = roll()
		}

	} else if random == "no" {
		fmt.Println("Please insert your stats.")
		var input int
		for i := 0; i < len(Cats); i++ {
			fmt.Printf("\n%v:", Cats[i])
			fmt.Scan(&input)
			Stats[Cats[i]] = input
		}
	}

	fmt.Println(Stats)
	fmt.Println(Mods)

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

	fmt.Println(Stats)
	fmt.Println(Mods)

}
