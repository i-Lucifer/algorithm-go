package tree

// 出栈
func pop(stack []*BNode) (*BNode, []*BNode) {
	lenth := len(stack)
	node := stack[lenth-1]
	stack = stack[0 : lenth-1]
	return node, stack
}

// 出队
func pull(queue []*BNode) (*BNode, []*BNode) {
	node := queue[0]
	queue = queue[1:len(queue)]
	return node, queue
}
