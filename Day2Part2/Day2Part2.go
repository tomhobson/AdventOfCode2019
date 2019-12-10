package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

func main() {
	addCode := "1"
	multiplyCode := "2"
	endCode := "99"

	start := time.Now()

	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	dataString := string(data)
	output := "FalseFlag"
	outputParm1 := "FalseFlag"
	outputParm2 := "FalseFlag"

	for x := 0; x <= 100; x++ {
		for y := 0; y <= 100; y++ {
			s := strings.Split(dataString, ",")
			s[1] = strconv.Itoa(x)
			s[2] = strconv.Itoa(y)
			for i := 0; i <= len(s); i = i + 4 {
				//Check endcode beforehand so we don't get index out of range exceptions
				if s[i] == endCode {
					break
				}
				loc1, _ := strconv.Atoi(s[i+1])
				loc2, _ := strconv.Atoi(s[i+2])
				val1, _ := strconv.Atoi(s[loc1])
				val2, _ := strconv.Atoi(s[loc2])
				storeLoc, _ := strconv.Atoi(s[i+3])
				if s[i] == addCode {
					add(val1, val2, storeLoc, s)
					continue
				}
				if s[i] == multiplyCode {
					multiply(val1, val2, storeLoc, s)
					continue
				}
			}

			output = s[0]
			if output == "19690720" {
				outputParm1 = s[1]
				outputParm2 = s[2]
				break
			}

		}
	}
	fmt.Println("s[1]:", outputParm1, "s[2]", outputParm2)
	intparm1, _ := strconv.Atoi(outputParm1)
	intparm2, _ := strconv.Atoi(outputParm2)
	fmt.Println("The answer is: ", 100*intparm1+intparm2)
	elapsed := time.Since(start)
	fmt.Println("Execution Time: %s", elapsed)
}

func add(inval1 int, inval2 int, storloc int, s []string) {
	s[storloc] = strconv.Itoa(inval1 + inval2)
}

func multiply(inval1 int, inval2 int, storloc int, s []string) {
	s[storloc] = strconv.Itoa(inval1 * inval2)
}
