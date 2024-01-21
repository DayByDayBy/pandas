package main

import (
	"fmt"
	"math/rand"
)

func main() {

	people := 1280
	childrenPerPair := 5
	generations := 0
	population := people

	for population < 7000000000 {
		newBirths := ((population / 2) * childrenPerPair) * int(randomMortalityFactor())
		population += newBirths 
		generations++
	}

	fmt.Printf("Population reached 7 billion in approximately %d generations, or %d years \n", generations, generations*25)
}

func randomMortalityFactor() float64 {
	if rand.Intn(2) == 0 {
		return 0.0
	}
	return 1.0
}
