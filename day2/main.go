package day2

import (
	"strings"
)

// ComputeChecksum takes a list of strings and returns the compute hash
func ComputeChecksum(list []string) int {
	two := 0
	three := 0
	for _, n := range list {
		hasTwo, hasThree := CountDuplicates(n)
		if hasTwo {
			two++
		}
		if hasThree {
			three++
		}
	}

	return two * three
}

// CountDuplicates takes a argument boxID and returns the number of of
// letters that occur exactly 2 times and the number of letters that
// occur exactly 3 times as intagers
func CountDuplicates(boxID string) (bool, bool) {
	list := strings.Split(boxID, "")
	duplicates := [][]string{}
	for len(list) > 0 {
		first, remaining := list[0], list[1:]
		match := false
		for i := 0; i < len(duplicates); i++ {
			if duplicates[i][0] == first {
				duplicates[i] = append(duplicates[i], first)
				match = true
				break
			}
		}

		if match == false {
			duplicates = append(duplicates, []string{first})
		}
		list = remaining
	}

	return ParseThing(duplicates)
}

// ParseThing parses thing
func ParseThing(input [][]string) (bool, bool) {
	a, b := false, false
	for _, n := range input {
		if len(n) == 2 {
			a = true
		}
		if len(n) == 3 {
			b = true
		}
	}
	return a, b
}

// FindPrototype finds the comman values of the prototype
func FindPrototype(input []string) string {
	newlist := []string{}
	for _, n := range input {
		for _, u := range input {
			if n == u {
				newlist = append(newlist, n)
			}
			firstString := strings.Split(n, "")
			secondString := strings.Split(u, "")
			sharedString := ""
			for i, a := range firstString {
				if a == secondString[i] {
					sharedString += a
				}
			}
			newlist = append(newlist, sharedString)
		}
	}
	finalString := ""
	for _, item := range newlist {
		if len(item) == len(input[0])-1 {
			finalString = item
			break
		}
	}
	return finalString
}
