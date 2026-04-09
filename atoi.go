package piscine

func Atoi(s string) int {
	if s == "" {
		return 0
	}

	i := 0
	sign := 1

	if s[0] == '+' {
		i = 1
	} else if s[0] == '-' {
		sign = -1
		i = 1
	}

	if i < len(s) && (s[i] == '+' || s[i] == '-') {
		return 0
	}

	result := 0
	for ; i < len(s); i++ {
		ch := s[i]
		if ch < '0' || ch > '9' {
			return 0
		}
		result = result*10 + int(ch-'0')
	}

	return sign * result
}
