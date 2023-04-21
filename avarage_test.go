package gomocktest_test

import (
	gomocktest "go-mock-test"
	"testing"
)

func TestDivideInteger(t *testing.T) {
	tests := map[string]struct {
		avg       int
		b         []int
		wantError bool
		expect    int
	}{
		`correct integer avarage`: {
			avg:    3,
			b:      []int{1, 2, 4},
			expect: 2,
		},
		`incorrect integer avarage`: {
			avg:       3,
			b:         []int{1, 2, 4},
			expect:    1,
			wantError: true,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			out := gomocktest.Avarage(test.avg, test.b...)
			if (out != test.expect) != test.wantError {
				t.Errorf("got :%d, expected :%d", out, test.expect)
			}
		})
	}
}

func TestDivideFloat(t *testing.T) {
	tests := map[string]struct {
		avg       float64
		b         []float64
		wantError bool
		expect    float64
	}{
		`correct float avarage`: {
			avg:    1.5,
			b:      []float64{2.5, 0.5},
			expect: 2.0,
		},
		`incorrect float avarage`: {
			avg:       1.2,
			b:         []float64{2},
			expect:    2,
			wantError: true,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			out := gomocktest.Avarage(test.avg, test.b...)
			if (out != test.expect) != test.wantError {
				t.Errorf("got :%f, expected :%f", out, test.expect)
			}
		})
	}
}
