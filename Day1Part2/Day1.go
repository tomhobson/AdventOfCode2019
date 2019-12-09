package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func main() {
	var totalFuel float64 = 0

	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	dataString := string(data)

	s := strings.Fields(dataString)

	fmt.Println(s)

	for _, value := range s {
		float, _ := strconv.ParseFloat(value, 64)
		totalFuel = totalFuel + calculateFuelForFuel(float)
	}

	outputstring := fmt.Sprintf("%f", totalFuel)

	fmt.Println("Output:", outputstring)
}

func calculateFuelForFuel(massInput float64) float64{
	var fuelreq float64 = math.Floor(massInput/3) - 2
	if fuelreq < 1 {
		return 0
	}
	CurrentCalcFuelReq := fmt.Sprintf("%f", fuelreq)
	fmt.Println("Output:", CurrentCalcFuelReq)
	return fuelreq + calculateFuelForFuel(fuelreq)
}
