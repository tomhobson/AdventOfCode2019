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

	var wiremap map[int]map[int]WirePoint
	wiremap = make(map[int]map[int]WirePoint)

	// Do wire 1
	currentX := 0
	currentY := 0

	// Init zero
	wiremap[currentX] = make(map[int]WirePoint)

	for _, dirAndLength := range wire1Directions {
		newX := currentX
		newY := currentY
		length, _ := strconv.Atoi(dirAndLength[1:])
		if dirAndLength[0] == 'U' {
			newY = currentY - length
		} else if dirAndLength[0] == 'D' {
			newY = currentY + length
		} else if dirAndLength[0] == 'L' {
			newX = currentX - length
		} else if dirAndLength[0] == 'R' {
			newX = currentX + length
		}

		for currentX != newX {
			if currentX > newX {
				currentX = currentX - 1
				if wiremap[currentX] == nil {
					wiremap[currentX] = make(map[int]WirePoint)
				}
				wp := wiremap[currentX][currentY]
				wp.Wire1HasHit = true
				wiremap[currentX][currentY] = wp
			} else if currentX < newX {
				currentX = currentX + 1
				if wiremap[currentX] == nil {
					wiremap[currentX] = make(map[int]WirePoint)
				}
				wp := wiremap[currentX][currentY]
				wp.Wire1HasHit = true
				wiremap[currentX][currentY] = wp
			} else {
				break
			}
		}

		for currentY != newY {
			if currentY > newY {
				currentY = currentY - 1
				wp := wiremap[currentX][currentY]
				wp.Wire1HasHit = true
				wiremap[currentX][currentY] = wp
			} else if currentY < newY {
				currentY = currentY + 1
				wp := wiremap[currentX][currentY]
				wp.Wire1HasHit = true
				wiremap[currentX][currentY] = wp
			} else {
				break
			}
		}
	}

	//Do wire 2
	currentX = 0
	currentY = 0

	for _, dirAndLength := range wire2Directions {
		newX := currentX
		newY := currentY
		length, _ := strconv.Atoi(dirAndLength[1:])
		if dirAndLength[0] == 'U' {
			newY = currentY - length
		} else if dirAndLength[0] == 'D' {
			newY = currentY + length
		} else if dirAndLength[0] == 'L' {
			newX = currentX - length
		} else if dirAndLength[0] == 'R' {
			newX = currentX + length
		}

		for currentX != newX {
			if currentX > newX {
				currentX = currentX - 1
				if wiremap[currentX] == nil {
					wiremap[currentX] = make(map[int]WirePoint)
				}
				wp := wiremap[currentX][currentY]
				wp.Wire2HasHit = true
				wiremap[currentX][currentY] = wp
			} else if currentX < newX {
				currentX = currentX + 1
				if wiremap[currentX] == nil {
					wiremap[currentX] = make(map[int]WirePoint)
				}
				wp := wiremap[currentX][currentY]
				wp.Wire2HasHit = true
				wiremap[currentX][currentY] = wp
			} else {
				break
			}
		}

		for currentY != newY {
			if currentY > newY {
				currentY = currentY - 1
				wp := wiremap[currentX][currentY]
				wp.Wire2HasHit = true
				wiremap[currentX][currentY] = wp
			} else if currentY < newY {
				currentY = currentY + 1
				wp := wiremap[currentX][currentY]
				wp.Wire2HasHit = true
				wiremap[currentX][currentY] = wp
			} else {
				break
			}
		}
	}

	lowestdistance := 100000
	// The issue here is that I don't care if it's itself or another
	for x, map1 := range wiremap {
		for y, yval := range map1 {
			if yval.Wire1HasHit && yval.Wire2HasHit {
				currentDis := intAbs(x) + intAbs(y)
				if currentDis < lowestdistance {
					lowestdistance = currentDis
				}
			}
		}
	}

	fmt.Println("Lowest Distance:", lowestdistance)

	elapsed := time.Since(start)
	fmt.Println("Execution Time: %s", elapsed)
}

type WirePoint struct {
	Wire1HasHit bool
	Wire2HasHit bool
}

func intAbs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
