package prettystring

import "testing"

type prettyStringTest struct {
	arg1      string
	arg2, exp int
}

var prettyStringTests = []prettyStringTest{
	{"abcaz", 2, 4},
	{"helto", 2, 3},
	{"a", 5, 1},
	{"aaaaa", 1, 5},
	{"ddaaadd", 3, 7},
	{"", 5, 0},
	{"afdsfhsdfbsdfhzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzz", 5, 40},
}

func TestPrettyString(t *testing.T) {
	for _, test := range prettyStringTests {
		if got := PrettyString(test.arg1, test.arg2); got != test.exp {
			t.Errorf("PrettyString(%q, %d) expected %d got %d", test.arg1, test.arg2, test.exp, got)
		}
	}
}
