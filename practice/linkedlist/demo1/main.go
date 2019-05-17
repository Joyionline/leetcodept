package main

import "fmt"

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
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	var cur = head
	for cur != nil {
		nexttemp := cur.Next
		cur.Next = pre
		pre = cur
		cur = nexttemp
		// pre, cur, cur.Next = cur, cur.Next, pre
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
	pre := reverseList(h1)
	fmt.Println("反转后的链表数据：", pre)
}
