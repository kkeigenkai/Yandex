package task25

import "testing"

type test struct {
	cnt  int
	pins []int
	exp  int
}

var tests = []test{
	{6, []int{3, 13, 12, 4, 14, 6}, 5},
	{5, []int{4, 10, 0, 12, 2}, 6},
}

func TestPins(t *testing.T) {
	for _, test := range tests {
		if got := Pins(test.cnt, test.pins); got != test.exp {
			t.Errorf("( %d, %v ) = %d, want: %d\n", test.cnt, test.pins, got, test.exp)
		}
	}
}
