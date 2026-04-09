package piscine

func Enigma(a ***int, b *int, c *******int, d ****int) {
	// store the original values
	tempA := ***a
	tempB := *b
	tempC := *******c
	tempD := ****d

	// reassign according to the instructions
	*******c = tempA // a -> c
	****d = tempC    // c -> d
	*b = tempD       // d -> b
	***a = tempB     // b -> a
}
