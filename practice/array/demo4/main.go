package main

import (
	"fmt"
	"sort"
)

func main() {
	// demo4_arrayPairSum_main()
	demo4_sortArrayByParity_main()
}

/*
	最小差值
*/

func demo4_smallestRangeI_main() {
	A := []int{1}
	K := 0
	fmt.Println("返回的最小差值：", smallestRangeI(A, K))
}

func smallestRangeI(A []int, K int) int {
	return 0
}

/*
数组拆分 I
	给定长度为 2n 的数组, 你的任务是将这些数分成 n 对,
	 例如 (a1, b1), (a2, b2), ..., (an, bn) ，使得从1 到 n 的 min(ai, bi) 总和最大。

	示例 1:
	输入: [1,4,3,2]
	输出: 4
	解释: n 等于 2, 最大总和为 4 = min(1, 2) + min(3, 4).
*/

func demo4_arrayPairSum_main() {
	// 测试用例： {1, 4, 3, 2}
	nums := []int{0, 0}
	fmt.Println("拆分之后的最大总和是：", arrayPairSum(nums))
}

func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	var sum int
	fmt.Println("当前的数组：", nums)
	// var lastres int
	for i := 0; i < len(nums); i++ {
		fmt.Println(i % 2)
		if i%2 == 0 {
			sum += nums[i]
			// lastres = nums[i] % 2
		}
		/* if lastres != 0 {
			lastres = 0
		} */
	}
	fmt.Println("总和是：", sum)
	return sum
}

func demo4_sortArrayByParity_main() {
	// 测试用例
	nums := []int{3, 1, 2, 4}
	fmt.Println("排序之后的数组：", sortArrayByParity(nums))
}

func sortArrayByParity(A []int) []int {
	newArr0 := make([]int, 0)
	newArr1 := make([]int, 0)
	for i := 0; i < len(A); i++ {
		if A[i]%2 == 0 {
			newArr0 = append(newArr0, A[i])
		} else {
			newArr1 = append(newArr1, A[i])
		}
	}
	fmt.Println("偶数数组：", newArr0)
	fmt.Println("奇数数组：", newArr1)
	return append(newArr0, newArr1...)
}
