package piscine

func Split(s, sep string) []string {
	if sep == "" {
		// If separator is empty string, return a slice of each rune as string
		runes := []rune(s)
		result := make([]string, len(runes))
		for i, r := range runes {
			result[i] = string(r)
		}
		return result
	}

	var result []string
	start := 0
	for i := 0; i <= len(s)-len(sep); {
		if s[i:i+len(sep)] == sep {
			result = append(result, s[start:i])
			i += len(sep)
			start = i
		} else {
			i++
		}
	}
	// Add last part after last sep (or whole string if no sep found)
	result = append(result, s[start:])
	return result
}
