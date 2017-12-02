package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string
	fmt.Scanln(&input)

	var head int
	var prev int
	var curr int
	var total int

	for i, r := range input {
		prev = curr
		c := string(r)
		curr, _ = strconv.Atoi(c)

		if i == 0 {
			head = curr
		}

		if prev == curr {
			total += curr
		}
	}

	if curr == head {
		total += curr
	}

	fmt.Println(total)
}
