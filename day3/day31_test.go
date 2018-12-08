package day3

import (
	"testing"
)

func TestFindPerfectClaim(t *testing.T) {
	input := []string{
		"#1 @ 1,3: 4x4",
		"#2 @ 3,1: 4x4",
		"#3 @ 5,5: 2x2",
	}
	expected := 3
	actual := FindPerfectClaim(input)

	if expected != actual {
		t.Errorf("Expected %d to equal %d!", expected, actual)
	}
}
