package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 1, 2}
	// fmt.Println(DelDuplicates_demo1(nums))
	fmt.Println(DelDuplicates_demo2(nums))
}

// 从排序数组中删除重复项
// TODO 本地测试没问题，leetcode上输出不一致
func DelDuplicates_demo1(nums []int) int {
	newnums := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		repeat := false
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newnums = append(newnums, nums[i])
		}
	}
	fmt.Println("新的数组数据：", newnums)
	return len(newnums)
}

// 从排序数组中删除重复项
/*
	双指针法：
		通过定义一个变量在一个循环中不断增加值，可以控制时间复杂度
*/

func DelDuplicates_demo2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var i int = 0
	for j := 0; j < len(nums); j++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}
