package postfixnotation

import "testing"

type postfixNotationTest struct {
	str string
	exp int
}

var tests = []postfixNotationTest{
	{"8 9 + 1 7 - *", -102},
}

func TestPostfixNotation(t *testing.T) {
	for _, test := range tests {
		if got := PostfixNotation(test.str); got != test.exp {
			t.Errorf("( %q ) = %d,\n\twant: %d", test.str, got, test.exp)
		}
	}
}
