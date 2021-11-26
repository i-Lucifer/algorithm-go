package linked_list

import "testing"

func TestRSort(t *testing.T) {
	head := NewNode(8)
	// PrintNode(head)
	RSortInsert(head)
	PrintNode(head)
}
