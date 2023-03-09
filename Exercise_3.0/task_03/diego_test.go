package task03

import "testing"

type diegoTest struct {
	diegoCnt int
	diegoS   []int
	collCnt  int
	collS    []int
	exp      []int
}

var tests = []diegoTest{
	{1, []int{5}, 2, []int{4, 6}, []int{0, 1}},
	{3, []int{100, 1, 50}, 3, []int{300, 0, 75}, []int{3, 0, 2}},
	{0, []int{}, 0, []int{}, []int{0}},
}

func TestDiego(t *testing.T) {
	for _, test := range tests {
		got := Diego(test.diegoCnt, test.diegoS, test.collCnt, test.collS)
		if len(got) != len(test.exp) {
			t.Errorf("( %d, %v, %d, %v ) = %v,\n\twant: %v\n",
				test.diegoCnt, test.diegoS, test.collCnt, test.collS, got, test.exp)
			continue
		}

		for i, g := range got {
			if test.exp[i] != g {
				t.Errorf("( %d, %v, %d, %v ) = %v,\n\twant: %v\n",
					test.diegoCnt, test.diegoS, test.collCnt, test.collS, got, test.exp)
				break
			}
		}
	}
}
