package tree

import "fmt"

// 深度优先

// 先序遍历 根 左 右
// 根节点优先输出
// 二叉树的深度优先遍历一般需要借助栈结构 (递归也是栈) 输出完下层节点，需要回朔
func PreOrder(node *BNode) {
	// 1. 栈 优化说明 for循环直接对node判空，不需要额外判定
	var stack []*BNode
	for node != nil || len(stack) > 0 {
		// 输出并入栈 优化说明，将父节点和子节点的入栈合并，可见git记录
		for node != nil {
			fmt.Println("先序遍历的核心是入栈前输出: ", node.Data)
			stack = append(stack, node)
			node = node.LChild // 递归左节点
		}

		// 出栈
		// 优化1 封装pop函数，看起来更简介
		// 优化2 取消node.RChild判空，如果是空，直接赋值，得到的也是node=nil
		// 优化3 进入循环1时
		// 		    node和stack必定有一个不为空
		//          如果stack为空，进入循环2是，必定入栈，所以stack不为空
		// 			如果node为空，stack不为空，所以此处不需要做stack空判断
		// if len(stack) > 0 {
		node, stack = pop(stack)
		node = node.RChild
		// }
	}
}

// 中序遍历 左 根 右
func InOrder(node *BNode) {
	var stack []*BNode
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.LChild
		} // 结束时 node必定为空

		// 出栈
		// 中序遍历的核心是出栈时输出
		// if len(stack) > 0 {
		node, stack = pop(stack) // 弹出时，node必定不为空
		fmt.Println(node.Data)
		node = node.RChild
		// }
	}
}

// 后序遍历 左 右 根
// 后序遍历 需要用双栈结构
func PostOrder(node *BNode) {
	var stack []*BNode
	// var preNode *BNode
	for node != nil || len(stack) > 0 {
		// 入栈时无须理会
		for node != nil {
			stack = append(stack, node)
			node = node.LChild
		}

		node, stack = pop(stack)
		node = node.RChild
	}
}
