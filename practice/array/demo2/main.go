package main

import "fmt"

func main() {
	demo2_searchInsert_1()
}

/*
给定一个排序数组和一个目标值，
在数组中找到目标值，并返回其索引。
如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
*/

func demo2_searchInsert_1() {
	nums := []int{1, 3, 5, 6}
	target := 0
	n := searchInsert(nums, target)
	fmt.Println("返回的值：", n)
}

func searchInsert(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
		if nums[i] == target {
			return i
		} else {
			if target < nums[i] {
				return i
			}
		}
	}
	return len(nums)
}
