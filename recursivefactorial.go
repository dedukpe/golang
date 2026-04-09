package piscine

func RecursiveFactorial(nb int) int {
	if nb < 0 || nb > 20 {
		return 0
	}
	if nb == 0 || nb == 1 {
		return 1
	}
	// Calculate factorial recursively
	prev := RecursiveFactorial(nb - 1)
	if prev == 0 {
		return 0
	}
	// Check for overflow: if nb * prev would overflow, return 0
	if nb > (1<<63-1)/prev {
		return 0
	}
	return nb * prev
}
