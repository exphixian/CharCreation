package main

import (
	"math/rand"
	"time"
)

func roll(sides int, dice int) []int {
	rand.Seed(time.Now().UnixNano())
	results := []int{}
	for i := 0; i < dice; i++ {
		results = append(results, rand.Intn(sides) + 1)
	}
	return results
}
