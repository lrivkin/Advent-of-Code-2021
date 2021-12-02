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
