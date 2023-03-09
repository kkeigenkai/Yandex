package task14

import "testing"

type sortLiteTest struct {
	trainCnt int
	trains   []int
	exp      bool
}

var tests = []sortLiteTest{
	{3, []int{3, 2, 1}, true},
	{4, []int{4, 1, 3, 2}, true},
	{3, []int{2, 3, 1}, false},
	{1, []int{1}, true},
	{2, []int{1, 2}, true},
	{2, []int{2, 1}, true},
}

func TestSortLite(t *testing.T) {
	for _, ti := range tests {
		if got := SortLite(ti.trainCnt, ti.trains); got != ti.exp {
			t.Errorf("( %d, %v ) = %t,\n\twant: %t", ti.trainCnt, ti.trains, got, ti.exp)
		}
	}
}
