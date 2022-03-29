package main

import (
	"fmt"
	"math/rand"
	"time"
)

var Strength, Constitution, Dexterity, Charisma, Wisdom, Intelligence int
var Str, Con, Dex, Cha, Wis, Int int

//var stats map[string] int
//var mods map[string] int

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

func classModifier() {
	fmt.Println("\nClass modifiers for stats is in development.")
	time.Sleep(time.Second)

}

func levelAdjustment() {
	fmt.Println("\nLevel Adjustment for stats is in development.")
	time.Sleep(time.Second)

}

func diceroll() {
	var random string
	fmt.Println("Do you want your stats randomly generated?")
	fmt.Scan(&random)

	if random == "yes" {
		Strength = roll()
		Constitution = roll()
		Dexterity = roll()
		Charisma = roll()
		Wisdom = roll()
		Intelligence = roll()

	} else {
		fmt.Println("Please insert your stats.")
		fmt.Print("Str: ")
		fmt.Scan(&Strength)
		fmt.Print("Con: ")
		fmt.Scan(&Constitution)
		fmt.Print("Dex: ")
		fmt.Scan(&Dexterity)
		fmt.Print("Cha: ")
		fmt.Scan(&Charisma)
		fmt.Print("Wis: ")
		fmt.Scan(&Wisdom)
		fmt.Print("Int: ")
		fmt.Scan(&Intelligence)
	}

	classModifier()

	levelAdjustment()

	Str := modifier(Strength)
	Con := modifier(Constitution)
	Dex := modifier(Dexterity)
	Cha := modifier(Charisma)
	Wis := modifier(Wisdom)
	Int := modifier(Intelligence)

	fmt.Println("\nYour base stats are...")
	time.Sleep(time.Second)
	fmt.Printf("\nStr: %d, mod %d\n", Strength, Str)
	time.Sleep(time.Second)
	fmt.Printf("Con: %d, mod %d\n", Constitution, Con)
	time.Sleep(time.Second)
	fmt.Printf("Dex: %d, mod %d\n", Dexterity, Dex)
	time.Sleep(time.Second)
	fmt.Printf("Cha: %d, mod %d\n", Charisma, Cha)
	time.Sleep(time.Second)
	fmt.Printf("Wis: %d, mod %d\n", Wisdom, Wis)
	time.Sleep(time.Second)
	fmt.Printf("Int: %d, mod %d\n", Intelligence, Int)
}
