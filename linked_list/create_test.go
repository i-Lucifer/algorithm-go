package linked_list

import (
	"testing"
)

// 递归打印
func TestCreate(t *testing.T) {
	arr := []int{1, 3, 1, 3, 7, 5, 7, 9, 11}
	head := ArrToLink(arr)
	PrintNode(head)
}
