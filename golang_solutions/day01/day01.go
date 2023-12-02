package day01

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func common(filename string) []string {

	// Read input file

	input, err := os.ReadFile(filename)

	if err != nil {
		log.Fatal("File not found")
	}

	lines := strings.Split(string(input), "\n")

	return lines
}

func Part1(filename string) int {

	lines := common(filename)

	sum := 0
	for _, line := range lines {

		chars := strings.Split(line, "")

		numbers := []int{}
		for _, char := range chars {
			number, err := strconv.Atoi(char)

			if err != nil {
				continue
			}

			numbers = append(numbers, number)

		}

		sum = sum + 10*numbers[0] + numbers[len(numbers)-1]
	}

	return sum

}

func Part2(filename string) int {

	lines := common(filename)

	sum := 0
	for _, line := range lines {

		chars := strings.Split(line, "")

		numbers := []int{}
		for index, char := range chars {
			number, err := strconv.Atoi(char)

			switch {
			case err == nil:
				numbers = append(numbers, number)
			case strings.HasPrefix(line[index:], "one"):
				numbers = append(numbers, 1)
			case strings.HasPrefix(line[index:], "two"):
				numbers = append(numbers, 2)
			case strings.HasPrefix(line[index:], "three"):
				numbers = append(numbers, 3)
			case strings.HasPrefix(line[index:], "four"):
				numbers = append(numbers, 4)
			case strings.HasPrefix(line[index:], "five"):
				numbers = append(numbers, 5)
			case strings.HasPrefix(line[index:], "six"):
				numbers = append(numbers, 6)
			case strings.HasPrefix(line[index:], "seven"):
				numbers = append(numbers, 7)
			case strings.HasPrefix(line[index:], "eight"):
				numbers = append(numbers, 8)
			case strings.HasPrefix(line[index:], "nine"):
				numbers = append(numbers, 9)
			}

		}

		sum = sum + 10*numbers[0] + numbers[len(numbers)-1]
	}

	return sum

}
