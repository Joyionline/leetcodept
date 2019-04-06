package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

/*
	Leetcode题
*/

func main() {
	// dmeo1_1()
	// demo1_isPalindrome_main()
	demo1_getSum_1_main()
}

/*
题目：两数之和
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，
并返回他们的数组下标。
你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。
*/

func dmeo1_1() {
	nums := []int{2, 7, 11, 15}
	var target = 9
	rt := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				rt = append(rt, i, j)
			}
		}
	}
	fmt.Println("两数之和为target的元素下标是：", rt)
}

/*
题目：两数相加
给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例：

输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807
*/

func demo1_2() {
	// var num1, num2 *int
	for i := 0; i < 8; i++ {

	}
}

/*
	回文数判断一个整数是否是回文数。
	回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。

*/

func demo1_isPalindrome_main() {
	fmt.Println(isPalindrome(0121))
}

func isPalindrome(x int) bool {
	if x < 0 { // 当这个数为负数时一定不是回文数
		return false
	}
	var num int
	var tmp int = x
	for x != 0 {
		num = num*10 + x%10
		x = x / 10
	}
	if num > math.MaxInt32 || num < math.MinInt32 {
		return false
	}
	if tmp != num {
		return false
	}
	return true
}

//
func isPalindrome2(x int) bool {
	if x < 0 { // 当这个数为负数时一定不是回文数
		return false
	}
	var num int
	var tmp int = x
	var ok bool
	for x != 0 {
		num = num*10 + x%10
		x = x / 10
	}
	if num > math.MaxInt32 || num < math.MinInt32 {
		return false
	}
	was1 := strings.Split(strconv.Itoa(tmp), "")
	was2 := strings.Split(strconv.Itoa(num), "")
	if len(was1) == len(was2) {
		var j int = len(was2) - 1
		for i := 0; i < len(was1); i++ {
			if was1[i] == was2[j] {
				ok = true
			} else {
				ok = false
			}
			j--
		}
	}
	return ok
}

/*
两整数之和：
不使用运算符 + 和 - ​​​​​​​，计算两整数 ​​​​​​​a 、b ​​​​​​​之和。

*/

func demo1_getSum_1_main() {
	fmt.Println(getSum(1, 4))
}

func getSum(a int, b int) int {
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}
	var sum int
	sum = a & b
	return getSum(a^b, sum<<1)
}
