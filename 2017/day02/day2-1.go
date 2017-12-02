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
	file, _ := os.Open("./day2-1.dat")

	defer file.Close()

	var checksum int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		min := math.MaxInt32
		max := math.MinInt32

		fields := strings.Fields(scanner.Text())

		for _, field := range fields {
			intField, _ := strconv.Atoi(field)

			if intField < min {
				min = intField
			}

			if intField > max {
				max = intField
			}
		}

		checksum += max - min
	}

	fmt.Println(checksum)
}
