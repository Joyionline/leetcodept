package main

import (
	"fmt"
	"sort"
)

func main() {
	demo6_maximumProduct_main()
}

func demo6_maximumProduct_main() {
	// 测试用例 {2, 2, 4, 5, 6, 6} | {1, 2, 3, 4} | {-4, -3, -2, -1, 60} | {-1, -2, -3}
	nums := []int{-1000, -1000, -1000}
	fmt.Println("该数组运算后最大的乘积是：", maximumProduct(nums))

}

/*
	分析情况：
		1. 全部是负数
			最大的三个负数相乘即最大
		2. 全部是正数
			最大的三个正数相乘即最大
		3. 正数个数小于3个，负负正
		4. 最小结果，数组中三个数均为范围内最小值
*/

func maximumProduct(nums []int) int {
	sort.Ints(nums)
	var result int = -1000000000
	var numlen = len(nums)
	if result < nums[numlen-1]*nums[numlen-2]*nums[numlen-3] {
		result = nums[numlen-1] * nums[numlen-2] * nums[numlen-3]
	}
	if result < nums[0]*nums[1]*nums[numlen-1] {
		result = nums[0] * nums[1] * nums[numlen-1]
	}
	return result
}
