package piscine

func Compact(ptr *[]string) int {
	s := *ptr
	var compacted []string

	for _, v := range s {
		if v != "" {
			compacted = append(compacted, v)
		}
	}

	*ptr = compacted

	count := 0
	for range compacted {
		count++
	}
	return count
}
