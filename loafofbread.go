package piscine

func LoafOfBread(str string) string {
	if len(str) == 0 {
		return "\n"
	}
	if len(str) < 5 {
		return "Invalid Output\n"
	}
	returnStr := ""
	wordCount := 0
	for i, ch := range str {
		if ch == ' ' && wordCount != 5 {
			continue
		}
		if wordCount == 5 {
			if i < len(str)-1 && !(ch == ' ' && str[i+1] == ' ') {
				wordCount = 0
				returnStr += " "
			}
			continue
		}
		returnStr += string(ch)
		wordCount++
	}
	return returnStr + "\n"
}
