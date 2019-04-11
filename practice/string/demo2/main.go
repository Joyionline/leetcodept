package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// demo2_countSegments_main()
	demo2_compress_main()
}

/*
字符串中的单词数
	统计字符串中的单词个数，这里的单词指的是连续的不是空格的字符。
	请注意，你可以假定字符串里不包括任何不可打印的字符。
	示例:
	输入: "Hello, my name is John"
	输出: 5
*/

func demo2_countSegments_main() {
	// 测试用例：", , , ,        a, eaefa"
	s := "Hello,  my  name  is John"
	fmt.Println("该字符串中的单词数是：", countSegments(s), "个")
}

func countSegments(s string) int {
	strarr := strings.Split(s, " ")
	fmt.Println("档期：", strarr)
	var count int
	for i := 0; i < len(strarr); i++ {
		if strarr[i] != "" {
			count++
		}
	}
	return count
}

/*
	TODO   未完成
	压缩字符串
	给定一组字符，使用原地算法将其压缩。
	压缩后的长度必须始终小于或等于原数组长度。
	数组的每个元素应该是长度为1 的字符（不是 int 整数类型）。
	在完成原地修改输入数组后，返回数组的新长度。
	进阶：
	你能否仅使用O(1) 空间解决问题？
*/

func demo2_compress_main() {
	s := []byte{'a', 'a', 'b', 'b', 'b', 'c', 'c', 'c'}
	fmt.Println("当前的字符串数组长度,", compress(s))
}

func compress(chars []byte) int {
	if len(chars) <= 1 {
		return len(chars)
	}
	var i, count int
	for j := 1; j < len(chars); j++ {
		if chars[i] == chars[j] {
			count++
			if j == len(chars)-1 {
				b := strconv.Itoa(count + 1)
				chars[i] = byte(b[0])
				break
			}
		} else {
			b := strconv.Itoa(count + 1)
			chars[j-count] = byte(b[0])
			chars[j-count+1] = chars[j]
			count = 0
		}
		i++
	}
	fmt.Println("当前字符串：", string(chars))
	return len(chars)
}
