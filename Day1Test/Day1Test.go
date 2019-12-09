package main

import "fmt"

import "math"

func main() {

	var inputMass float64 = 14

	var outputMass float64 = math.Floor(inputMass/3) - 2

	fmt.Println("Output:", outputMass)
}
