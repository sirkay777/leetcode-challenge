package linkedlist

import "testing"

func TestNewLinkedListInputValidation(t *testing.T) {
	data := []struct {
		input string
		valid bool
	}{
		{"[]", true},
		{"[1]", true},
		{"[1,2,3]", true},
		{"[1,   2,3] ", true},
		{"", false},
		{"[", false},
		{"]", false},
		{"[a]", false},
		{"[1,2,a]", false},
		{"[1,2,3]b", false},
	}
	for _, d := range data {
		_, err := NewLinkedList(d.input)
		got := err == nil
		if got != d.valid {
			t.Errorf("input: %s, expteced: %v, got: %v\n", d.input, d.valid, got)
		}
	}
}

func TestConvert(t *testing.T) {
	data := []string{
		"[]",
		"[1]",
		"[1,2]",
		"[1,2,3,4,5]",
		"[2,5,3,1]",
	}
	for _, v := range data {
		list, err := NewLinkedList(v)
		if err != nil {
			t.Fatal(err)
		}
		if v != LinkedListToString(list) {
			t.Errorf("convert failed, input:%s\n", v)
		}
	}
}
