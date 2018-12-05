package main

import (
	"fmt"

	"github.com/thomasankcorn/adventofcode/day1"
	"github.com/thomasankcorn/adventofcode/day2"
)

func main() {
	list, err := FileToInts("./day1/data.txt")
	if err != nil {
		panic(err)
	}
	fmt.Printf("The answer to day 1 part 1 is %d\n", day1.UpdateFrequencyDrift(list))
	fmt.Printf("The answer to day 1 part 2 is %d\n", day1.DetectDuplicateFrequencies(list))

	day2data, err := FileToStrArr("./day2/data.txt")
	if err != nil {
		panic(err)
	}
	fmt.Printf("The answer to day 2 part 1 is %d\n", day2.ComputeChecksum(day2data))
}
