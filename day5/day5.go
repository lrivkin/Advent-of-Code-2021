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

func absValue(num int) int {
	if num < 0 {
		return num * -1
	}
	return num
}

func fillMap(parsedLines []Line) map[Spot]int {
	spotMap := make(map[Spot]int)
	for _, line := range parsedLines {
		//fmt.Println(line)
		//fmt.Println(spotMap)
		xStart := line.start.x
		xEnd := line.end.x
		yStart := line.start.y
		yEnd := line.end.y
		deltaX := xEnd - xStart
		deltaY := yEnd - yStart
		if deltaX == 0 {
			direction := deltaY / absValue(deltaY)
			for i := 0; i <= absValue(deltaY); i++ {
				spotMap[Spot{xStart, yStart + direction*i}]++
			}
		} else if deltaY == 0 {
			direction := deltaX / absValue(deltaX)
			for i := 0; i <= absValue(deltaX); i++ {
				spotMap[Spot{xStart + direction*i, yStart}]++
			}
		} else {
			xDirection := deltaX / absValue(deltaX)
			yDirection := deltaY / absValue(deltaY)
			for i := 0; i <= absValue(deltaY); i++ {
				spotMap[Spot{xStart + xDirection*i, yStart + yDirection*i}]++
			}
		}
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
