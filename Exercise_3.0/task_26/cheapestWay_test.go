package task26

import "testing"

type test struct {
	n, m int
	ssi  [][]int
	exp  int
}

var tests = []test{
	{5, 5, [][]int{
		{1, 1, 1, 1, 1},
		{3, 100, 100, 100, 100},
		{1, 1, 1, 1, 1},
		{2, 2, 2, 2, 1},
		{1, 1, 1, 1, 1},
	}, 11},
	{1, 1, [][]int{
		{1},
	}, 1},
}

func TestCheapestWay(t *testing.T) {
	for _, test := range tests {
		if got := CheapestWay(test.n, test.m, test.ssi); got != test.exp {
			t.Errorf("( %d, %d, %v ) = %d, want: %d\n", test.n, test.m, test.ssi, got, test.exp)
		}
	}
}
