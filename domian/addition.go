package domian

// go:generate mockgen -source=addition_interface
type Additional[T int | float64] interface {
	Addition(a, b T) T
}
