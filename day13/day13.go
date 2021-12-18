package main

import (
	//utils "adventofcode2021"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type Paper struct {
	dots map[Coordinate]struct{}
	maxX int
	maxY int
}

type Coordinate struct{
	x int
	y int
}

func (p *Paper) fillPaper(rawInput string) {
	for _, c := range strings.Split(rawInput, "\n") {
		splitCoordinate := strings.Split(c, ",")
		x, _ := strconv.Atoi(splitCoordinate[0])
		y, _ := strconv.Atoi(splitCoordinate[1])
		p.dots[Coordinate{x, y}] = struct{}{}
		if x > p.maxX {
			p.maxX = x
		}
		if y > p.maxY {
			p.maxY = y
		}
	}
}

func (p *Paper) print() {
	var paperGrid [][]string
	for i:=0; i < p.maxY+1; i++ {
		col := strings.Split(strings.Repeat(".", p.maxX+1), "")
		paperGrid = append(paperGrid, col)
	}
	for c := range p.dots {
		paperGrid[c.y][c.x] = "#"
	}
	for _, row := range paperGrid {
		fmt.Println(row)
	}
	fmt.Println()
}

func (p *Paper) fold(instruction string) {
	//fmt.Println(instruction)
	re := regexp.MustCompile(`fold along (x|y)=([0-9]+)`)
	parsedInstructions := re.FindStringSubmatch(instruction)
	direction := parsedInstructions[1]
	lineNum, _ := strconv.Atoi(parsedInstructions[2])

	switch direction {
	case "x": {
		for c := range p.dots {
			if c.x > lineNum {
				delete(p.dots, c)
				c.x = lineNum - (c.x-lineNum)
				p.dots[c] = struct{}{}
			}

		}
		p.maxX = lineNum-1
	}
	case "y": {
		for c := range p.dots {
			if c.y > lineNum {
				delete(p.dots, c)
				c.y = lineNum - (c.y-lineNum)
				p.dots[c] = struct{}{}
			}
		}
		p.maxY = lineNum-1
	}
	}
	//p.print()
	//fmt.Println(p.dots)
	//fmt.Println(len(p.dots))
	//fmt.Println()
}

func main() {
	contents, _ := ioutil.ReadFile("inputs/day13.txt")
	sliceContents := strings.Split(string(contents), "\n\n")
	paperInput := sliceContents[0]
	folding := sliceContents[1]

	paperGraphMap := make(map[Coordinate]struct{})
	paperGraph := Paper{paperGraphMap, 0, 0}
	paperGraph.fillPaper(paperInput)
	//fmt.Println(paperGraph.dots)
	//paperGraph.print()

	for _, fold := range strings.Split(folding, "\n"){
		paperGraph.fold(fold)
	}
	paperGraph.print()
}