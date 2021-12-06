package main

import (
	utils "adventofcode2021"
	"fmt"
	"strconv"
	"strings"
)

func doPart1(input []string) int64 {
	countZeros := make([]int, len(input[0]))
	for _, num := range input {
		for i, bit := range []rune(num) {
			char := string(bit)
			if char == "0" {
				countZeros[i]++
			}
		}
	}
	fmt.Println(countZeros)
	totalInput := len(input)

	gamma := make([]string, len(input[0]))
	epsilon := make([]string, len(input[0]))
	for i, num := range countZeros {
		if num > totalInput/2 {
			gamma[i] = "0"
			epsilon[i] = "1"
		} else {
			gamma[i] = "1"
			epsilon[i] = "0"
		}
	}
	fmt.Println(strings.Join(gamma, ""))
	fmt.Println(strings.Join(epsilon, ""))
	gammaInt, _ := strconv.ParseInt(strings.Join(gamma, ""), 2, 16)
	epsilonInt, _ := strconv.ParseInt(strings.Join(epsilon, ""), 2, 16)

	fmt.Printf("gamma= %d, epsilon= %d Final answer = %d\n", gammaInt, epsilonInt, gammaInt*epsilonInt)
	return gammaInt * epsilonInt
}

func doPart2(input []string) int64 {
	var (
		mostZerosC []string
		mostOnesC  []string
		mostZerosO []string
		mostOnesO  []string
	)
	oxygenRating := input[:]
	co2Rating := input[:]

	for i := range input[0] {
		if len(oxygenRating) > 1 {
			mostZerosO = mostZerosO[:0]
			mostOnesO = mostOnesO[:0]
			for _, val := range oxygenRating {
				if string(val[i]) == "0" {
					mostZerosO = append(mostZerosO, val)
				} else {
					mostOnesO = append(mostOnesO, val)
				}
			}
			if len(mostZerosO) > len(mostOnesO) {
				oxygenRating = mostZerosO[:]
			} else {
				oxygenRating = mostOnesO[:]
			}
		}
		if len(co2Rating) > 1 {
			mostZerosC = mostZerosC[:0]
			mostOnesC = mostOnesC[:0]
			for _, val := range co2Rating {
				if string(val[i]) == "0" {
					mostZerosC = append(mostZerosC, val)
				} else {
					mostOnesC = append(mostOnesC, val)
				}
			}
			if len(mostZerosC) > len(mostOnesC) {
				co2Rating = mostOnesC[:]
			} else {
				co2Rating = mostZerosC[:]
			}
		}
		fmt.Printf("oxygenRating %v \n", oxygenRating)
		fmt.Printf("co2Rating %v \n", co2Rating)
	}

	o2RatingInt, _ := strconv.ParseInt(oxygenRating[0], 2, 16)
	co2RatingInt, _ := strconv.ParseInt(co2Rating[0], 2, 16)

	fmt.Printf("o2Rating= %d, co2Rating= %d Final answer = %d\n", o2RatingInt, co2RatingInt, o2RatingInt*co2RatingInt)
	return o2RatingInt * co2RatingInt
}

func main() {
	fmt.Println("Advent of Code Day 2")
	testInput, _ := utils.ReadLines("tests/day3.txt")
	doPart2(testInput)
	//fmt.Println(moveSubmarine(testInput, 1))
	//input, _ := utils.ReadLines("inputs/day3.txt")
	//doPart2(input)
}
