package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	var totalFuel float64 = 0

	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	dataString := string(data)

	s := strings.Fields(dataString)

	for _, value := range s {
		float, _ := strconv.ParseFloat(value, 64)
		totalFuel = totalFuel + calculateFuelForFuel(float)
	}

	outputstring := fmt.Sprintf("%f", totalFuel)

	fmt.Println("Output:", outputstring)
	elapsed := time.Since(start)
	fmt.Println("Execution Time: %s", elapsed)
}

func calculateFuelForFuel(massInput float64) float64 {
	var fuelreq float64 = math.Floor(massInput/3) - 2
	if fuelreq < 1 {
		return 0
	}
	return fuelreq + calculateFuelForFuel(fuelreq)
}
