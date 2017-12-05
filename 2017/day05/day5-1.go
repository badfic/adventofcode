package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("./day5.dat")

	defer file.Close()

	list := []int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parsed, _ := strconv.Atoi(scanner.Text())
		list = append(list, parsed)
	}

	steps := 0
	idx := 0
	for {
		if idx < 0 || idx >= len(list) {
			break
		}

		curr := list[idx]
		list[idx] = curr + 1

		idx += curr
		steps += 1
	}

	fmt.Println(steps)
}
