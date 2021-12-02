package tree

// 深度优先
// 深度优先遍历的本质是，按照一定的顺序访问各个节点，先，中，后的区别在于，是访问时输出，还是返回时输出

import "fmt"

// 先序遍历 根 左 右
func PreOrderRec(root *BNode) {
	if root == nil {
		return
	}

	fmt.Println(root.Data) // 先序的本质是入栈时输出

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
	// 入栈时不输出
	// 出栈时
	// 1. 叶子节点 无左 无右 直接出
	// 2. 非叶子节点 无左 压入的就是右栈 出栈时也直接出右，跳4
	// 3. 非叶子节点 如果是左节点(if.L)出栈 -> 执行右节点(if.R) -> 无右时 -> 当前节点输出
	// 4. 非叶子节点 如果是右节点(if.R)出栈 -> 当前节点输出
	fmt.Println(root.Data)
}
