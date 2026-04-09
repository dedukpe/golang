package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	summary := make(map[string]int)
	word := ""

	for i := 0; i < len(str); i++ {
		if str[i] == ' ' || str[i] == '\t' || str[i] == '\n' || str[i] == '\r' {
			summary[word]++
			word = ""
		} else {
			word += string(str[i])
		}
	}

	summary[word]++
	return summary
}
