package main

import (
	"math/rand"
	"time"
)

//currently used for HP rolls, but can be used for any set of dicerolls needed.
// input takes what dice you need (ie d4, d20, etc) and the number of rolls needed (ie 5 d4)

func roll(sides int, dice int) []int {
	rand.Seed(time.Now().UnixNano())
	results := []int{}
	for i := 0; i < dice; i++ {
		results = append(results, rand.Intn(sides)+1)
	}
	return results
}
