package piscine

func SplitWhiteSpaces(s string) []string {
	var words []string
	start := -1 // start index of a word, -1 means not inside a word

	for i, r := range s {
		if r == ' ' || r == '\t' || r == '\n' {
			// if we were inside a word, slice it and append
			if start != -1 {
				words = append(words, s[start:i])
				start = -1
			}
		} else {
			// if not inside a word, mark start
			if start == -1 {
				start = i
			}
		}
	}

	// add last word if string ended without whitespace
	if start != -1 {
		words = append(words, s[start:])
	}

	return words
}
