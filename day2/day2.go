package main

import (
	utils "adventofcode2021"
	"fmt"
	"strconv"
	"strings"
)


func moveSubmarine(input []string, part int) int {
	var depth = 0
	var horiz = 0
	aim := 0
	for _, line := range input {
		splitLine := strings.Split(line, " ")
		dir := splitLine[0]
		moves, _ := strconv.Atoi(splitLine[1])
		if part == 1 {
			switch dir {
			case "forward":
				horiz += moves
			case "down":
				depth += moves
			case "up":
				depth -= moves
			}
		} else {
			switch dir {
			case "forward":
				horiz += moves
				depth += aim * moves
			case "down":
				aim += moves
			case "up":
				aim -= moves
			}
		}
		//fmt.Printf("Move %v, new position = %v, %v\n", dir, horiz, depth)
	}
	return depth * horiz
}


func main() {
	fmt.Println("Advent of Code Day 2")
	testInput, _ := utils.ReadLines("tests/day2.txt")
	fmt.Println(moveSubmarine(testInput, 1))
	input, _ := utils.ReadLines("inputs/day2.txt")
	fmt.Println(moveSubmarine(input, 1))

	fmt.Println("Part 2")
	testInput, _ = utils.ReadLines("tests/day2.txt")
	fmt.Println(moveSubmarine(testInput, 2))
	input, _ = utils.ReadLines("inputs/day2.txt")
	fmt.Println(moveSubmarine(input, 2))
}
