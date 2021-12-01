package tree

import (
	"fmt"
	"testing"
)

func TestLayer(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("array: ", data) // 二分法转为树后，6 3,9 2,5,8,10 1,4,7
	root := Array2Tree(data)
	// LayerOrder(root)
	LayerOrder2(root)
}
