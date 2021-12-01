package tree

import "fmt"

// 深度优先
// 先序遍历 根 左 右
func PreOrderRec(root *BNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Data)
	if root.LChild != nil {
		PreOrderRec(root.LChild)
	}
	if root.RChild != nil {
		PreOrderRec(root.RChild)
	}
}

// 中序遍历 左 根 右
func InOrderRec(root *BNode) {
	if root == nil {
		return
	}
	if root.LChild != nil {
		InOrderRec(root.LChild)
	}
	fmt.Println(root.Data)
	if root.RChild != nil {
		InOrderRec(root.RChild)
	}
}

// 后序遍历 左 右 根
func PostOrderRec(root *BNode) {
	if root == nil {
		return
	}
	if root.LChild != nil {
		PostOrderRec(root.LChild)
	}
	if root.RChild != nil {
		PostOrderRec(root.RChild)
	}
	fmt.Println(root.Data)
}
