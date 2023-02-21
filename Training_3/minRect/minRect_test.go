package minrect

import "testing"

type minRectTest struct {
	sqCnt  int
	coords [][2]int
	exp    [4]int
}

var tests = []minRectTest{
	{3, [][2]int{{1, 1}, {1, 10}, {5, 5}}, [4]int{1, 1, 5, 10}},
	{5, [][2]int{{2, 2}, {3, 5}, {20, 6}, {11, 7}, {8, 1}}, [4]int{2, 1, 20, 7}},
	{5, [][2]int{{2, 2}, {3, 5}, {20, 6}, {11, 7}, {8, 3}}, [4]int{2, 2, 20, 7}},
	{5, [][2]int{{668, 1676}, {1810, 6372}, {3301, 4297}, {2297, 6251}, {2598, 3714}}, [4]int{0, 0, 0, 0}},
}

// 668 1676
// 1810 6372
// 3301 4297
// 2297 6251
// 2598 3714

func TestMinRect(t *testing.T) {
	for _, test := range tests {
		if got := MinRect(test.sqCnt, test.coords); test.exp != got {
			t.Errorf("MinRect(%d, %v) = %v,\n\twant: %v", test.sqCnt, test.coords, got, test.exp)
		}
	}
}
