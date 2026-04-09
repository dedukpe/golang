package piscine

func IsNumeric(s string) bool {
	if len(s) == 0 {
		return false // Usually empty string is considered *not* numeric
	}
	for i := 0; i < len(s); i++ {
		c := s[i]
		if !(c >= '0' && c <= '9') {
			return false
		}
	}
	return true
}
