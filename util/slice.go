package util

// InitSliceInt set defaultValue to all element of uint slice
func InitSliceInt(slice []int, defaultValue int) {
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
	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}
