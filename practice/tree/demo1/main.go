package main

import "fmt"

func main() {
	demo1_main_tree_1()
}

/**
对称二叉树
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func demo1_main_tree_1() {
	var root *TreeNode
	fmt.Println("是否为平衡二叉树：", isSymmetric(root))
}

func isSymmetric(root *TreeNode) bool {
	return isMirror(root, root)
}

func isMirror(t1 *TreeNode, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 == nil || t2 == nil {
		return false
	}
	return (t1.Val == t2.Val) && isMirror(t1.Right, t2.Left) && isMirror(t1.Left, t2.Right)
}
