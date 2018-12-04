package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

// FileToNums takes a text file and returns a list of numbers
func FileToNums(path string) ([]int, error) {
	b, err := ioutil.ReadFile(path) // just pass the file name
	if err != nil {
		return nil, err
	}
	// convert content to a 'string'
	str := string(b)
	// parses string to list of strings
	list := strings.Split(str, "\n")

	var ints = []int{}

	// loops through list of strings
	for _, i := range list {
		// parses each string to a number
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		// appends string to list
		ints = append(ints, j)
	}

	return ints, nil
}
