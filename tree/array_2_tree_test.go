package tree

import (
	"fmt"
	"testing"
)

func TestToTree(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("array: ", data)
	root := Array2Tree(data) // 二分法转为树后，6 3,9 2,5,8,10 1,4,7
	fmt.Println("tree: ", root)
	fmt.Println("-----")
	fmt.Println(root.LChild)
	fmt.Println(root.RChild)
	fmt.Println("-----")
	fmt.Println(root.LChild.LChild)
	fmt.Println(root.LChild.RChild)
	fmt.Println(root.RChild.LChild)
	fmt.Println(root.RChild.RChild)
	fmt.Println("-----")
	fmt.Println(root.LChild.LChild.LChild)
	fmt.Println(root.LChild.RChild.LChild)
	fmt.Println(root.RChild.LChild.LChild)
}
