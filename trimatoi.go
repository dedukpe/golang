package piscine

func TrimAtoi(s string) int {
	sign := 1
	result := 0
	signFound := false
	digitFound := false

	for i := 0; i < len(s); i++ {
		c := s[i]

		if c == '-' && !digitFound && !signFound {
			// Only allow '-' if no digit/sign found yet
			sign = -1
			signFound = true
		} else if c >= '0' && c <= '9' {
			digitFound = true
			result = result*10 + int(c-'0')
		}
	}

	return sign * result
}
