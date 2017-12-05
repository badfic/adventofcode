package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, _ := os.Open("./day4.dat")

	defer file.Close()

	var numValid int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		set := map[string]bool{}
		dupes := false

		for _, field := range fields {
			if set[field] {
				dupes = true
				break
			}

			set[field] = true
		}

		if !dupes {
			numValid += 1
		}
	}

	fmt.Println(numValid)
}
