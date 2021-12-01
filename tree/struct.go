package tree

type BNode struct {
	Data   interface{}
	LChild *BNode
	RChild *BNode
}

func NewNode() *BNode {
	return &BNode{}
}
