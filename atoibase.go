package piscine

func isValidBasePrintNbr(base string) bool {
	if len(base) < 2 {
		return false
	}
	seen := make(map[rune]bool)
	for _, r := range base {
		if r == '+' || r == '-' || seen[r] {
			return false
		}
		seen[r] = true
	}
	return true
}

func AtoiBase(s string, base string) int {
	if !isValidBasePrintNbr(base) {
		return 0
	}

	baseLen := len(base)
	value := 0

	for _, r := range s {
		digitValue := -1
		for i, br := range base {
			if br == r {
				digitValue = i
				break
			}
		}
		if digitValue == -1 {
			return 0
		}
		value = value*baseLen + digitValue
	}

	return value
}
