package util

// InitSliceInt set defaultValue to all element of uint slice
func InitSliceInt(slice []int, defaultValue int) {
	for i := range slice {
		slice[i] = defaultValue
	}
}

// InitSliceUint set defaultValue to all element of uint slice
func InitSliceUint(slice []uint, defaultValue uint) {
	for i := range slice {
		slice[i] = defaultValue
	}
}

// Compare two slice of uint and returns true if they
// contains same values in same order
func EqualsSliceUint(s1, s2 []uint) bool {
	if len(s1) != len(s2) {
		return false
	}
	for s1Index := range s1 {
		for s2Index := range s2 {
			if s1[s1Index] != s2[s2Index] {
				return false
			}
		}
	}
	return true
}
