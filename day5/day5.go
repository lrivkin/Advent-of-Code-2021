package main

import (
	utils "adventofcode2021"
	"fmt"
	"strconv"
	"strings"
)

type Line struct {
	start Spot
	end   Spot
}

type Spot struct {
	x int
	y int
}

func parseInput(input []string) []Line {
	var lines []Line
	for _, l := range input {
		splitLine := strings.Split(l, " -> ")
		xy1 := strings.Split(splitLine[0], ",")
		x1, _ := strconv.Atoi(xy1[0])
		y1, _ := strconv.Atoi(xy1[1])

		xy2 := strings.Split(splitLine[1], ",")
		x2, _ := strconv.Atoi(xy2[0])
		y2, _ := strconv.Atoi(xy2[1])

		start := Spot{x1, y1}
		end := Spot{x2, y2}
		line := Line{start, end}
		lines = append(lines, line)
	}
	return lines
}
func fillMap(parsedLines []Line) map[Spot]int {
	spotMap := make(map[Spot]int)
	for _, line := range parsedLines {
		if line.start.x == line.end.x {
			y1 := line.start.y
			y2 := line.end.y
			bigY := y1
			smallY := y2
			if y2 > y1 {
				bigY = y2
				smallY = y1
			}
			for i := smallY; i <= bigY; i++ {
				spotMap[Spot{line.start.x, i}] += 1
			}
		} else if line.start.y == line.end.y {
			x1 := line.start.x
			x2 := line.end.x
			bigX := x1
			smallX := x2
			if x2 > x1 {
				bigX = x2
				smallX = x1
			}
			for i := smallX; i <= bigX; i++ {
				spotMap[Spot{i, line.start.y}] += 1
			}
		} else {
			// handle diagonals
			xStart := line.start.x
			yStart := line.start.y
			xEnd := line.end.x
			yEnd := line.end.y
			if xStart > xEnd {
				for i := 0; i <= xStart-xEnd; i++ {
					if yStart < yEnd {
						spotMap[Spot{xStart - i, yStart + i}] += 1
						//fmt.Println("a", xStart-i, yStart+i)
					} else {
						spotMap[Spot{xStart + i, yStart - i}] += 1
						//fmt.Println("b", xStart-i, yStart-i)
					}
				}
			} else {
				for i := 0; i <= xEnd-xStart; i++ {
					if yStart < yEnd {
						spotMap[Spot{xStart + i, yStart + i}] += 1
						//fmt.Println("c", xStart+i, yStart+i)
					} else {
						spotMap[Spot{xStart + i, yStart - i}] += 1
						//fmt.Println("d", xStart+i, yStart-i)
					}
				}
			}
		}
		//fmt.Println(line)
		//fmt.Println(spotMap)
	}
	//fmt.Println(spotMap)
	return spotMap
}

func main() {
	fmt.Println("Advent of code day 5")
	testLines, _ := utils.ReadLines("inputs/day5.txt")
	parsedLines := parseInput(testLines)

	//fmt.Println(spotMap)
	intersections := 0
	spotMap := fillMap(parsedLines)
	for _, count := range spotMap {
		if count > 1 {
			intersections++
		}
	}
	fmt.Println(intersections)
}
