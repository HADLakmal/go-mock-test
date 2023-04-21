package divide_test

import (
	"go-mock-test/divide"
	"testing"
)

func TestDivideInteger(t *testing.T) {
	tests := map[string]struct {
		a         int
		b         int
		wantError bool
		expect    int
	}{
		`correct integer divide`: {
			a:      1,
			b:      2,
			expect: 0,
		},
		`incorrect integer divide`: {
			a:         4,
			b:         2,
			expect:    1,
			wantError: true,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			out := divide.Divide(test.a, test.b)
			if (out != test.expect) != test.wantError {
				t.Errorf("got :%d, expected :%d", out, test.expect)
			}
		})
	}
}

func TestDivideFloat(t *testing.T) {
	tests := map[string]struct {
		a         float64
		b         float64
		wantError bool
		expect    float64
	}{
		`correct float divide`: {
			a:      2.5,
			b:      2.5,
			expect: 1.0,
		},
		`incorrect float divide`: {
			a:         1.2,
			b:         2,
			expect:    2,
			wantError: true,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			out := divide.Divide(test.a, test.b)
			if (out != test.expect) != test.wantError {
				t.Errorf("got :%f, expected :%f", out, test.expect)
			}
		})
	}
}
