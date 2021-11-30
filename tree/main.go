package main

import "fmt"

// BinaryTreeNode
type BNode struct {
	Data       interface{}
	LeftChild  *BNode
	RightChild *BNode
}

func NewNode() *BNode {
	return &BNode{}
}

// ⭐️ 注意，树结构，可以用数组实现，也可以用链表实现，平常编写代码的时候，用链表是比较直观，并且方便扩容。
// ⭐️ 如果用二分法把数据写入数组，这不是也是树吗
//
// 把有序数组放到二叉树
// 这里把数组二分，但不是二分查找，而是遍历了一遍数组，时间复杂度O(n)
func Array2Tree(arr []int) *BNode {
	return ArrayToTree(arr, 0, len(arr)-1)
}
func ArrayToTree(arr []int, start int, end int) *BNode {
	var root *BNode
	if end < start {
		return root
	}
	root = NewNode()
	// fmt.Println(start, end)
	mid := (start + end + 1) / 2
	// fmt.Println(mid)
	// fmt.Println(arr)
	root.Data = arr[mid]
	root.LeftChild = ArrayToTree(arr, start, mid-1)
	root.RightChild = ArrayToTree(arr, mid+1, end)
	return root
}

// 中序遍历
func MidOrder(root *BNode) {

}

// 层次遍历 (可以中序打印，也可以中序塞到数组里)
func LayerOrder(root *BNode) {
	if root == nil {
		return
	}
}

func main() {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(data)
	root := Array2Tree(data)
	fmt.Println(root)
	// fmt.Println(root.LeftChild)
	// fmt.Println(root.RightChild)
	MidOrder(root)
	LayerOrder(root)
}
