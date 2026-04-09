package piscine

func ConcatParams(args []string) string {
	if len(args) == 0 {
		return ""
	}

	// Calculate total length to allocate buffer once
	totalLen := 0
	for _, s := range args {
		totalLen += len(s)
	}
	totalLen += len(args) - 1 // for the '\n' characters between args

	b := make([]byte, 0, totalLen)

	for i, s := range args {
		for j := 0; j < len(s); j++ {
			b = append(b, s[j])
		}
		if i < len(args)-1 {
			b = append(b, '\n')
		}
	}

	return string(b)
}
