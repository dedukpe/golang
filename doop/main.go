package main

import "os"

func main() {
	args := os.Args[1:]
	if len(args) != 3 {
		return
	}

	a, ok := parseInt(args[0])
	b, ok2 := parseInt(args[2])
	if !ok || !ok2 {
		return
	}

	op := args[1]

	switch op {
	case "+":
		res, ok := safeAdd(a, b)
		if !ok {
			return
		}
		printInt(res)
	case "-":
		res, ok := safeSub(a, b)
		if !ok {
			return
		}
		printInt(res)
	case "*":
		res, ok := safeMul(a, b)
		if !ok {
			return
		}
		printInt(res)
	case "/":
		if b == 0 {
			printString("No division by 0\n")
			return
		}
		printInt(a / b)
	case "%":
		if b == 0 {
			printString("No modulo by 0\n")
			return
		}
		printInt(a % b)
	default:
		return
	}
}

func parseInt(s string) (int64, bool) {
	var n int64 = 0
	neg := false
	i := 0
	if len(s) == 0 {
		return 0, false
	}
	if s[0] == '-' {
		neg = true
		i++
	} else if s[0] == '+' {
		i++
	}
	if i == len(s) {
		return 0, false
	}
	for ; i < len(s); i++ {
		c := s[i]
		if c < '0' || c > '9' {
			return 0, false
		}
		digit := int64(c - '0')

		if n > (9223372036854775807-digit)/10 {
			return 0, false
		}
		n = n*10 + digit
	}
	if neg {
		n = -n
		if n > 0 {
			return 0, false
		}
	}
	return n, true
}

func printInt(n int64) {
	if n == 0 {
		os.Stdout.Write([]byte{'0', '\n'})
		return
	}

	var buf [20]byte
	i := len(buf)

	neg := false
	if n < 0 {
		neg = true
		n = -n
	}

	for n > 0 {
		i--
		buf[i] = byte('0' + n%10)
		n /= 10
	}
	if neg {
		i--
		buf[i] = '-'
	}

	s := buf[i:]

	s = append(s, '\n')

	os.Stdout.Write(s)
}

func printString(s string) {
	os.Stdout.Write([]byte(s))
}

func safeAdd(a, b int64) (int64, bool) {
	res := a + b
	if (b > 0 && res < a) || (b < 0 && res > a) {
		return 0, false
	}
	return res, true
}

func safeSub(a, b int64) (int64, bool) {
	res := a - b
	if (b < 0 && res < a) || (b > 0 && res > a) {
		return 0, false
	}
	return res, true
}

func safeMul(a, b int64) (int64, bool) {
	if a == 0 || b == 0 {
		return 0, true
	}

	res := a * b
	if res/b != a {
		return 0, false
	}

	if (a == -1 && b == -9223372036854775808) || (b == -1 && a == -9223372036854775808) {
		return 0, false
	}

	return res, true
}
