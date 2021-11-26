package linked_list

import "fmt"

// 递归打印
func PrintByFunc(head *LNode) {
	if head.Data != 0 { // 头节点这里不存储数据，也不输出 这里的缺陷是，如果数据为0，就不输出了
		fmt.Print(head.Data, " ")
	}
	if head.Next != nil {
		PrintByFunc(head.Next) // 传递指针给下一次输出
	}
}

// 循环打印
func PrintNode(head *LNode) {
	// fmt.Println(head.Data)  // 头节点这里不存储数据，也不输出
	next := head.Next
	for next != nil {
		fmt.Print(next.Data, " ")
		next = next.Next // 传递指针给下一次循环
	}
	fmt.Println()
}
