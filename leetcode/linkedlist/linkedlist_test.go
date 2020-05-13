package linkedlist

import "testing"

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
