package tree

func pop(stack []*BNode) (*BNode, []*BNode) {
	lenth := len(stack)
	node := stack[lenth-1]
	stack = stack[0 : lenth-1]
	return node, stack
}
