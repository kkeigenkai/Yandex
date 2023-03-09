package task10

import (
	"reflect"
	"testing"
)

type test struct {
	str string
	exp map[rune]int
}

var tests = []test{
	{"hello", map[rune]int{'e': 8, 'h': 5, 'l': 17, 'o': 5}},
	{"abacaba", map[rune]int{'a': 44, 'b': 24, 'c': 16}},
	{"a", map[rune]int{'a': 1}},
}

func TestBoringLecture(t *testing.T) {
	for _, test := range tests {
		if got := BoringLecture(test.str); !reflect.DeepEqual(test.exp, got) {
			t.Errorf("( %s ) = %v,\n\twant: %v\n", test.str, got, test.exp)
		}
	}
}
