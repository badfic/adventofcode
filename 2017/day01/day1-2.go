package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string
	fmt.Scanln(&input)

	stepsForward := len(input) / 2
	var total int

	for i, r := range input {
		c := string(r)
		curr, _ := strconv.Atoi(c)
		halfwayAroundIdx := i + stepsForward

		if halfwayAroundIdx > len(input)-1 {
			halfwayAroundIdx = halfwayAroundIdx - len(input)
		}

		halfwayValue, _ := strconv.Atoi(string(input[halfwayAroundIdx]))
		if halfwayValue == curr {
			total += curr
		}
	}

	fmt.Println(total)
}
