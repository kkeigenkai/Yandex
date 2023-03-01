package task22

import "testing"

type grasshopperTest struct {
	size, jump, exp int
}

var tests = []grasshopperTest{
	{1, 2, 1},
	{2, 2, 1},
	{8, 2, 21},
	{8, 4, 56},
}

func TestGrasshopper(t *testing.T) {
	for _, test := range tests {
		if got := Grasshopper(test.size, test.jump); got != test.exp {
			t.Errorf("( %d, %d ) = %d, want: %d", test.size, test.jump, got, test.exp)
		}
	}
}
