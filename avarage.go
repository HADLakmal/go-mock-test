package gomocktest

import (
	"go-mock-test/addition"
	"go-mock-test/divide"
)

func Avarage[T int | float64](avg T, value ...T) T {
	var sum T
	for _, t := range value {
		sum = addition.Addition(sum, t)
	}
	return divide.Divide(sum, avg)
}
