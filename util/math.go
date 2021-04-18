package util

// MaxInt compares int values and returns the maximum
func MaxInt(values ...int) int {
	var max int = values[0]
	for _, value := range values[1:] {
		if value > max {
			max = value
		}
	}
	return max
}

// MaxUint compares uint values and returns the maximum
func MaxUint(values ...uint) uint {
	var max uint = values[0]
	for _, value := range values[1:] {
		if value > max {
			max = value
		}
	}
	return max
}

// SumUint returns the sum of the uint values given
func SumUint(values ...uint) uint {
	var total uint = 0
	for _, value := range values {
		total += value
	}
	return total
}
