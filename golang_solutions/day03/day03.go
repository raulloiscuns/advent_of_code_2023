package day03

import (
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func checkSurrounded(nline int, match []int, lines []string) bool {
	notSymbols := []string{".", "1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
	surrounded := false
	for ncol := match[0]; ncol < match[1]; ncol++ {
		switch {
		case nline > 0 && !slices.Contains(notSymbols, string(lines[nline-1][ncol])):
			surrounded = true
		case nline < len(lines)-1 && !slices.Contains(notSymbols, string(lines[nline+1][ncol])):
			surrounded = true
		case ncol > 0 && !slices.Contains(notSymbols, string(lines[nline][ncol-1])):
			surrounded = true
		case ncol < len(lines[nline])-1 && !slices.Contains(notSymbols, string(lines[nline][ncol+1])):
			surrounded = true
		case nline > 0 && ncol > 0 && !slices.Contains(notSymbols, string(lines[nline-1][ncol-1])):
			surrounded = true
		case nline > 0 && ncol < len(lines[nline])-1 && !slices.Contains(notSymbols, string(lines[nline-1][ncol+1])):
			surrounded = true
		case nline < len(lines)-1 && ncol > 0 && !slices.Contains(notSymbols, string(lines[nline+1][ncol-1])):
			surrounded = true
		case nline < len(lines)-1 && ncol < len(lines[nline])-1 && !slices.Contains(notSymbols, string(lines[nline+1][ncol+1])):
			surrounded = true
		}

	}
	return surrounded
}

func checkGear(nline int, match []int, lines []string) (bool, [][2]int) {
	notSymbols := []string{".", "1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
	gear := false
	pos := [][2]int{}
	for ncol := match[0]; ncol < match[1]; ncol++ {
		switch {
		case nline > 0 && string(lines[nline-1][ncol]) == "*":
			gear = true
			pos = append(pos, [2]int{nline - 1, ncol})
		case nline < len(lines)-1 && string(lines[nline+1][ncol]) == "*":
			gear = true
			pos = append(pos, [2]int{nline + 1, ncol})
		case ncol > 0 && string(lines[nline][ncol-1]) == "*":
			gear = true
			pos = append(pos, [2]int{nline, ncol - 1})
		case ncol < len(lines[nline])-1 && string(lines[nline][ncol+1]) == "*":
			gear = true
			pos = append(pos, [2]int{nline, ncol + 1})
		case nline > 0 && ncol > 0 && !slices.Contains(notSymbols, string(lines[nline-1][ncol-1])):
			gear = true
			pos = append(pos, [2]int{nline - 1, ncol - 1})
		case nline > 0 && ncol < len(lines[nline])-1 && !slices.Contains(notSymbols, string(lines[nline-1][ncol+1])):
			gear = true
			pos = append(pos, [2]int{nline - 1, ncol + 1})
		case nline < len(lines)-1 && ncol > 0 && !slices.Contains(notSymbols, string(lines[nline+1][ncol-1])):
			gear = true
			pos = append(pos, [2]int{nline + 1, ncol - 1})
		case nline < len(lines)-1 && ncol < len(lines[nline])-1 && !slices.Contains(notSymbols, string(lines[nline+1][ncol+1])):
			gear = true
			pos = append(pos, [2]int{nline + 1, ncol + 1})
		}
		if gear == true {
			break
		}
	}
	return gear, pos
}

func Part1(filename string) int {

	input, err := os.ReadFile(filename)

	if err != nil {
		log.Fatal("File not found")
	}

	lines := strings.Split(string(input), "\n")

	sum := 0
	for i, line := range lines {
		matches := regexp.MustCompile(`(\d+)`).FindAllStringIndex(line, -1)
		for _, match := range matches {
			if checkSurrounded(i, match, lines) {
				number, _ := strconv.Atoi(line[match[0]:match[1]])
				sum = sum + number
			}
		}
	}

	return sum

}

func Part2(filename string) int {

	input, err := os.ReadFile(filename)

	if err != nil {
		log.Fatal("File not found")
	}

	lines := strings.Split(string(input), "\n")

	gearsDict := make(map[[2]int][]int)
	for i, line := range lines {
		matches := regexp.MustCompile(`(\d+)`).FindAllStringIndex(line, -1)
		for _, match := range matches {
			check, gearPositions := checkGear(i, match, lines)
			if check {
				number, _ := strconv.Atoi(line[match[0]:match[1]])
				for _, position := range gearPositions {
					gearsDict[position] = append(gearsDict[position], number)
				}
			}
		}
	}

	sum := 0
	for _, gear := range gearsDict {
		if len(gear) == 2 {
			sum = sum + gear[0]*gear[1]
		}
	}

	return sum

}
