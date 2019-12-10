package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	dataString := string(data)

	wires := strings.Fields(dataString)
	wire1Directions := strings.Split(wires[0], ",")
	wire2Directions := strings.Split(wires[1], ",")

	wire1 := convertDirectionsToMapOfCoordinates(wire1Directions)
	wire2 := convertDirectionsToMapOfCoordinates(wire2Directions)

	elapsed := time.Since(start)
	fmt.Println("Execution Time: %s", elapsed)
}

func convertDirectionsToMapOfCoordinates(directions []string) map[int]wireCoordinate{
	currentX := 0
	currentY := 0

	var wire map[int]wireCoordinate; 
	for dirAndLength := range directions{
		length = strconv.Atoi(dirAndLength[1:])
		if dirAndLength[0] == 'U'{
			
		}else if dirAndLength[0] == 'D'{

		}else if dirAndLength[0] == 'L'{
			
		}else if dirAndLength[0] == 'R'{
			
		}
	}
}

type wireCoordinate struct {
	x int
	y int
}
