package piscine

func StringToIntSlice(str string) []int {
	if str == "" {
		return nil
	}

	result := make([]int, 0, len(str)) // preallocate max possible length
	for _, r := range str {
		result = append(result, int(r)) // use rune values, not bytes
	}
	return result
}
