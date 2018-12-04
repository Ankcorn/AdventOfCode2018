package main

import (
	"fmt"

	"github.com/thomasankcorn/adventofcode/day1"
)

func main() {
	list, err := FileToNums("./day1/data.txt")
	if err != nil {
		panic(err)
	}
	fmt.Printf("The answer to day 1 is %d\n", day1.UpdateFrequencyDrift(list))
	fmt.Printf("The answer to day 1 part 2 is %d\n", day1.DetectDuplicateFrequencies(list))
}
