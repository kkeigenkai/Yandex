package oslite

import "testing"

type osLiteTest struct {
	sectCount, partCount int
	sects                [][2]int
	exp                  int
}

var osLiteTests = []osLiteTest{
	{10, 3, [][2]int{{1, 3}, {4, 7}, {3, 4}}, 1},
	{10, 4, [][2]int{{1, 3}, {4, 5}, {7, 8}, {4, 6}}, 3},
	{10, 3, [][2]int{{1, 4}, {2, 6}, {7, 10}}, 2},
	{200, 4, [][2]int{{100, 150}, {110, 175}, {176, 180}, {101, 174}}, 2},
	{1, 5, [][2]int{{1, 1}, {1, 1}, {1, 1}, {1, 1}, {1, 1}}, 1},
}

func TestOsLite(t *testing.T) {
	for _, test := range osLiteTests {
		if got := Oslite(test.sectCount, test.partCount, test.sects); test.exp != got {
			t.Errorf("Oslite(%d, %d, %+v) = %d, want %d", test.sectCount, test.partCount, test.sects, got, test.exp)
		}
	}
}
