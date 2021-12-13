package main

import (
	utils "adventofcode2021"
	"fmt"
)

type Lava struct {
	lava [][]int
}

func (l Lava) adjacentSpots(x int, y int) []int {
	if x < 0 || y < 0 || x > len(l.lava) || y > len(l.lava[0]) {
		return nil
	}
	var spots []int
	for r:=-1; r<=1; r++ {
		for c:=-1; c<=1; c++ {
			if r==0 && c==0 {
				continue
			}
			if x+r >= 0 && x+r < len(l.lava) && y+c >= 0 && y+c < len(l.lava[0]) {
				spots = append(spots, l.lava[x+r][y+c])
			}
		}
	}
	return spots
}

func (l Lava) isLowest(x int, y int) bool {
	adjacent := l.adjacentSpots(x, y)
	val := l.lava[x][y]
	for _, i := range adjacent {
		if val >= i {
			return false
		}
	}
	return true
}

func (l Lava) findRisk() int {
	lava := l.lava
	riskLevel := 0
	for r := 0; r < len(lava); r++ {
		for c := 0; c < len(lava[0]); c++ {
			if l.isLowest(r, c) {
				riskLevel += lava[r][c]+1
			}
		}
	}
	return riskLevel
}

func main() {
	inputs, _ := utils.ParseIntGrid("inputs/day9.txt")
	//fmt.Println(inputs)
	lava := Lava{inputs}
	fmt.Println(lava.findRisk())
}