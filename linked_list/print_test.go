package linked_list

import (
	"fmt"
	"testing"
)

// 递归打印
func TestPrint(t *testing.T) {
	head := NewNode(8)
	PrintByFunc(head)
	fmt.Println("")
	PrintNode(head)
}
