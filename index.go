package piscine

func Index(s string, toFind string) int {
	if toFind == "" {
		return 0
	}

	sLen := len(s)
	tLen := len(toFind)

	if tLen > sLen {
		return -1
	}

	for i := 0; i <= sLen-tLen; i++ {
		if s[i:i+tLen] == toFind {
			return i
		}
	}
	return -1
}
