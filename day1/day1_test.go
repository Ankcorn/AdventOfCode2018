package day1

import (
	"testing"
)

func TestUpdateFrequencyDrift1(t *testing.T) {
	input := []int{0, 1}
	expected := 1
	actual := UpdateFrequencyDrift(input)

	if actual != expected {
		t.Errorf("Expected %d but instead got %d!", expected, actual)
	}
}

func TestUpdateFrequencyDrift2(t *testing.T) {
	input := []int{1, 1, -2}
	expected := 0
	actual := UpdateFrequencyDrift(input)

	if actual != expected {
		t.Errorf("Expected %d but instead got %d!", expected, actual)
	}
}

func TestDuplicateFrequencyDrift1(t *testing.T) {
	input := []int{1, -1, 3}
	expected := 0
	actual := DetectDuplicateFrequencies(input)

	if actual != expected {
		t.Errorf("Expected %d but instead got %d!", expected, actual)
	}
}

func TestDuplicateFrequencyDrift2(t *testing.T) {
	input := []int{+7, +7, -2, -7, -4}
	expected := 14
	actual := DetectDuplicateFrequencies(input)

	if actual != expected {
		t.Errorf("Expected %d but instead got %d!", expected, actual)
	}
}

func TestIndexOfTest(t *testing.T) {
	input := []int{7, 7, -2, -7, -4}

	expected := 2
	actual := indexOf(-2, input)

	if actual != expected {
		t.Errorf("Expected %d but instead got %d!", expected, actual)
	}
}
