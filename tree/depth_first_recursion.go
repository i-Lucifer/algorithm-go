package tree

import "fmt"

// 深度优先
// 先序遍历 根 左 右
func PreOrderRec(root *BNode) {
	if root == nil {
		return
	}

	fmt.Println(root.Data) // 输出

	if root.LChild != nil { // 入栈
		PreOrderRec(root.LChild)
	} // 出栈

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
	} // 出栈

	fmt.Println(root.Data) // 中序的本质是出栈后输出

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
	fmt.Println(root.Data) // 后序的本质也是调整当前节点的输出顺序，从入栈和出栈的角度来讲，是交叉的
}
