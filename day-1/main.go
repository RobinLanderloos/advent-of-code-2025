package day1

import (
	"fmt"
	"log"
	"regexp"
	"robinlanderloos/aoc2025/io"
	"strconv"
)

func Day1() {
	fmt.Printf("Puzzle P1: %s \r\n", solveP1("day-1/input.txt"))
	fmt.Printf("Puzzle P2: %s \r\n", solveP2("day-1/input.txt"))
}

func solveP2(path string) string {
	pos := 50
	password := 0
	for line := range io.EnumerateFile(path) {
		instruction := parseInstruction(line)
		clicks := 0
		pos, clicks = calculatePosition(pos, instruction)
		password += clicks
	}

	return strconv.Itoa(password)
}

func solveP1(path string) string {
	pos := 50
	password := 0
	for line := range io.EnumerateFile(path) {
		instruction := parseInstruction(line)
		pos, _ = calculatePosition(pos, instruction)
		if pos == 0 {
			password += 1
		}
	}

	return strconv.Itoa(password)
}

func calculatePosition(pos, instruction int) (int, int) {
	clicks := 0
	if instruction < 0 {
		for range -instruction {
			if pos > 0 {
				pos -= 1
			} else {
				pos = 99
			}
			if pos == 0 {
				clicks += 1
			}
		}
	} else {
		for range instruction {
			if pos < 99 {
				pos += 1
			} else {
				pos = 0
			}
			if pos == 0 {
				clicks += 1
			}
		}
	}

	return pos, clicks
}

func parseInstruction(instruction string) int {
	regex, err := regexp.Compile(`(\w)(\d*)`)
	if err != nil {
		log.Fatalf("an error occurred while compiling the regex %s", err.Error())
	}

	instructionGroups := regex.FindStringSubmatch(instruction)

	if instructionGroups == nil {
		log.Fatalf("did not find an instruction in the provided string")
	}

	value, err := strconv.Atoi(instructionGroups[2])

	if err != nil {
		log.Fatalf("could not parse the provided value %s", err.Error())
	}

	if instructionGroups[1] == "L" {
		return 0 - value
	}

	return value
}
