package main

import (
	"fmt"
	"sort"
)

func main() {
	demo5_singleNumber_main()
}

/*
	只出现一次的数字
	给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
说明：
你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？
示例 1:
输入: [2,2,1]
输出: 1
示例 2:
输入: [4,1,2,1,2]
输出: 4

*/

func demo5_singleNumber_main() {
	// 测试用例{2, 2, 1} ,{4, 1, 2, 1, 2}
	nums := []int{1, 2, 3, 3, 1}
	fmt.Println("只出现一次的数据：", singleNumber(nums))
}

func singleNumber(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	sort.Ints(nums)
	var lastnum int
	for i := 0; i < len(nums); i++ {
		if i != 0 && i != len(nums)-1 {
			if nums[i] != nums[i-1] && nums[i] != nums[i+1] {
				lastnum = nums[i]
			}
		}
		if i == 0 {
			lastnum = nums[i]
		}
		if i == len(nums)-1 && nums[i] != nums[i-1] {
			lastnum = nums[i]
		}
	}
	return lastnum
}
