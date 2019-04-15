package main

import "fmt"

/*
	关于链表操作
*/

func main() {
	demo1_List_main()
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func demo1_List_main() {
	head := new(ListNode)
	head.Val = 1
	L2 := new(ListNode)
	L2.Val = 2
	L3 := new(ListNode)
	L3.Val = 3
	head.Next = L2
	L2.Next = L3
	fmt.Printf("当前链表数据是：%v\n", head.Val)
	fmt.Println("经过反转后的链表", reverserList(head).Val)
}

/*
	TODO: 反转
	反转链表
*/

func reverserList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr.Next != nil {
		nextTemp := curr.Next
		curr.Next = prev
		prev = curr
		curr = nextTemp
	}
	return prev
}
