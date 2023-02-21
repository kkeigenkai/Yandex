package thedrunkgame

import "testing"

type theDrunkGameTest struct {
	fs, ss []int
	exp    string
}

var tests = []theDrunkGameTest{
	{[]int{1, 3, 5, 7, 9}, []int{2, 4, 6, 8, 0}, "second 5"},
	{[]int{2, 4, 6, 8, 0}, []int{1, 3, 5, 7, 9}, "first 5"},
	{[]int{1, 7, 3, 9, 4}, []int{5, 8, 0, 2, 6}, "second 23"},
	{[]int{9, 0, 9, 0, 9}, []int{0, 9, 0, 9, 0}, "botva"},
}

func TestTheDrunkGame(t *testing.T) {
	for _, test := range tests {
		if got := TheDrunkGame(test.fs, test.ss); got != test.exp {
			t.Errorf("( %v, %v ) = %q,\n\twant %q", test.fs, test.ss, got, test.exp)
		}
	}
}
