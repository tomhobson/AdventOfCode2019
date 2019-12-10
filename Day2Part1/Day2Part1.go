package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	start := time.Now()
	addCode := "1"
	multiplyCode := "2"
	endCode := "99"

	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	dataString := string(data)

	s := strings.Split(dataString, ",")

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

	fmt.Println("The answer is: ", s[0])
	elapsed := time.Since(start)
	fmt.Println("Execution Time: %s", elapsed)
}

func add(inval1 int, inval2 int, storloc int, s []string) {
	s[storloc] = strconv.Itoa(inval1 + inval2)
}

func multiply(inval1 int, inval2 int, storloc int, s []string) {
	s[storloc] = strconv.Itoa(inval1 * inval2)
}
