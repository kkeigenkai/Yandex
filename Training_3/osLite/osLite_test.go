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
}

func TestOsLite(t *testing.T) {

}
