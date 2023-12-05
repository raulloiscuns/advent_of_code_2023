package day04

import (
	"bufio"
	"log"
	"math"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func Part1(filename string) int {
	cardData := readInput(filename)

	sum := 0
	for _, card := range cardData {
		count := 0
		for _, number := range card[1] {
			if slices.Contains(card[0], number) {
				count++
			}
		}
		if count > 0 {
			sum = sum + int(math.Pow(2, float64(count-1)))
		}
	}

	return sum
}

func Part2(filename string) int {
	cardData := readInput(filename)

	numberOfCards := make([]int, len(cardData))
	for i := range numberOfCards {
		numberOfCards[i] = 1
	}

	for index, card := range cardData {
		for i := 0; i < numberOfCards[index]; i++ {
			count := 0
			for _, number := range card[1] {
				if slices.Contains(card[0], number) {
					count++
				}
			}
			for j := index + 1; j < index+count+1; j++ {
				numberOfCards[j]++

			}
		}

	}

	sum := 0
	for _, n := range numberOfCards {
		sum = sum + n
	}

	return sum
}

func readInput(filename string) [][][]int {

	// Open input file
	readFile, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	cardData := make([][][]int, 0)
	numbersRegex := regexp.MustCompile(`(\d+)`)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		splitedLine := strings.Split(line, ":")
		splitedNumbers := strings.Split(splitedLine[1], "|")

		cardNumbers := make([][]int, 2)

		// Extracting winning numbers
		winningNumbersMatches := numbersRegex.FindAllStringSubmatch(splitedNumbers[0], -1)
		winningNumbers := make([]int, len(winningNumbersMatches))
		for i, match := range winningNumbersMatches {
			number, _ := strconv.Atoi(match[1])
			winningNumbers[i] = number
		}
		cardNumbers[0] = winningNumbers

		// Extracting player numbers
		playerNumbersMatches := numbersRegex.FindAllStringSubmatch(splitedNumbers[1], -1)
		playerNumbers := make([]int, len(playerNumbersMatches))
		for i, match := range playerNumbersMatches {
			number, _ := strconv.Atoi(match[1])
			playerNumbers[i] = number
		}
		cardNumbers[1] = playerNumbers

		cardData = append(cardData, cardNumbers)
	}

	return cardData

}
