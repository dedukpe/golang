package piscine

func JumpOver(str string) string {
	runes := []rune(str)

	// Count runes manually (avoid len)
	length := 0
	for range runes {
		length++
	}

	// If empty or fewer than 3 characters
	if length < 3 {
		return "\n"
	}

	result := ""
	for i := 2; i < length; i += 3 {
		result += string(runes[i])
	}

	// Return every third character + newline
	return result + "\n"
}
