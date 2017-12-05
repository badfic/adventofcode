package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	file, _ := os.Open("./day4.dat")

	defer file.Close()

	var numValid int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		valid := true

		dictionary := map[string]bool{}

		for _, field := range fields {
			sorted := SortString(field)

			if dictionary[sorted] {
				valid = false
				break
			}

			dictionary[sorted] = true
		}

		if valid {
			numValid += 1
		}
	}

	fmt.Println(numValid)
}

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
