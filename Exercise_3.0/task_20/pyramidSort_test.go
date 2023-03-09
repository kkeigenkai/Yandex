package task20

import "testing"

type pyramidSortTest struct {
	cnt     int
	si, exp []int
}

var tests = []pyramidSortTest{
	{1, []int{1}, []int{1}},
	{2, []int{3, 1}, []int{1, 3}},
	{5, []int{3, 8, 1, 2, 9}, []int{1, 2, 3, 8, 9}},
	{
		5,
		[]int{970230416, 819698993, 12233245, 983424780, 762955113, 552371399, 728993264, 46396314, 339526042},
		[]int{12233245, 46396314, 339526042, 552371399, 728993264, 762955113, 819698993, 970230416, 983424780},
	},
}

func TestPyramidSort(t *testing.T) {
	for _, test := range tests {
		got := PyramidSort(test.cnt, test.si)
		if len(got) != len(test.exp) {
			t.Errorf("( %d, %v ) = %v,\n\twant: %v\n", test.cnt, test.si, got, test.exp)
			continue
		}
		for i := range got {
			if got[i] != test.exp[i] {
				t.Errorf("( %d, %v ) = %v,\n\twant: %v\n", test.cnt, test.si, got, test.exp)
				break
			}
		}
	}
}
