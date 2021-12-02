package main

import (
	"fmt"
)

func part1(measurements []int64) int {
	count := 0
	for i := 0; i < len(measurements) - 1; i+=1 {
		if measurements[i+1] > measurements[i] {
			count+=1
		}
	}
	return count
}

func part2Result(measurements []int64) int {
	windows := make([]int64, len(measurements)-2)
	for i := range windows {
		windows[i] = measurements[i] + measurements[i+1] + measurements[i+2]
	}
	return part1(windows)
}

func main() {
	content, _ := readLines("inputs/day1-input.txt")
	//fmt.Println(content)
	measurements, _ := stringSliceToIntSlice(content)
	//fmt.Println(numbers)
	part1Result := part1(measurements)
	fmt.Printf("Part 1: Total increased = %v\n", part1Result)
	fmt.Printf("Part 2: With sliding windows = %v", part2Result(measurements))

}