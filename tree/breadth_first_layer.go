package tree

import (
	"fmt"
)

// 广度优先
// 层次遍历
// 二叉树的层次遍历一般需要借助队列
func LayerOrder(root *BNode) {
	if root == nil {
		return
	}
	var queue []*BNode
	queue = append(queue, root)

	for len(queue) > 0 {
		node := queue[0]
		// for _, node := range queue { // range不能感知queue的变化，如果queue遍历完不空，再遍历一次
		fmt.Println(node.Data)
		if node.LChild != nil {
			queue = append(queue, node.LChild)
		}
		if node.RChild != nil {
			queue = append(queue, node.RChild)
		}
		queue = queue[1:len(queue)]
		// }
	}
}

func LayerOrder2(root *BNode) {
	if root == nil {
		return
	}
	var queue []*BNode
	queue = append(queue, root)

	for i := 0; i < len(queue); i++ {
		node := queue[i]
		fmt.Println(node.Data)
		if node.LChild != nil {
			queue = append(queue, node.LChild)
		}
		if node.RChild != nil {
			queue = append(queue, node.RChild)
		}
		// queue = queue[1:len(queue)] // for i 不需要对queue进行截取，也就是queue会增大，时间复杂度和空间复杂度都是O(n)
	}
}
