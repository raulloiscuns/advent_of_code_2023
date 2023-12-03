# Advent of Code 2023 :santa:

This repo covers some [Go](https://go.dev/) and [Julia](https://julialang.org/) answers to the [Advent of Code 2023](https://adventofcode.com/) problems. A brief summary of each of the statements of each problem is shown below.

1. [Trebuchet?!](#day-1-trebuchet)
2. [Cube Conundrum](#day-2-cube-conundrum)
3. [Gear Ratios](#day-3-gear-ratios)

## Day 1: Trebuchet?!

The input is a list of strings with letters and numbers. From each string is possible to obtain a number by taking the first and last number and combining it together (using the first one to build the tens and the last one to build the units). The goals of the problem are:
- Part 1: Calculate the sum of all the numbers that can be extracted from each string.
- Part 2: Now spelled numbers ("one", "two", ..., "eight" and "nine) are also treated as numbers and used to construct the number assigned to each string. Calculate the new sum of all the numbers.

## Day 2: Cube Conundrum

The input is a list of strings of games results. In the game (each game is identified by an ID number) cubes of different colors (red, green and blue) were extracted in each round (delimited by a semicolon-separated). The goals of the problem are:
- Part 1: If we only have 12 red, 13 green and 14 blues cubes; a game is only possible if the numbers of cubes in each round are not greater that the total number of cubes of each color. Calculate the sum of all the IDs of the possible games. 
- Part 2: We can define the power of each game as the product of the minimum number of cubes of each color that we need for the game to be possible. Calculate the sum of powers of all the games.

## Day 3: Gear Ratios

The input is a visual representation of an engine (engine schematic) by means of numbers and symbols.
- Part 1: If any number adjacent to a symbol ("." do not count as symbol) is considered a part number, compute the sum of all of the part numbers in the engine schematic.
- Part 2: If we define a gear as all the "*" symbols that are adjacent to exactly two part numbers and its gear ratio as the result of multiplying those two numbers together, compute the sum of all of the gear ratios in the engine schematic.