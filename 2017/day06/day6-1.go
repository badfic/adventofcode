package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("./day6.dat")

	defer file.Close()

	list := []int{}
	layouts := map[string]bool{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input := scanner.Text()

		for _, field := range strings.Fields(input) {
			parsed, _ := strconv.Atoi(field)
			list = append(list, parsed)
		}
	}

	steps := 0
	for {
		stringRepresentation := strings.Join(strings.Fields(fmt.Sprint(list)), " ")

		if layouts[stringRepresentation] {
			break
		}

		layouts[stringRepresentation] = true

		idx := 0
		max := math.MinInt32
		for maxIdx, val := range list {
			if val > max {
				max = val
				idx = maxIdx
			}
		}

		val := list[idx]
		list[idx] = 0

		for i := 0; i < val; i++ {
			next := (idx + i + 1) % len(list)

			list[next] = list[next] + 1
		}

		steps += 1
	}

	fmt.Println(steps)
}
