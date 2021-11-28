package linked_list

import (
	"testing"
)

func TestDeWeight(t *testing.T) {
	arr := []int{1, 7, 3, 1, 3, 7, 5, 7, 9, 11}
	head := ArrToLink(arr)
	PrintNode(head)
	// QueryWeight(head)
	DeWeight(head)
	PrintNode(head)
}
