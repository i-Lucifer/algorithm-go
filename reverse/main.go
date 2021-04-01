package main

import (
	"fmt"
)

// 左节点
type LNode struct {
	Data int
	Next *LNode
}

func main() {
	head := &LNode{}
	CreateNode(head, 7)
	PrintNode1(head)
	Revserse1(head)
	fmt.Println("------")
	PrintNode2(head)
}

func CreateNode(head *LNode, num int) {
	lNode := head
	for i := 1; i < num; i++ {
		node := new(LNode)
		node.Data = i
		lNode.Next = node
		lNode = node
	}
}

// 递归法
// 需要入栈和出栈(压和弹)
func PrintNode1(head *LNode) {
	if head.Data != 0 {
		fmt.Println(head.Data)
	}
	if head.Next != nil {
		PrintNode1(head.Next)
	}
}

func PrintNode2(head *LNode) {
	next := head.Next
	for next != nil {
		fmt.Println(next.Data)
		next = next.Next
	}

}

// 循环直接逆序
// 时间复杂度O(n)
// 空间复杂度O(1)
func Revserse1(head *LNode) {
	if head == nil || head.Next == nil {
		return
	}

	var tmp *LNode
	var pre *LNode
	next := head.Next // 节点a

	// 每次循环 当前节点 指向前一个节点 原本a 指b 现在b指a
	for next != nil {
		tmp = next.Next // 临时记录下一个节点b的地址
		next.Next = pre // 修改指针 为前一个节点
		pre = next      // 当前节点 会作为下一次循环的前一个节点
		next = tmp      // 传递地址 循环下一个节点
	}

	head.Next = pre
}

// 递归
func Revserse2(head *LNode) {

}

// 插入法
func Revserse3(head *LNode) {

}
