package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	if len(a) < 2 {
		return true
	}

	firstCmp := f(a[0], a[1])

	for i := 1; i < len(a)-1; i++ {
		cmp := f(a[i], a[i+1])
		if firstCmp == 0 && cmp != 0 {
			firstCmp = cmp
		}

		if firstCmp != 0 && (cmp != 0) && (cmp*firstCmp < 0) {
			return false
		}
	}
	return true
}
