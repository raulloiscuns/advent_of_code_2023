package day02

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

func common(filename string) map[string][3][]int {

	// Read input file
	readFile, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	indexRegex := regexp.MustCompile(`Game (\d+):`)
	redRegex := regexp.MustCompile(`(\d+) red`)
	greenRegex := regexp.MustCompile(`(\d+) green`)
	blueRegex := regexp.MustCompile(`(\d+) blue`)

	gameDict := make(map[string][3][]int)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		gameIndex := indexRegex.FindStringSubmatch(line)[1]

		colorTrack := [3][]int{}
		redCubes := redRegex.FindAllStringSubmatch(line, -1)
		if len(redCubes) > 0 {
			for _, result := range redCubes {
				number, _ := strconv.Atoi(result[1])
				colorTrack[0] = append(colorTrack[0], number)
			}
		}
		greenCubes := greenRegex.FindAllStringSubmatch(line, -1)
		if len(greenCubes) > 0 {
			for _, result := range greenCubes {
				number, _ := strconv.Atoi(result[1])
				colorTrack[1] = append(colorTrack[1], number)
			}
		}
		blueCubes := blueRegex.FindAllStringSubmatch(line, -1)
		if len(blueCubes) > 0 {
			for _, result := range blueCubes {
				number, _ := strconv.Atoi(result[1])
				colorTrack[2] = append(colorTrack[2], number)
			}
		}

		gameDict[gameIndex] = colorTrack
	}

	readFile.Close()

	return gameDict

}

func anyHigher(slice []int, num int) bool {
	for _, val := range slice {
		if val > num {
			return true
		}
	}
	return false
}

func max(slice []int) int {
	max := 0
	for _, val := range slice {
		if val > max {
			max = val
		}
	}

	return max
}

func Part1(filename string) int {
	gameData := common(filename)

	sum := 0
	for index, game := range gameData {
		if !(anyHigher(game[0], 12) || anyHigher(game[1], 13) || anyHigher(game[2], 14)) {
			intIndex, _ := strconv.Atoi(index)
			sum = sum + intIndex
		}

	}

	return sum
}

func Part2(filename string) int {
	gameData := common(filename)

	sum := 0
	for _, game := range gameData {
		sum = sum + max(game[0])*max(game[1])*max(game[2])
	}

	return sum
}
