package task28

import "testing"

type test struct {
	n, m int
	exp  int
}

var tests = []test{
	{3, 2, 1},
	{31, 34, 293930},
}

func TestKnightMove(t *testing.T) {
	for _, test := range tests {
		if got := KnightMove(test.n, test.m); got != test.exp {
			t.Errorf("( %d, %d ) = %d, want: %d", test.n, test.m, got, test.exp)
		}
	}
}
