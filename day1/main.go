package day1

// UpdateFrequencyDrift takes an argument of any number of ints and
// returns the total drift (sum) of all those elements
func UpdateFrequencyDrift(drifts []int) int {
	drift := 0
	for _, n := range drifts {
		drift += n
	}
	return drift
}

// DetectDuplicateFrequencies returns the first frequency that appears twice
func DetectDuplicateFrequencies(drifts []int) int {

	drift := 0
	list := []int{0}

	index := 0
	duplicateFound := false
	// loop through drifts until duplicate is found
	for !duplicateFound {
		if index >= len(drifts) {
			index = 0
		}

		drift += drifts[index]

		if indexOf(drift, list) != -1 {
			break
		}
		// append drift
		list = append(list, drift)
		index++
	}
	return drift
}

func indexOf(element int, data []int) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1 //not found.
}
