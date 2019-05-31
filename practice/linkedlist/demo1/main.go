package main

import "fmt"

//TODO today

func main() {
	demo1_reverseList()
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func demo1_reverseList() {
	InsertData()
}

// 链表反转
/*
	整个循环过程分析：
	1  nil 1 2 3 4 5 >> 2 3 4 5 1 nil
	2  2 3 4 5 1 nil >> 3 4 5 1 2 nil
	3  3 4 5 1 2 nil >> 4 5 1 2 3 nil
	4  4 5 1 2 3 nil >> 5 1 2 3 4 nil
	5  5 1 2 3 4 nil >> 1 2 3 4 5 nil
*/
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	var cur = head
	for cur != nil {
		nexttemp := cur.Next // 以当前为h1, 下一个为h2 作为临时数据
		cur.Next = pre       // 当前的下一个值指向前一个，即当前前一个为pre,即 nil
		pre = cur            // 将当前的值赋值给前一个值 将nil 向后转移
		cur = nexttemp       // 将之前缓存的下一个赋值给当前
		// pre, cur, cur.Next = cur, cur.Next, pre
	}
	return pre
}

func reverseList_2(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		pre, cur, cur.Next = cur, cur.Next, pre
	}
	return pre
}

// 插入数据
func InsertData() {
	h1 := new(ListNode)
	h1.Val = 1
	h2 := new(ListNode)
	h2.Val = 2
	h3 := new(ListNode)
	h3.Val = 3
	h4 := new(ListNode)
	h4.Val = 4
	h5 := new(ListNode)
	h5.Val = 5
	h1.Next = h2
	h2.Next = h3
	h3.Next = h4
	h4.Next = h5

	fmt.Println("当前链表数据是：", h1)
	// pre := reverseList(h1)
	pre := reverseList_2(h1)
	fmt.Println("反转后的链表数据：", pre)
}
