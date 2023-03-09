package controlwork

import "testing"

type controlWorkTest struct {
	arg1, agr2, arg3, arg4 int
	exp1, exp2             int
}

var controlWorkTests = []controlWorkTest{
	{25, 2, 1, 2, 2, 2},
	{25, 13, 7, 1, -1, 0},
	{10, 3, 2, 2, 1, 1},
	{9, 2, 2, 2, 3, 2},
	{16, 8, 6, 2, 2, 2},
	{16, 3, 7, 1, 8, 2},
	{6, 6, 1, 1, -1, 0},
	{11, 5, 4, 1, 1, 2},
}

func TestControlWork(t *testing.T) {
	for _, test := range controlWorkTests {
		if row, side := ControlWork(test.arg1, test.agr2, test.arg3, test.arg4); row != test.exp1 || side != test.exp2 {
			t.Errorf("ControlWork(%d, %d, %d, %d) = (%d, %d), want (%d, %d)", test.arg1, test.agr2, test.arg3, test.arg4, row, side, test.exp1, test.exp2)
		}
	}
}
