package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// demo2_countSegments_main()
	// demo2_compress_main()
	// demo2_longestCommonPrefix_main()
	firstUniqChar_main()
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

/*
	最长公共前缀
	编写一个函数来查找字符串数组中的最长公共前缀。
	如果不存在公共前缀，返回空字符串 ""。
	示例 1:
	输入: ["flower","flow","flight"]
	输出: "fl"
	示例 2:
	输入: ["dog","racecar","car"]
	输出: ""
	解释: 输入不存在公共前缀。
*/
func demo2_longestCommonPrefix_main() {
	strarr := []string{"flower", "flow", "flight"}
	fmt.Println("该字符串数组中最长的公共前缀是：", longestCommonPrefix(strarr))

	// str := "hello"
	// fmt.Println(strings.Index(str, "e"))

}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	var comstr string = strs[0]
	for i := 1; i < len(strs); i++ {
		for strings.Index(strs[i], comstr) != 0 {
			comstr = comstr[0 : len(comstr)-1]
			if comstr != "" {
				return ""
			}
		}
		fmt.Println("截取数据：", comstr)
	}
	return comstr
}

/*
	字符串中的第一个唯一字符
	给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。
	案例:
	s = "leetcode"
	返回 0.
	s = "loveleetcode",
	返回 2.
*/

func firstUniqChar_main() {
	// 测试用例  "leetcode"  | "loveleetcode" | "oooooo" | "dddccdbba" | "cc"
	s := "dddccdbba"
	fmt.Println("当前字符串中第一个唯一字符索引是：", firstUniqChar_2(s))
}

// 暴力解法超时
func firstUniqChar(s string) int {
	var count int
	var laststr string
	for i := 0; i < len(s); i++ {
		if string(s[i]) != laststr {
			count = 0
		}
		for j := 0; j < len(s); j++ {
			if s[i] == s[j] {
				count++
			}
			if count > 1 {
				continue
			}
			if j == len(s)-1 && count == 1 {
				return i
			}
		}
	}
	return -1
}

// map 实现
func firstUniqChar_2(s string) int {
	h := map[rune]int{}
	for _, v := range s {
		h[v]++
	}
	for k, v := range s {
		if h[v] == 1 {
			return k
		}
	}
	return -1
}
