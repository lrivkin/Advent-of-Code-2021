package main

import (
	utils "adventofcode2021"
	"fmt"
)

type Grid struct {
	numFlashes int
	grid [][]int
}

type Coordinate struct {
	x int
	y int
}

func (g *Grid) flash(c Coordinate) []Coordinate {
	var nextFlash []Coordinate
	//fmt.Printf("Flashing %v\n", c)
	for i:=-1; i<=1; i++ {
		for j:=-1; j<=1; j++ {
			if i==0 && j==0 {
				continue
			}
			if c.x+i >= 0 && c.x+i < len(g.grid) && c.y+j >= 0 && c.y+j < len(g.grid[0]) {
				if g.grid[c.x+i][c.y+j] == 9 {
					nextFlash = append(nextFlash, Coordinate{c.x+i, c.y+j})
				}
				if g.grid[c.x+i][c.y+j] <= 9 {
					g.grid[c.x+i][c.y+j]++
				}
			}
		}
	}
	//g.print()
	return nextFlash
}

func (g *Grid) print() {
	for i := 0; i < len(g.grid); i++ {
		for j := 0; j < len(g.grid[i]); j++ {
			if g.grid[i][j] > 9 {
				fmt.Print("*")
			} else {
				fmt.Print(g.grid[i][j])
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func (g *Grid) playRound() bool {
	flashed := make(map[Coordinate]bool)
	needsFlash := 0

	// First, the energy level of each octopus increases by 1.
	for i := 0; i < len(g.grid); i++ {
		for j:= 0; j < len(g.grid[0]); j++ {
			g.grid[i][j]++
			if g.grid[i][j] > 9 {
				flashed[Coordinate{i, j}] = false
				needsFlash++
			}
		}
	}
	for needsFlash > 0 {
		for i := range flashed {
			if flashed[i] {
				// we've already flashed this, skip it
				continue
			}
			nextFlash := g.flash(i)
			needsFlash--
			for _, c := range nextFlash {
				// queue this spot up to be flashed next if it hasn't been
				_, ok := flashed[c]
				if !ok {
					flashed[c] = false
					needsFlash++
				}
			}
			flashed[i] = true
		}
	}
	if len(flashed) == len(g.grid)*len(g.grid[0]) {
		return true
	}
	// Finally, any octopus that flashed during this step has its energy level set to 0
	for i := range flashed {
		g.grid[i.x][i.y] = 0
		g.numFlashes++
	}
	return false
}

func main() {
	gridIn, _ := utils.ParseIntGrid("inputs/day11.txt")
	grid := Grid{0, gridIn}
	grid.print()

	numRounds := 0
	for i := 0; i< 100; i++ {
		grid.playRound()
		//grid.print()
		numRounds++
	}
	grid.print()
	fmt.Println(grid.numFlashes)
	allFlashed := false
	for !allFlashed {
		allFlashed = grid.playRound()
		numRounds++
	}
	fmt.Println(numRounds)
	//grid.print()
	//grid.playRound()
	//grid.print()
	//grid.playRound()
	//grid.print()

}