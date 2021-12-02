package tree

// 深度优先

import (
	"fmt"
)

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
// 先序遍历 中序遍历都很好理解，后续遍历暂时无法独立实现
// 借助文档 https://studygolang.com/articles/13444 见PostOrderDoc
func PostOrder(node *BNode) {
	var stack []*BNode
	var preNode *BNode
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.LChild
		}
		// 出栈
		node, stack = pop(stack)

		noLR := node.LChild == nil && node.RChild == nil    // 无左右 可直接输出
		noR := node.RChild == nil && preNode == node.LChild // 无右 且左节点上次已经输出
		preIsR := preNode == node.RChild                    // 有右 且右节点上次已经输出 (无左和此相同)

		// 将上面的三个条件 直接放在if里可以触发短路 不过为了可读性，就放在上面
		if noLR || noR || preIsR {
			fmt.Println(node.Data)
			preNode = node // 1. 记录子节点 下一个循环 作为条件，判断左右节点是否都已经输出

			// 节点既然已经输出，就要置空，不然下次循环又要入栈
			node = nil
		} else {
			// 左节点已经输出，右节点还未输出，将指针移到右节点，下次循环压栈
			stack = append(stack, node) // 父节点暂时不输出
			node = node.RChild          // 指向右节点 进行压栈
		}
	}
}

// 见文档 标了注释才好理解
func PostOrderDoc(node *BNode) {
	var stack []*BNode
	var preNode *BNode
	var top *BNode                       // 使用top作为弹出的栈顶，避免污染循环体node
	for node != nil || len(stack) != 0 { // node主要指当前已经遍历的节点
		for node != nil {
			stack = append(stack, node)
			node = node.LChild
		}
		top, stack = pop(stack)
		if (top.LChild == nil && top.RChild == nil) || (top.RChild == nil && preNode == top.LChild) || preNode == top.RChild {
			fmt.Println(top.Data)
			preNode = top
			// 和上面相比 不用对node置空，因为栈顶是用top接收的
		} else {
			stack = append(stack, top) // 右节点未输出时，父节点暂时不出栈
			node = top.RChild          // 但是将node指向3的右节点
		}
	}
}
