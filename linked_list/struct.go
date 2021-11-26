package linked_list

// 左节点
type LNode struct {
	Data int
	Next *LNode
}

// 头节点 这里头节点不存储数据
func NewNode(len int) *LNode {
	head := new(LNode)
	CreateNode(head, len)
	return head
}

func CreateNode(head *LNode, len int) {
	lNode := head
	for i := 1; i < len; i++ {
		node := new(LNode)
		node.Data = i
		lNode.Next = node
		lNode = node
	}
}
