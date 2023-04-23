package addition_test

import (
	"go-mock-test/addition"
	"testing"
)

func TestAdditionInteger(t *testing.T) {
	tests := map[string]struct {
		a         int
		b         int
		wantError bool
		expect    int
	}{
		`correct integer addition`: {
			a:      1,
			b:      2,
			expect: 3,
		},
		`incorrect integer addition`: {
			a:         1,
			b:         2,
			expect:    4,
			wantError: true,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ad := addition.Addition{}
			out := ad.Add(test.a, test.b)
			if (out != test.expect) != test.wantError {
				t.Errorf("got :%d, expected :%d", out, test.expect)
			}
		})
	}
}
