package main

import (
	"math/rand"
	"time"
)

func ArrayGenerate(size int) []int {

	rand.Seed(time.Now().UnixNano())

	slice := make([]int, size)

	for i := range slice {
		slice[i] = rand.Intn(10001)
	}

	return slice
}
