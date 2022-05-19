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
	//formatting for easier readability
	time.Sleep(time.Second)
	fmt.Println("\n------------\n")
}

func stringInput(req string) string {
	//reads string input from the user and returns it.
	stringInput := bufio.NewScanner(os.Stdin)
	fmt.Printf("What is your character's %s?\n", req)
	stringInput.Scan()
	variable := stringInput.Text()
	return variable
}

func roll() int { //3 d6
	rand.Seed(time.Now().UnixNano())
	stat := rand.Intn(15) + 4
	return stat
}

func modifier(modifier int) int { //calculates modifiers from stats
	if modifier < 10 {
		modifier -= 1
	}
	modifier = (modifier - 10) / 2
	return modifier
}

func basicInfo() (string, string, string, int) { //gathers needed input for use is remainder of script
	Name = stringInput("name")

	fmt.Printf("What level is your character?\n")
	fmt.Scan(&Level)
	if Level > 20 {
		fmt.Println("This sheet does not support legendary characters at this time.")
		os.Exit(1)
	}

	Species = stringInput("species")
	Job = stringInput("class")

	sleep()

	return Name, Species, Job, Level
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

	//assigns stats
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

	//ability points based on level/job
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
		fmt.Printf("%v: %v\n", Cats[i], Stats[Cats[i]])
		fmt.Printf("%v: %v\n", Modifiers[i], Mods[Modifiers[i]])
		fmt.Println()
	}
	return Stats, Mods
}

func main() {
	//fmt.Println("placeholder")\n", Name

	basicInfo()
	fmt.Printf("%v is a level %v %v %v.\n\n", Name, Level, Species, Job)

	diceroll(Level, Job)

}
