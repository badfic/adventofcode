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
	file, _ := os.Open("./day8.dat")

	defer file.Close()

	registers := map[string]int{}

	maxEverSeen := math.MinInt32
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input := scanner.Text()

		fields := strings.Fields(input)

		registerName := fields[0]
		op := fields[1]
		delta, _ := strconv.Atoi(fields[2])
		conditionRegister := fields[4]
		conditionalComparator := fields[5]
		conditionalValue, _ := strconv.Atoi(fields[6])

		if _, ok := registers[registerName]; !ok {
			registers[registerName] = 0
		}

		if _, ok := registers[conditionRegister]; !ok {
			registers[conditionRegister] = 0
		}

		if conditionalComparator == "<" {
			if registers[conditionRegister] < conditionalValue {
				modifyRegister(registers, registerName, op, delta)
			}
		} else if conditionalComparator == "<=" {
			if registers[conditionRegister] <= conditionalValue {
				modifyRegister(registers, registerName, op, delta)
			}
		} else if conditionalComparator == ">" {
			if registers[conditionRegister] > conditionalValue {
				modifyRegister(registers, registerName, op, delta)
			}
		} else if conditionalComparator == ">=" {
			if registers[conditionRegister] >= conditionalValue {
				modifyRegister(registers, registerName, op, delta)
			}
		} else if conditionalComparator == "==" {
			if registers[conditionRegister] == conditionalValue {
				modifyRegister(registers, registerName, op, delta)
			}
		} else if conditionalComparator == "!=" {
			if registers[conditionRegister] != conditionalValue {
				modifyRegister(registers, registerName, op, delta)
			}
		}

		if registers[registerName] > maxEverSeen {
			maxEverSeen = registers[registerName]
		}
	}

	max := math.MinInt32
	var maxName string
	for k, v := range registers {
		if v > max {
			max = v
			maxName = k
		}
	}

	fmt.Printf("Part 1: Largest register is %s with %d\n", maxName, max)
	fmt.Printf("Part 2: Largest value ever seen is %d\n", maxEverSeen)
}

func modifyRegister(registers map[string]int, name string, operation string, delta int) {
	if operation == "inc" {
		registers[name] = registers[name] + delta
	} else if operation == "dec" {
		registers[name] = registers[name] - delta
	}
}
