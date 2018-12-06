package day2

import "testing"
import "strconv"

func TestComputeChecksum(t *testing.T) {
	input := []string{
		"abcdef",
		"bababc",
		"abbcde",
		"abcccd",
		"aabcdd",
		"abcdee",
		"ababab",
	}
	expected := 12
	actual := ComputeChecksum(input)

	if expected != actual {
		t.Errorf(
			"Expected %d but instead got %d!",
			expected,
			actual,
		)
	}
}

func TestCountDuplicates(t *testing.T) {
	input := "abbcde"
	two := true
	three := false
	actualtwo, actualthree := CountDuplicates(input)

	if two != actualtwo || three != actualthree {
		t.Errorf(
			"Expected %s and %s but instead got %s and %s!",
			strconv.FormatBool(two),
			strconv.FormatBool(three),
			strconv.FormatBool(actualtwo),
			strconv.FormatBool(actualthree),
		)
	}
}

func TestCountDuplicates1(t *testing.T) {
	input := "abcdef"
	two := false
	three := false
	actualtwo, actualthree := CountDuplicates(input)

	if two != actualtwo || three != actualthree {
		t.Errorf(
			"Expected %s and %s but instead got %s and %s!",
			strconv.FormatBool(two),
			strconv.FormatBool(three),
			strconv.FormatBool(actualtwo),
			strconv.FormatBool(actualthree),
		)
	}
}

func TestParseThingy(t *testing.T) {
	input := [][]string{
		{"a", "a", "a"},
		{"b", "b"},
		{"c"},
	}
	two := true
	three := true

	actualtwo, actualthree := ParseThing(input)

	if two != actualtwo || three != actualthree {
		t.Errorf(
			"Expected %s and %s but instead got %s and %s!",
			strconv.FormatBool(two),
			strconv.FormatBool(three),
			strconv.FormatBool(actualtwo),
			strconv.FormatBool(actualthree),
		)
	}
}

func TestPart2(t *testing.T) {
	input := []string{
		"abcde",
		"fghij",
		"klmno",
		"pqrst",
		"fguij",
		"axcye",
		"wvxyz",
	}
	expected := "fgij"
	actual := FindPrototype(input)

	if expected != actual {
		t.Errorf("Expected %s to equal %s!", expected, actual)
	}
}
