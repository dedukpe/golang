package piscine

func Capitalize(s string) string {
	runes := []rune(s)
	newWord := true

	for i := 0; i < len(runes); i++ {
		c := runes[i]

		if isAlphaNumeric(c) {
			if newWord {
				runes[i] = toUpper(c)
				newWord = false
			} else {
				runes[i] = toLower(c)
			}
		} else {
			newWord = true
		}
	}

	return string(runes)
}

func isAlphaNumeric(r rune) bool {
	return (r >= 'A' && r <= 'Z') ||
		(r >= 'a' && r <= 'z') ||
		(r >= '0' && r <= '9')
}

func toUpper(r rune) rune {
	if r >= 'a' && r <= 'z' {
		return r - ('a' - 'A')
	}
	return r
}

func toLower(r rune) rune {
	if r >= 'A' && r <= 'Z' {
		return r + ('a' - 'A')
	}
	return r
}
