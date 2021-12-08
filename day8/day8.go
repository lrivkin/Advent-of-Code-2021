package main

import (
	utils "adventofcode2021"
	"fmt"
	"strings"
)

func countUnique(output string) int {
	count := 0
	for _, val := range strings.Split(output, " ") {
		if len(val) == 2 || len(val) == 4 || len(val) == 3 || len(val) == 7 {
			count++
		}

	}
	return count
}

func main() {
	notes, _ := utils.ReadLines("inputs/day8.txt")
	runningSum := 0
	for _, note := range notes {
		//fmt.Println(note)
		x := strings.Split(note, " | ")
		//signal := x[0]
		output := x[1]
		runningSum += countUnique(output)
	}
	fmt.Println(runningSum)

}
