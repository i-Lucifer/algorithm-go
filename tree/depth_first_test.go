package tree

import (
	"fmt"
	"testing"
)

func TestOrder(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("array: ", data) // 二分法转为树后，6 3,9 2,5,8,10 1,4,7
	root := Array2Tree(data)
	// fmt.Println("-----pre-----")
	// PreOrderRec(root)
	// PreOrder(root)
	// fmt.Println("-----in-----")
	// InOrderRec(root)
	// InOrder(root)
	// fmt.Println("-----post-----")
	// PostOrderRec(root)
	PostOrder(root)
}
