package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	fromLen := len(baseFrom)
	toLen := len(baseTo)

	// Step 1: Convert nbr (baseFrom) to integer
	value := 0
	for _, digit := range nbr {
		value = value*fromLen + indexInBase(digit, baseFrom)
	}

	// Step 2: Convert integer value to baseTo string
	if value == 0 {
		return string(baseTo[0])
	}

	result := ""
	for value > 0 {
		remainder := value % toLen
		result = string(baseTo[remainder]) + result
		value /= toLen
	}

	return result
}

// Helper: find index of rune in base string
func indexInBase(r rune, base string) int {
	for i, b := range base {
		if b == r {
			return i
		}
	}
	return -1 // Should never happen if input is valid
}
