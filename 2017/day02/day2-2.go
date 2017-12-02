package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("./day2.dat")

	defer file.Close()

	var checksum int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())

		for outerIdx, outer := range fields {
			outerField, _ := strconv.Atoi(outer)

			for innerIdx, inner := range fields {
				if innerIdx == outerIdx {
					continue
				}

				innerField, _ := strconv.Atoi(inner)

				if outerField%innerField == 0 {
					checksum += outerField / innerField
				}
			}
		}
	}

	fmt.Println(checksum)
}
