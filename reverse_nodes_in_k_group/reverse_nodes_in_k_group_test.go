package reverse_nodes_in_k_group

import (
	. "github.com/sirkay777/leetcode-challenge/leetcode/linkedlist"
	"testing"
)

func TestReverseNodesInKGroup(t *testing.T) {
	data := []struct {
		input    string
		k        int
		expected string
	}{
		{"[1,2,3,4,5]", 2, "[2,1,4,3,5]"},
		{"[1,2,3,4,5]", 3, "[3,2,1,4,5]"},
		{"[1,2,3,4,5]", 1, "[1,2,3,4,5]"},
		{"[1,2,3,4,5]", 4, "[4,3,2,1,5]"},
		{"[1,2,3,4,5]", 5, "[5,4,3,2,1]"},
	}
	for _, d := range data {
		list, err := NewLinkedList(d.input)
		if err != nil {
			t.Fatal(err)
		}
		newList := reverseKGroup(list, d.k)
		got := LinkedListToString(newList)
		if got != d.expected {
			t.Errorf("expected: %s, got: %s\n", d.expected, got)
		}
	}

}
