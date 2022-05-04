//Exit status 1: level over 20 detected

package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var Name, Species, Job string
var Level int

func sleep() {
	time.Sleep(time.Second)

}

func basicInfo() (string, string, string, int) {
	input := bufio.NewScanner(os.Stdin)
	fmt.Println("What is your character's name?")
	input.Scan()
	Name = input.Text()

	fmt.Printf("What level is %v?\n", Name)
	fmt.Scan(&Level)
	if Level > 20 {
		fmt.Println("This sheet does not support legendary characters at this time.")
		os.Exit(1)
	}

	input = bufio.NewScanner(os.Stdin)
	fmt.Printf("What species is %v?\n", Name)
	input.Scan()
	Species = input.Text()

	input = bufio.NewScanner(os.Stdin)
	fmt.Println("Which class will you be playing?")
	input.Scan()
	Job = input.Text()

	sleep()

	return Name, Species, Job, Level
}

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

func diceroll(Level int, Job string) (map[string]int, map[string]int) {
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

	for i := 0; i < len(Cats); i++ {
		fmt.Printf("\n%v: %v", Cats[i], Stats[Cats[i]])
		fmt.Println()
	}

	sleep()

	if Level > 3 {
		abilityPoints := (Level / 4) * 2
		if Level >= 19 {
			abilityPoints += 2
		}

		if Job == "fighter" {
			if Level >= 14 {
				abilityPoints += 4
			} else if Level >= 6 {
				abilityPoints += 2
			}

		} else if Job == "rogue" {
			if Level >= 10 {
				abilityPoints += 2
			}
		}

		fmt.Printf("Where do you want to allocate your %d ability points?\n", abilityPoints)
		for i := 0; i < abilityPoints; i++ {
			var input string
			fmt.Scan(&input)
			Stats[input] += 1
		}
	}

	for i := 0; i < len(Modifiers); i++ {
		Mods[Modifiers[i]] = modifier(Stats[Cats[i]])
	}

	sleep()

	for i := 0; i < len(Cats); i++ {
		fmt.Printf("\n%v: %v", Cats[i], Stats[Cats[i]])
		fmt.Printf("\n%v: %v", Modifiers[i], Mods[Modifiers[i]])
		fmt.Println()
	}
	return Stats, Mods
}

func main() {
	//fmt.Println("placeholder")

	basicInfo()
	fmt.Printf("%v is a level %v %v %v.\n\n", Name, Level, Species, Job)

	diceroll(Level, Job)

}
