package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/raulloiscuns/advent_of_code_2023/golang_solutions/day01"
	"github.com/raulloiscuns/advent_of_code_2023/golang_solutions/day02"
	"github.com/raulloiscuns/advent_of_code_2023/golang_solutions/day03"
)

func day() int {
	last := 3
	if len(os.Args) == 1 {
		return last
	}
	day, err := strconv.Atoi(os.Args[1])

	if err != nil || day < 0 || day > 25 {
		log.Fatal("Day ", os.Args[1], " is not a valid day.")
	}

	if day > last {
		log.Fatal("Day ", day, " has not yet been solved.")
	}

	return day
}

func main() {
	d := day() // Select day, if not last day is selected
	fmt.Printf("Running day %02d\n", d)

	// Filename (input in day<NN>/input.txt)
	filename := fmt.Sprintf("day%02d/input.txt", d)

	switch d {
	case 1:
		fmt.Printf("-> Part 1: %d\n", day01.Part1(filename))
		fmt.Printf("-> Part 2: %d\n", day01.Part2(filename))
	case 2:
		fmt.Printf("-> Part 1: %d\n", day02.Part1(filename))
		fmt.Printf("-> Part 2: %d\n", day02.Part2(filename))
	case 3:
		fmt.Printf("-> Part 1: %d\n", day03.Part1(filename))
		fmt.Printf("-> Part 2: %d\n", day03.Part2(filename))
	}

}
