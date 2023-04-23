package gomocktest_test

import (
	gomocktest "go-mock-test"
	"go-mock-test/addition"
	"go-mock-test/divide"
	"go-mock-test/mocks"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestAvgInteger(t *testing.T) {
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

	mat := gomocktest.Mathematics{}
	mat.Addition = addition.Addition{}
	mat.Divide = divide.Divider{}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			out := mat.Avarage(test.avg, test.b...)
			if (out != test.expect) != test.wantError {
				t.Errorf("got :%d, expected :%d", out, test.expect)
			}
		})
	}
}

func TestAvgIntegerMock(t *testing.T) {
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
			b:         []int{3, 4, 6},
			expect:    3,
			wantError: true,
		},
	}

	mat := gomocktest.Mathematics{}
	ctrl := gomock.NewController(t)
	mockAdd := mocks.NewMockAdditional(ctrl)
	mockAdd.EXPECT().Add(gomock.Any(), gomock.Any()).Return(6).AnyTimes()

	mockDivide := mocks.NewMockDivide(ctrl)
	mockDivide.EXPECT().Divide(gomock.Any(), gomock.Any()).Return(2).AnyTimes()

	mat.Addition = mockAdd
	mat.Divide = mockDivide
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			out := mat.Avarage(test.avg, test.b...)
			if (out != test.expect) != test.wantError {
				t.Errorf("got :%d, expected :%d", out, test.expect)
			}
		})
	}
}
