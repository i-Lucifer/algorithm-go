package linked_list

// 循环直接逆序 相邻节点交换指向
// 时间复杂度O(n)
// 空间复杂度O(1)
func RSortExchange(head *LNode) {
	if head == nil || head.Next == nil {
		return
	}

	var tmp *LNode
	var pre *LNode    // 存储交换后的第一个节点
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

// 递归 先逆序子节点，然后把当前节点放到最后
func RSortFunc(head *LNode) {
	NewHead := RSortFuncChild(head.Next)
	head.Next = NewHead

}
func RSortFuncChild(node *LNode) *LNode {
	if node == nil || node.Next == nil {
		return node
	}
	NewHead := RSortFuncChild(node.Next)
	node.Next.Next = node
	node.Next = nil
	return NewHead
}

// 插入法 遍历到第n个节点，插入最前面
// 这个比较像构建一个新的链表
func RSortInsert(head *LNode) {
	if head == nil || head.Next == nil {
		return
	}
	var tem *LNode       // 临时节点
	var next *LNode      // 下一个节点
	tem = head.Next.Next // 第一个节点的指针
	head.Next.Next = nil // 临时移除第一个的指针

	for tem != nil {
		next = tem.Next
		tem.Next = head.Next
		head.Next = tem
		tem = next
	}
}
