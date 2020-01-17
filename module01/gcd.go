package module01

// GCD gets the greatest common denominator using the Euclidean algorithm
func GCD(a, b int) int {
	if b == 0 {
		return a
	}
	a, b = b, a%b
	return GCD(a, b)
}
