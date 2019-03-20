package main

import "fmt"

/*
	Leetcode题
*/

func main() {
	dmeo1_1()
}

/*
题目：两数之和
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，
并返回他们的数组下标。
你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。
*/

func dmeo1_1() {
	nums := []int{2, 7, 11, 15}
	var target = 9
	rt := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				rt = append(rt, i, j)
			}
		}
	}
	fmt.Println("两数之和为target的元素下标是：", rt)
}
