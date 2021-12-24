package main

import (
	"flag"
	"fmt"
)

func main() {

	amountPizza := flag.Int("AmountPizza", 1, "an int")
	DoughBallWeight := flag.Int("DoughBallWeight", 220, "a int")

	// var baseFlour int = 100
	var baseWater int = 65
	var baseYeast float32 = 0.1
	var baseSalt float32 = 1.8
	flag.Parse()

	var flour int = *DoughBallWeight * *amountPizza
	var water float64 = float64(*DoughBallWeight) * (float64(baseWater) / 100)
	var yeast float64 = float64(*DoughBallWeight) * (float64(baseYeast) / 10) // magic
	var salt float64 = float64(*DoughBallWeight) * (float64(baseSalt) / 100)

	fmt.Println("Flour Total: ", flour)
	fmt.Println("Water Total: ", water)
	fmt.Println("Yeast Total: ", yeast)
	fmt.Println("Salt Total: ", salt)

}
