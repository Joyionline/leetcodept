package main

import (
	"fmt"
	"sort"
)

func main() {
	demo3_dominantIndex_main()
}

/*
	至少是其他数字两倍的最大数
	在一个给定的数组nums中，总是存在一个最大元素 。

	查找数组中的最大元素是否至少是数组中每个其他数字的两倍。
	如果是，则返回最大元素的索引，否则返回-1。
	示例 1:
	输入: nums = [3, 6, 1, 0]
	输出: 1
	解释: 6是最大的整数, 对于数组中的其他整数,
	6大于数组中其他元素的两倍。6的索引是1, 所以我们返回1.

	思考：
		1. 不适用第三方排序
		2. 比较仅需要比较最大值和第二大的两个是否符合就可
*/
func demo3_dominantIndex_main() {
	//  测试用例： {3, 6, 1, 0} {1, 2, 3, 4}
	nums := []int{1, 1}
	fmt.Println("返回值：", dominantIndex(nums))
}

func dominantIndex(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	tmp := make([]int, 0)
	tmp = append(tmp, nums...)
	var flag bool
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[len(nums)-1] >= nums[i]*2 {
			flag = true
		} else {
			flag = false
		}
	}
	if flag {
		for j := 0; j < len(tmp); j++ {
			if nums[len(nums)-1] == tmp[j] {
				return j
			}
		}
	}
	return -1
}
