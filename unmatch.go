package piscine

func Unmatch(a []int) int {
	counts := make(map[int]int)

	for _, n := range a {
		counts[n]++
	}

	for _, n := range a {
		if counts[n]%2 != 0 {
			return n
		}
	}

	return -1
}
