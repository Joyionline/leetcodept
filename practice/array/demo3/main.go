package main

import (
	"fmt"
	"sort"
)

func main() {
	// demo3_dominantIndex_main()
	// demo3_PrintArray_main()
	demo3_findDiagonalOrder_main()
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
