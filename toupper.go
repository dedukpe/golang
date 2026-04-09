package piscine

func ToUpper(s string) string {
	b := []byte(s)
	for i := 0; i < len(b); i++ {
		if b[i] >= 'a' && b[i] <= 'z' {
			b[i] = b[i] - ('a' - 'A')
		}
	}
	return string(b)
}
