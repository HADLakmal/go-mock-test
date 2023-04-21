package addition

func Addition[T int | float64](a, b T) T {
	return a + b
}
