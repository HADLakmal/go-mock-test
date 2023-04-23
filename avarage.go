package gomocktest

import (
	domian "go-mock-test/domain"
)

type Mathematics struct {
	Addition domian.Additional
	Divide   domian.Divide
}

func (m Mathematics) Avarage(avg int, value ...int) int {
	var sum int
	for _, t := range value {
		sum = m.Addition.Add(sum, t)
	}
	return m.Divide.Divide(sum, avg)
}
