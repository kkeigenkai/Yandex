package task12

import "testing"

type correctParenthesisTest struct {
	str string
	exp bool
}

var tests = []correctParenthesisTest{
	{"()[]{}", true},
	{"([)", false},
	{"", true},
	{"(", false},
}

func TestCorrectParenthesis(t *testing.T) {
	for _, test := range tests {
		if got := CorrectParenthesis(test.str); got != test.exp {
			t.Errorf("( %q ) = %t,\n\twant: %t", test.str, got, test.exp)
		}
	}
}
