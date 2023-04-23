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
			dv := divide.Divider{}
			out := dv.Divide(test.a, test.b)
			if (out != test.expect) != test.wantError {
				t.Errorf("got :%d, expected :%d", out, test.expect)
			}
		})
	}
}
