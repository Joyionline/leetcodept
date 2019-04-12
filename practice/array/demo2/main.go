package main

import (
	"fmt"
	"sort"
)

func main() {
	// demo2_searchInsert_1()
	// demo2_plusOne_1_main()
	// demo2_containsDuplicate_1_main()
	// demo2_containsNearbyDuplicate_1_main()
	// demo2_thirdMax_main()
	// demo2_findKthLargest_main()
	// demo2_findMaxConsecutiveOnes_main()
	demo2_longestOnes_main()
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

/*
给定一个整数数组和一个整数 k，判断数组中是否存在两个不同的索引 i 和 j，
使得 nums [i] = nums [j]，并且 i 和 j 的差的绝对值最大为 k。

*/

func demo2_containsNearbyDuplicate_1_main() {
	nums := []int{0, 1, 1, 1}
	k := 1
	fmt.Println("result：", containsNearbyDuplicate(nums, k))
}

func containsNearbyDuplicate(nums []int, k int) bool {
	var tmp int
	// nm := make(map[int]int, 0)
	// s := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				if tmp == 0 {
					tmp = i
				}
				tmp += j
			}
		}
	}
	fmt.Println("当前的tmp是：", tmp)
	/* 	for j := 0; j < len(nums); j++ {
		if _, ok := nm[nums[i]]; ok {
			s = append(s, i)
		}
		nm[nums[i]] = i
	} */

	// fmt.Println("存在元素的索引是：", s)
	return true
}

/*
找到第三个最大的数
给定一个非空数组，返回此数组中第三大的数。
如果不存在，则返回数组中最大的数。要求算法时间复杂度必须是O(n)。

*/

func demo2_thirdMax_main() {
	// 测试用例[1,3,5,6,6,1] {2, 2, 3, 1}
	nums := []int{1, 3, 5, 6, 6, 1}
	fmt.Println("当前数组中第三大值是：", thirdMax(nums))
}

func thirdMax(nums []int) int {
	n := make([]int, 0) // 初始化空切片
	sort.Ints(nums)     // 对原始数组进行排序 使得不相邻的两个相等的数可以相邻排列
	i := 0
	n = append(n, nums[0])
	for j := 1; j < len(nums); j++ {
		if nums[i] != nums[j] {
			n = append(n, nums[j])
			i++
			nums[i] = nums[j]
		}
	}
	ln := len(n)
	if ln <= 2 {
		return n[ln-1]
	}
	return n[ln-3]
}

/*
	数组中的第K个最大元素
	在未排序的数组中找到第 k 个最大的元素。
	请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。

	考虑只有一个元素的情况
*/

func demo2_findKthLargest_main() {
	// 测试用例  输入: [3,2,1,5,6,4] 和 k = 2  输出: 5  输入: [3,2,3,1,2,4,5,5,6] 和 k = 4 输出: 4 ; {1} out:1
	//   in: {2,1} out: {2}
	nums := []int{1, 3}
	k := 3
	fmt.Printf("当前数组中第%d大的数是: %d\n", k, findKthLargest(nums, k))
}

func findKthLargest(nums []int, k int) int {
	ln := len(nums)
	if ln == 1 {
		return nums[0]
	}
	sort.Ints(nums)
	return nums[ln-k]
}

/*
	最大连续1的个数

*/

func demo2_findMaxConsecutiveOnes_main() {
	// 测试用例；{1,0,1,1,0,1},{1, 1, 0, 1, 1, 1, 0, 1, 1, 1, 1, 1, 8}
	// {1, 0, 1, 1, 0, 1, 1, 1, 1, 1, 1, 0, 1, 0, 1}
	// {1}
	nums := []int{1, 1, 0, 1, 1, 1, 0, 1, 1, 1, 1, 1, 0}
	fmt.Println("当前数组中最大连续1的个数是：", findMaxConsecutiveOnes(nums))
}

func findMaxConsecutiveOnes(nums []int) int {
	var count, ic int
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			ic++
			// 以下判断加在此避免漏掉第一个数即为符合值
			if ic >= count {
				count = ic
			}
		} else {
			ic = 0
		}
	}
	fmt.Println("当前值是：", count)
	return count
}

/*
	TODO: to be continued
	最大连续1的个数 III
	示例 1：
	输入：A = [1,1,1,0,0,0,1,1,1,1,0], K = 2
	输出：6
	解释：
	[1,1,1,0,0,1,1,1,1,1,1]
	粗体数字从 0 翻转到 1，最长的子数组长度为 6。
	示例 2：

	输入：A = [0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1], K = 3
	输出：10
	解释：
	[0,0,1,1,1,1,1,1,1,1,1,1,0,0,0,1,1,1,1]
	粗体数字从 0 翻转到 1，最长的子数组长度为 10。
*/

func demo2_longestOnes_main() {
	nums := []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}
	k := 2
	fmt.Println("该数组中最长的子数组长度为：", longestOnes(nums, k))
}

func longestOnes(A []int, K int) int {
	Tmp := make([]int, 0)
	var count, max int
	var num int
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A); j++ {
			if A[j] == 0 && num <= K {
				A[j] = 1
				num++
				Tmp = append(Tmp, j)
			}
			if num == K {
				break
			}
		}
		fmt.Println("修改的结果：", A)
		if A[i] == 1 {
			max++
			if max >= count {
				count = max
			}
		} else {
			max = 0
		}
		fmt.Println("shdfnsndfionsdf:", Tmp)
		for n := 0; n < len(Tmp); n++ {
			A[Tmp[n]] = 0
		}
		fmt.Println("当前更新的数据：", A)
	}
	return count
}
