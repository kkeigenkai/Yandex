package task05

import "testing"

type test struct {
	cnt int
	si  []int
	exp int
}

var tests = []test{
	{3, []int{1, 1, 1}, 2},
	{2, []int{3, 4}, 3},
	{1, []int{10}, 0},
	{4, []int{5, 4, 3, 2}, 9},
	{5, []int{1, 2, 3, 4, 5}, 10},
}

func TestGoodString(t *testing.T) {
	for _, test := range tests {
		if got := GoodString(test.cnt, test.si); got != test.exp {
			t.Errorf(" ( %d, %v ) = %d, want: %d\n", test.cnt, test.si, got, test.exp)
		}
	}
}
