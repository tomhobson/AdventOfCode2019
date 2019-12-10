package main

import (
	"fmt"
	"io/ioutil"
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
	s := strings.Split(dataString, ",")
	fmt.Println(s)
	elapsed := time.Since(start)
	fmt.Println("Execution Time: %s", elapsed)
}
