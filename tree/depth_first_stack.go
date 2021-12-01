package tree

import "fmt"

// 深度优先

// 先序遍历 根 左 右
// 根节点优先输出
// 二叉树的深度优先遍历一般需要借助栈结构 (递归也是栈) 输出完下层节点，需要回朔
func PreOrder(root *BNode) {
	if root == nil {
		return
	}
	var stack []*BNode
	node := root

	for node != nil || len(stack) > 0 {
		// 有节点时入栈 先序遍历的核心是入栈时输出
		if node != nil {
			fmt.Println("父节点入栈并输出: ", node.Data)
			stack = append(stack, node)
			for node.LChild != nil {
				node = node.LChild
				stack = append(stack, node)
				fmt.Println("子节点入栈并输出: ", node.Data)
			}
		}

		// 开始出栈
		newLen := len(stack)
		if newLen > 0 {
			node = stack[newLen-1]
			stack = stack[0 : newLen-1]
			if node.RChild != nil {
				node = node.RChild
			} else {
				node = nil
			}
		}
	}
}

// 中序遍历 左 根 右
func InOrder(root *BNode) {
	if root == nil {
		return
	}
	var stack []*BNode
	node := root

	for node != nil || len(stack) > 0 {
		// 入栈
		if node != nil {
			stack = append(stack, node)
			for node.LChild != nil {
				node = node.LChild
				stack = append(stack, node)
			}
		}

		// 出栈
		newLen := len(stack)
		if newLen > 0 {
			node = stack[newLen-1]
			stack = stack[0 : newLen-1]
			fmt.Println(node.Data) // 中序遍历的核心是，出栈后输出
			if node.RChild != nil {
				node = node.RChild
			} else {
				node = nil
			}
		}
	}
}

// 后序遍历 左 右 根
// 后序遍历 需要用双栈结构
func PostOrder(root *BNode) {
	// if root == nil {
	// 	return
	// }
	// var stack []*BNode
	// var stackVal []interface{}
	// node := root

	// for node != nil || len(stack) > 0 {
	// 	// 入栈
	// 	if node != nil {
	// 		stack = append(stack, node)
	// 		for node.LChild != nil {
	// 			node = node.LChild
	// 			stack = append(stack, node)
	// 		}
	// 	}

	// 	// 出栈
	// 	newLen := len(stack)
	// 	if newLen > 0 {
	// 		node = stack[newLen-1]
	// 		stack = stack[0 : newLen-1]
	// 		node = pop(stack)

	// 		if node.LChild != nil {
	// 			fmt.Println(node.Data)
	// 		}

	// 		if node.RChild != nil {
	// 			node = node.RChild
	// 			fmt.Println(node.Data)
	// 		} else {
	// 			node = nil
	// 		}
	// 	}
	// }
}
