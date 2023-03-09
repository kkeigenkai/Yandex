package task24

import "testing"

type test struct {
	cnt int
	que [][3]int
	exp int
}

var tests = []test{
	{5, [][3]int{{5, 10, 15}, {2, 10, 15}, {5, 5, 5}, {20, 20, 1}, {20, 1, 1}}, 12},
	{1, [][3]int{{1, 2, 3}}, 1},
}

func TestTicketPurchase(t *testing.T) {
	for _, test := range tests {
		if got := TicketPurchase(test.cnt, test.que); got != test.exp {
			t.Errorf("( %d, %v ) = %d,\n\twant: %d\n", test.cnt, test.que, got, test.exp)
		}
	}
}
