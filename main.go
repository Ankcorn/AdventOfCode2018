package main

import (
	"fmt"

	"github.com/thomasankcorn/adventofcode/day1"
	"github.com/thomasankcorn/adventofcode/day2"
	"github.com/thomasankcorn/adventofcode/day3"
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
	fmt.Printf("The answer to day 2 part 2 is %s\n", day2.FindPrototype(day2data))

	day3data, err := FileToStrArr("./day3/data.txt")
	if err != nil {
		panic(err)
	}
	fmt.Printf("The answer to day 3 part 1 is %d\n", day3.Overlaps(day3data))
	fmt.Printf("The answer to day 3 part 2 is %d\n", day3.FindPerfectClaim(day3data))
}
