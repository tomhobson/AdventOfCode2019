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
		var outputMass float64 = math.Floor(float/3) - 2
		totalFuel = totalFuel + outputMass
	}

	outputstring := fmt.Sprintf("%f", totalFuel)

	fmt.Println("Output:", outputstring)
}
