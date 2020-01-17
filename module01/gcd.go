package module01

// GCDEuclidean gets the greatest common denominator using the Euclidean algorithm recursively
func GCDEuclidean(a, b int) int {
	if b == 0 {
		return a
	}
	a, b = b, a%b
	return GCDEuclidean(a, b)
}

// GCD gets the greatest common denominator without using recursion
func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
