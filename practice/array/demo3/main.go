package main

import (
	"fmt"
	"sort"
)

func main() {
	// demo3_dominantIndex_main()
	// demo3_PrintArray_main()
	// demo3_findDiagonalOrder_main()
	// demo3_findRestaurant_main()
	demo3_findLHS_main()
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

/*
	TODO:  关于初始值问题 ，使用make声明的二维数组中没有初始值，因为是切片，所以需要手动赋初始值
	关于二维数组
*/
func demo3_PrintArray_main() {
	fmt.Println("Example I:")
	a := make([][]int, 2, 5) // 该声明方式可兼容所有方法的传递(方法中可不指定长度)
	fmt.Println("该数组是:", a)
	PrintArray(a)
	/* 	fmt.Println("Example II:")
	   	b := make([][]int, 2)
	   	PrintArray(b) */

}

func PrintArray(a [][]int) {
	for i := 0; i < len(a); i++ {
		m2 := make([]int, 2, 5)
		// fmt.Println(a[i])
		a[i] = m2
	}
	fmt.Println("当前数组是:", a)
	/* 	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			fmt.Print(a[i][j], " ")
		}
		fmt.Println()
	} */
}

/*
对角线遍历
给定一个含有 M x N 个元素的矩阵（M 行，N 列），
请以对角线遍历的顺序返回这个矩阵中的所有元素，对角线遍历如下图所示。

*/

func demo3_findDiagonalOrder_main() {
	a := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	// fmt.Println("当前数组是：", a)
	fmt.Println("新数组数据是:", findDiagonalOrder(a))
}

func findDiagonalOrder(matrix [][]int) []int {
	newarr := make([]int, 0)
	for i := 0; i < len(matrix); i++ {
		fmt.Println(matrix[i])
	}
	return newarr
}

/*
	关于双指针
*/

func demo3_minSubArrayLen() {
	s := 7
	nums := []int{2, 3, 1, 2, 4, 3}
	fmt.Println(s, nums)
}

func minSubArrayLen(s int, nums []int) int {
	// var len, i int
	/* for j := 0; j < len(nums); j++ {

	} */

	return 0
}

/*
	两个列表的最小索引总和
	假设Andy和Doris想在晚餐时选择一家餐厅，并且他们都有一个表示最喜爱餐厅的列表，
	每个餐厅的名字用字符串表示。
	你需要帮助他们用最少的索引和找出他们共同喜爱的餐厅。
	如果答案不止一个，则输出所有答案并且不考虑顺序。 你可以假设总是存在一个答案。

	输入:
	["Shogun", "Tapioca Express", "Burger King", "KFC"]
	["Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"]
	输出: ["Shogun"]
	解释: 他们唯一共同喜爱的餐厅是“Shogun”。
*/

func demo3_findRestaurant_main() {
	/*
		测试用例：
		{"Shogun", "Tapioca Express", "Burger King", "KFC"} || {"KFC", "Shogun", "Burger King"}

	*/
	list1 := []string{"Shogun", "Tapioca Express", "Burger King", "KFC"}
	list2 := []string{"KFC", "Burger King", "Tapioca Express", "Shogun"}
	rlist := findRestaurant(list1, list2)
	fmt.Println("他们两个共同的餐厅是：", rlist)
}

func findRestaurant(list1 []string, list2 []string) []string {
	var common []string
	var index int
	for i := 0; i < len(list1); i++ {
		for j := 0; j < len(list2); j++ {
			if list1[i] == list2[j] {
				if index == 0 && len(common) == 0 {
					index = i + j
					common = append(common, list1[i])
				} else {
					if i+j == index {
						common = append(common, list1[i])
					}
				}

			}
		}
	}
	return common
}

/*
	TODO： to be coutinued
	最长的和谐子序列
	和谐数组是指一个数组里元素的最大值和最小值之间的差别正好是1。
	现在，给定一个整数数组，你需要在所有可能的子序列中找到最长的和谐子序列的长度。
	示例 1:

	输入: [1,3,2,2,5,2,3,7]
	输出: 5
	原因: 最长的和谐数组是：[3,2,2,2,3].
*/

func demo3_findLHS_main() {
	//  测试用例：{1, 3, 2, 2, 5, 2, 3, 7}  {1,2,3,4}
	nums := []int{1, 2, 3, 4}
	fmt.Println("最长的和谐数组长度是：", findLHS(nums))
}

func findLHS(nums []int) int {
	sort.Ints(nums)
	var newslice []int
	var recount, max, lastnum int
	for i := 0; i < len(nums); i++ {
		newslice = append(newslice, nums[i])
		if nums[i] == lastnum && i > 0 {
			if max >= recount {
				max++
				recount = max
				newslice = append(newslice, nums[i])
			} else {
				recount++
			}
		} else {
			newslice = nil
			max = 0
			newslice = append(newslice, nums[i])
			recount = len(newslice)
		}
		lastnum = nums[i]
	}
	if len(newslice) > 0 {
		var countone, counttwo int
		for j := 0; j < len(nums); j++ {
			if nums[j] == newslice[0]-1 {
				countone++
			}
			if nums[j] == newslice[0]+1 {
				counttwo++
			}
		}
		if countone > counttwo {
			recount += countone
		} else if countone < counttwo {
			recount += counttwo
		} else {
			recount += countone
		}
	}
	return recount
}
