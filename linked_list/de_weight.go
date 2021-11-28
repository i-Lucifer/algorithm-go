package linked_list

import "fmt"

// 查重
func QueryWeight(head *LNode) {
	next := head.Next
	for next != nil {
		tmpdata := next.Data
		tmpnode := next.Next
		for tmpnode != nil {
			if tmpdata == tmpnode.Data {
				fmt.Println(tmpdata)
				break // 7 重复了两次，如果不加break，会输出 2 + 1次7 加上会输出1+1次7
			}
			tmpnode = tmpnode.Next
		}
		next = next.Next // 传递指针给下一次循环
	}
}

// 去重
func DeWeight(head *LNode) {
	next := head.Next
	for next != nil {
		tmpdata := next.Data
		tmpnode := next.Next
		for tmpnode != nil {
			if tmpdata == tmpnode.Data {
				// 删除节点
			}
			tmpnode = tmpnode.Next
		}
		next = next.Next // 传递指针给下一次循环
	}
}
