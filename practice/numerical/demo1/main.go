package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

/*
	数值类型操作
*/

func main() {
	numerical_demo1_main()
	// fmt.Println("Int类型的最大值：", math.MaxInt32)
	// fmt.Println("Int类型的最小值：", math.MinInt32)
}

/*
	整数反转
*/

func numerical_demo1_main() {
	// numerical_demo1(321)
	fmt.Println(numerical_demo1_2(321))
}

/*
	注意点： 整数的长度，
	思路1:   将其转换成字符串，根据字符串的长度特性进行逐一变换
	TODO 未完成
*/
func numerical_demo1_1(x int32) int {
	var tmp string
	s := strconv.Itoa(int(x))
	strarr := strings.Split(s, "")
	var i int = len(strarr) - 1
	fmt.Println("获取的字符串数组是：", strarr, "长度是：", i)
	for j := 0; j < len(strarr); j++ {
		if strarr[i] != strarr[j] {
			tmp = strarr[i]
			strarr[i] = strarr[j]
			strarr[j] = tmp
			i--
		}
		fmt.Println("当前数呼呼是", math.Ceil(float64(i)))
	}
	fmt.Println("当前字符串是：", strarr)
	return i
}

// 整数反转 No.2
func numerical_demo1_2(x int) (num int) {
	for x != 0 {
		fmt.Println("num*10 + x%10：", num*10, x%10, num*10+x%10)
		num = num*10 + x%10
		x = x / 10
	}
	// 使用math 包中定义好的最大最小值
	if num > math.MaxInt32 || num < math.MinInt32 {
		return 0
	}
	fmt.Println("计算之后的num值：", num)
	return
}

// 整数反转No.3
func numerical_demo1_3(x int) (num int) {
	// An highlighted block
	var nums, newnums int
	for x != 0 { //直到x等于0，跳出循环
		a := x % 10
		newnums = nums*10 + a

		nums = newnums
		x = x / 10

		//题目要求其数值范围是 [−2^31,  2^31 − 1]。如果反转后的整数溢出，则返回 0。
		MaxInt32 := 1<<31 - 1
		MinInt32 := -1 << 31
		if nums > MaxInt32 || nums < MinInt32 {
			return 0
		}
	}
	return nums

}

/*
	报数：
*/
func demo1_countAndSay_main() {
	countAndSay(0)
}

// 报数
func countAndSay(n int) string {
	return ""
}
