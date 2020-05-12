package reverse_linked_list

import (
	. "github.com/sirkay777/leetcode-challenge/leetcode/linkedlist"
	"testing"
)

func BenchmarkReverseList(b *testing.B) {
	head, err := NewLinkedList("[1,2,3,4,5]")
	if err != nil {
		b.Fatal(err)
	}
	for i := 0; i < b.N; i++ {
		ReverseList(head)
	}
}

func TestReverseLinkedList(t *testing.T) {
	data := []struct {
		input    string
		expected string
	}{
		{"[1,2,3,4,5]", "[5,4,3,2,1]"},
		{"[1]", "[1]"},
		{"[]", "[]"},
	}

	for _, d := range data {
		head, err := NewLinkedList(d.input)
		if err != nil {
			t.Fatal(err)
		}
		reversedHead := ReverseList(head)
		got := LinkedListToString(reversedHead)
		if got != d.expected {
			t.Errorf("got: %s, expected: %s", got, d.expected)
		}
	}
}
