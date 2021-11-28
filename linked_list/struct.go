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

// 数组转链表
func ArrToLink(arr []int) *LNode {
	head := new(LNode)
	lNode := head
	for _, item := range arr {
		node := new(LNode)
		node.Data = item
		lNode.Next = node
		lNode = node
	}
	return head
}
