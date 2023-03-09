package task21

import "testing"

type threeOnesTest struct {
	n, exp int
}

var tests = []threeOnesTest{
	{1, 2},
	{2, 4},
	{3, 7},
	{4, 13},
}

func TestThreeOnes(t *testing.T) {
	for _, test := range tests {
		if got := ThreeOnes(test.n); got != test.exp {
			t.Errorf("( %d ) = %d, want: %d\n", test.n, got, test.exp)
		}
	}
}
