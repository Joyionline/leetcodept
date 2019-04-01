package main

import (
	"fmt"
)

func main() {
	// demo2_searchInsert_1()
	// demo2_plusOne_1_main()
	demo2_containsDuplicate_1_main()
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

/*
给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。

最高位数字存放在数组的首位， 数组中每个元素只存储一个数字。

你可以假设除了整数 0 之外，这个整数不会以零开头。

示例 1:

输入: [1,2,3]
输出: [1,2,4]
解释: 输入数组表示数字 123。

*/

func demo2_plusOne_1_main() {
	arr := []int{1, 2, 9, 9}
	newarr := plusOne(arr)
	fmt.Println("新的数组是：", newarr)
}

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] == 9 {
			digits[i] = 0
		} else {
			digits[i] += 1
			return digits
		}
	}
	return append([]int{1}, digits...)
}

/*
给定一个整数数组，判断是否存在重复元素。

如果任何值在数组中出现至少两次，函数返回 true。如果数组中每个元素都不相同，则返回 false。

示例 1:

输入: [1,2,3,1]
输出: true
*/

func demo2_containsDuplicate_1_main() {
	nums := []int{7, 3, 2, 1, 2}
	fmt.Println(containsDuplicate(nums))
}

func containsDuplicate(nums []int) bool {
	m := make(map[int]int, 0)
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = nums[i]
	}
	if len(nums) == len(m) {
		return false
	}
	return true
}
