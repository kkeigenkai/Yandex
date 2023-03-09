package sntp

import "testing"

type SNTPTest struct {
	ta, tb, tc string
	exp        string
}

var tests = []SNTPTest{
	{"15:01:00", "18:09:45", "15:01:40", "18:10:05"},
}

func TestSNTP(t *testing.T) {
	for _, test := range tests {
		if got := SNTP(test.ta, test.tb, test.tc); got != test.exp {
			t.Errorf("( %q, %q, %q ) = %q,\n\twant: %q\n", test.ta, test.tb, test.tc, got, test.exp)
		}
	}
}
