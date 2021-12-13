package utils

import (
	"io/ioutil"
	"strconv"
	"strings"
)

// ReadLines reads a whole file into memory
// and returns a slice of its lines.
func ReadLines(path string) ([]string, error) {
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(contents), "\n"), nil
}

func StringSliceToIntSlice(input []string) ([]int64, error) {
	var numbers []int64
	for _, x := range input {
		num, err := strconv.ParseInt(x, 0, 64)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, num)
	}
	return numbers, nil
}

func ParseIntGrid(path string) ([][]int, error) {
	contents, err := ReadLines(path)
	if err != nil {
		return nil, err
	}

	var results [][]int
	for _, line := range contents {
		var row []int
		for _, x := range strings.Split(line, "") {
			num, _ := strconv.Atoi(x)
			row = append(row, num)
		}
		results = append(results, row)
	}
	return results, nil
}


func StringListToIntSlice(input string) ([]int, error) {
	var numbers []int
	separated := strings.Split(input, ",")
	for _, x := range separated {
		num, err := strconv.Atoi(strings.TrimSpace(x))
		if err != nil {
			return numbers, err
		}
		numbers = append(numbers, num)
	}
	return numbers, nil
}
