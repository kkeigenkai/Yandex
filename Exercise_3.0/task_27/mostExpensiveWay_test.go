package task27

import "testing"

type test struct {
	n, m int
	ssi  [][]int
	expI int
	expS string
}

var tests = []test{
	{5, 5, [][]int{
		{9, 9, 9, 9, 9},
		{3, 0, 0, 0, 0},
		{9, 9, 9, 9, 9},
		{6, 6, 6, 6, 8},
		{9, 9, 9, 9, 9},
	}, 74, "D D R R R R D D"},
}

func TestMostExpensiveWay(t *testing.T) {
	for _, test := range tests {
		if gotI, gotS := MostExpensiveWay(test.n, test.m, test.ssi); gotI != test.expI || gotS != test.expS {
			t.Errorf("( %d, %d, %v )\n\t= %d, %s, want: %d, %s\n", test.n, test.m, test.ssi, gotI, gotS, test.expI, test.expS)
		}
	}
}
