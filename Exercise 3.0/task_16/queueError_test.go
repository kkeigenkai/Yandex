package queueerror

import (
	"testing"
)

type queueErrorTest struct {
	cmds []string
	exp  []string
}

var tests = []queueErrorTest{
	{[]string{"push 1", "front", "exit"}, []string{"ok", "1", "bye"}},
	{
		[]string{"size", "push 1", "size", "push 2", "size", "push 3", "size", "exit"},
		[]string{"0", "ok", "1", "ok", "2", "ok", "3", "bye"},
	},
	{[]string{"size", "exit", "push 2", "size"}, []string{"0", "bye"}},
	{[]string{"pop", "exit"}, []string{"error", "bye"}},
}

func TestQueueError(t *testing.T) {
	for _, test := range tests {
		got := QueueError(test.cmds)
		if len(got) != len(test.exp) {
			t.Errorf(" ( %v ) = %v,\n\twant: %v", test.cmds, got, test.exp)
			continue
		}
		for i, cmd := range test.exp {
			if cmd != got[i] {
				t.Errorf(" ( %v ) = %v,\n\twant: %v", test.cmds, got, test.exp)
			}
		}
	}
}
