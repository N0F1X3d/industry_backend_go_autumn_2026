package main

func rotateRunes(s string, shift int) string {
	runes := []rune(s)
	n := len(runes)
	if n == 0 {
		return s
	}

	effective := shift % n
	if effective < 0 {
		effective += n
	}

	rotated := make([]rune, n)
	copy(rotated, runes[effective:])
	copy(rotated[n-effective:], runes[:effective])
	return string(rotated)
}
