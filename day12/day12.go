package main

import (
	utils "adventofcode2021"
	"fmt"
	"strings"
)

type Cave struct {
	paths map[string]map[string]struct{}
}

func fillCave(input []string) Cave {
	paths := make(map[string]map[string]struct{})
	for _, i := range input {
		i2 := strings.Split(i, "-")
		start := i2[0]
		end := i2[1]
		_, ok := paths[start]
		if !ok {
			paths[start] = map[string]struct{}{end : {}}
		} else {
			paths[start][end] = struct{}{}
		}
		_, ok = paths[end]
		if !ok {
			paths[end] = map[string]struct{}{start : {}}
		} else {
			paths[end][start] = struct{}{}
		}
	}
	return Cave{paths}
}

func main() {
	things, _ := utils.ReadBlock("tests/day12.txt")
	test1 := things[0]
	c := fillCave(test1)
	for i := range c.paths {
		fmt.Printf("%s: %v\n", i, utils.GetMapKeys(c.paths[i]))
	}
	fmt.Println()
	test2 := things[1]
	c = fillCave(test2)
	for i := range c.paths {
		fmt.Printf("%s: %v\n", i, utils.GetMapKeys(c.paths[i]))
	}

	fmt.Println()
	input, _ := utils.ReadBlock("inputs/day12.txt")
	c = fillCave(input[0])
	for i := range c.paths {
		fmt.Printf("%s: %v\n", i, utils.GetMapKeys(c.paths[i]))
	}
}