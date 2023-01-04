package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
type char struct {
	name       string
	level      int
	hp         int
	ac         int
	throws     []int
	species    string
	subspecies string
	job        string
	stats      map[string]int
	mods       map[string]int
	languages  []string
	feats      []string
	spells     []string
}

var character char
*/
func stringInput(req string) string {
	//reads string input from the user and returns it.
	stringInput := bufio.NewScanner(os.Stdin)
	fmt.Printf("What is your character's %s?\n", req)
	stringInput.Scan()
	variable := stringInput.Text()
	return variable
}

func main() {
	name := stringInput("name")

	level := 0
	fmt.Printf("What level is your character?\n")
	fmt.Scan(&level)
	if level > 20 {
		fmt.Println("This sheet does not support legendary characters at this time.")
		os.Exit(1)
	}

	species := stringInput("species")
	subspecies := ""

	switch {
	case species == "Dwarf":
		fmt.Println("Which Dwarven subspecies are you playing? (Hill or Mountain)")
		fmt.Scanln(&subspecies)
	case species == "Elf":
		fmt.Println("Which Elven subspecies are you playing? (High, Wood or Dark)")
		fmt.Scanln(&subspecies)
	case species == "Halfling":
		fmt.Println("Which Halfling subspecies are you playing? (Lightfoot or Stout)")
		fmt.Scanln(&subspecies)
	case species == "Gnome":
		fmt.Println("Which Gnome subspecies are you playing? (Forest or Rock)")
		fmt.Scanln(&subspecies)
	case species == "Human" || species == "Dragonborn" || species == "Half-Elf" || species == "Half-Orc" || species == "Tiefling":
		fmt.Printf("\nNo %v subspecies supported at this time.\n", species)

	default:
		fmt.Println("That species is not supported by this script. Species related character adjustments will not be added.")

	}

	//Need to support multiclassing.
	job := stringInput("job")

	sleep()

	fmt.Printf("%v is a level %v %v %v.\n\n", name, level, species, job)

	var generate string
	var stats map[string]int
	var mods map[string]int

	fmt.Println("Yes or No: do you want your stats to be randomly generated?")

	fmt.Scan(&generate)
	if generate == "yes" || generate == "Yes" {
		stats, mods = randomizedStats(level, species, subspecies, job)
	} else {
		stats, mods = manualStats(level, species, subspecies, job)
	}

	fmt.Println(stats, mods)

}
