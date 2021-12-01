package tree

import (
	"fmt"
)

// 广度优先
// 层次遍历
// 二叉树的层次遍历一般需要借助队列
func LayerOrder(node *BNode) {
	if node == nil {
		return
	}
	var queue []*BNode
	queue = append(queue, node)

	for len(queue) > 0 {
		node, queue = pull(queue)
		// for _, node := range queue { // range不能感知queue的变化，如果queue遍历完不空，再遍历一次
		fmt.Println(node.Data)
		if node.LChild != nil {
			queue = append(queue, node.LChild)
		}
		if node.RChild != nil {
			queue = append(queue, node.RChild)
		}
		// queue = queue[1:len(queue)]
		// }
	}
}

// 缺点，队列的值没有被取出，空间复杂度为O(n)
func LayerOrder2(node *BNode) {
	if node == nil {
		return
	}
	var queue []*BNode
	queue = append(queue, node)

	for i := 0; i < len(queue); i++ {
		node := queue[i] // 遍历取值
		fmt.Println(node.Data)
		if node.LChild != nil {
			queue = append(queue, node.LChild)
		}
		if node.RChild != nil {
			queue = append(queue, node.RChild)
		}
	}
}
