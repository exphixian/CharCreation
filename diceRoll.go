package main

import (
	"fmt"
	"math/rand"
)

var strength, constitution, dexterity, charisma, wisdom, intelligence int

func diceroll() {
	var random string
	fmt.Println("Do you want your stats randomly generated?")
	fmt.Scan(&random)

	if random == "yes" {
		strength = rand.Intn(12) + 6
		constitution = rand.Intn(12) + 6
		dexterity = rand.Intn(12) + 6
		charisma = rand.Intn(12) + 6
		wisdom = rand.Intn(12) + 6
		intelligence = rand.Intn(12) + 6

	} else {
		fmt.Println("Please insert your stats.")
		fmt.Print("Str: ")
		fmt.Scan(&strength)
		fmt.Print("Con: ")
		fmt.Scan(&constitution)
		fmt.Print("Dex: ")
		fmt.Scan(&dexterity)
		fmt.Print("Cha: ")
		fmt.Scan(&charisma)
		fmt.Print("Wis: ")
		fmt.Scan(&wisdom)
		fmt.Print("Int: ")
		fmt.Scan(&intelligence)
	}

}
