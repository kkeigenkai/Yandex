package task19

import "testing"

type test struct {
	ss  [][2]int
	exp []int
}

var tests = []test{
	{[][2]int{{0, 10000}, {1}}, []int{10000}},
	{
		[][2]int{{0, 1}, {0, 345}, {1}, {0, 4346}, {1}, {0, 2435}, {1}, {0, 235}, {0, 5}, {0, 365}, {1}, {1}, {1}, {1}},
		[]int{345, 4346, 2435, 365, 235, 5, 1},
	},
}

func TestHeap(t *testing.T) {
	for _, test := range tests {
		got := Heap(test.ss)
		if len(got) != len(test.exp) {
			t.Errorf("( %v )\n\t= %v,\n\twant: %v\n", test.ss, got, test.exp)
			continue
		}
		for i := range got {
			if got[i] != test.exp[i] {
				t.Errorf("( %v )\n\t= %v,\n\twant: %v\n", test.ss, got, test.exp)
				break
			}
		}
	}
}
