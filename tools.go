package tools

import (
	"fmt"
	"math/rand"
	"time"
)

// diceRoll takes the number of sides and the number of dice to simulate a random dice roll.
func diceRoll(sides int, dice int) []int {
	results := []int{}
	for i := 0; i < dice; i++ {
		results = append(results, rand.Intn(sides)+1)
	}
	return results
}

func sleep() {
	//formatting for easier readability
	time.Sleep(time.Second)
	fmt.Println("\n------------\n")
}
