package main

import (
	"fmt"
	"strconv"
	"strings"
)

func parseInput(input string) []int {
	fishIn := strings.Split(input, ",")
	var fish []int
	for _, f := range fishIn {
		fishNum, _ := strconv.Atoi(f)
		fish = append(fish, fishNum)
	}
	return fish
}

func countFish(input []int, rounds int) int {
	for i := 0; i < rounds; i++ {
		//fmt.Printf("Move %d\n", i+1)
		currSize := len(input)
		for x := 0; x < currSize; x++ {
			if input[x] == 0 {
				input[x] = 6
				input = append(input, 8)
			} else {
				input[x]--
			}
		}
		//fmt.Println(input)
	}
	return len(input)
}

func smartCount(input[]int, rounds int) int {
	pond := make([]int, 9)
	for _, fish := range input {
		pond[fish]++
	}
	for i:=0; i<rounds; i++ {
		deadFish := pond[0]
		pond = pond[1:]
		pond[6]+=deadFish
		pond = append(pond, deadFish)
		//fmt.Println(pond)
	}
	totalFish := 0
	for _, fish := range pond {
		totalFish += fish
	}
	return totalFish
}

func main() {
	part1test := "3,4,3,1,2"
	part1input := "1,1,3,5,1,1,1,4,1,5,1,1,1,1,1,1,1,3,1,1,1,1,2,5,1,1,1,1,1,2,1,4,1,4,1,1,1,1,1,3,1,1,5,1,1,1,4,1,1,1,4,1,1,3,5,1,1,1,1,4,1,5,4,1,1,2,3,2,1,1,1,1,1,1,1,1,1,1,1,1,1,5,1,1,1,1,1,1,1,1,2,2,1,1,1,1,1,5,1,1,1,3,4,1,1,1,1,3,1,1,1,1,1,4,1,1,3,1,1,3,1,1,1,1,1,3,1,5,2,3,1,2,3,1,1,2,1,2,4,5,1,5,1,4,1,1,1,1,2,1,5,1,1,1,1,1,5,1,1,3,1,1,1,1,1,1,4,1,2,1,1,1,1,1,1,1,1,1,1,1,1,1,3,2,1,1,1,1,2,2,1,2,1,1,1,5,5,1,1,1,1,1,1,1,1,1,1,1,1,2,2,1,1,4,2,1,4,1,1,1,1,1,1,1,2,1,2,1,1,1,1,1,1,1,1,1,1,1,1,1,2,2,1,5,1,1,1,1,1,1,1,1,3,1,1,3,3,1,1,1,3,5,1,1,4,1,1,1,1,1,4,1,1,3,1,1,1,1,1,1,1,1,2,1,5,1,1,1,1,1,1,1,1,1,1,4,1,1,1,1"
	fmt.Println(smartCount(parseInput(part1test), 256))
	fmt.Println(smartCount(parseInput(part1input), 256))
}