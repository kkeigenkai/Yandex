package task09

import "testing"

type test struct {
	base [3]int
	ssi  [][]int
	sr   [][4]int
	exp  []int
}

var tests = []test{
	{[3]int{3, 3, 2}, [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}, [][4]int{
		{2, 2, 3, 3},
		{1, 1, 2, 3},
	}, []int{28, 21}},
	{[3]int{4, 4, 2}, [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}, [][4]int{{1, 1, 2, 2}, {2, 2, 3, 3}}, []int{14, 34}},
}

func TestRectSum(t *testing.T) {
	for _, test := range tests {
		got := RectSum(test.base, test.ssi, test.sr)
		if len(got) != len(test.exp) {
			t.Errorf("( %v, %v, %v )\n\t= %v,\n\twant: %v\n",
				test.base, test.ssi, test.sr, got, test.exp)
			continue
		}
		for i := range got {
			if got[i] != test.exp[i] {
				t.Errorf("( %v, %v, %v )\n\t= %v,\n\twant: %v\n",
					test.base, test.ssi, test.sr, got, test.exp)
				break
			}
		}
	}
}
