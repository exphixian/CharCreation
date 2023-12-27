package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

// reads string input from the user and returns it.
func stringInput(req string) string {
	stringInput := bufio.NewScanner(os.Stdin)
	fmt.Printf("\nWhat is your character's %s?\n", req)
	stringInput.Scan()
	variable := stringInput.Text()
	return variable
}

// diceRoll takes the number of sides and the number of dice to simulate a random dice roll.
func diceRoll(sides int, dice int) []int {
	results := []int{}
	for i := 0; i < dice; i++ {
		results = append(results, rand.Intn(sides)+1)
	}
	return results
}

// sleep forces a break between printing to assist with response readability.
func sleep() {
	time.Sleep(time.Second)
	fmt.Println("\n------------\n")
}
