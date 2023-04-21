package domian

// go:generate mockgen -source=addition_interface
type Divide[T int | float64] interface {
	Divide(a, b T) T
}
