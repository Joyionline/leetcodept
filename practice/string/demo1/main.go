package main

import (
	"fmt"
	"strings"
)

func main() {
	// fmt.Println(lengthOfLongestSubstring("cb"))
	// demo1_strStr_1()
	// demo1_lengthOfLastWord_main()
	// demo1_reverseString_main()

}

/*
	给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

		示例 1:
		输入: "abcabcbb"
		输出: 3
		解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。

		示例 2:
		输入: "bbbbb"
		输出: 1
		解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。

		示例 3:
		输入: "pwwkew"
		输出: 3
		解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
			请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。

思路：
	将字符串进行切分，得到一个字符串切片，然后将切片中的数据进行遍历，每次遍历找到相邻的不重复的字符串，
	当遇到相同的字符串便停止，同时记录当前字符串的长度，如果遇到更长的字符串将覆盖之前的长度，同时从下一个字符串开始
	再次遍历
	1. 通过遍历将切片中的每个字符进行对比
	2. 利用map无重复的特性，将数据进行分别存储
*/
// TODO 注意点：最长子串
func demo1_1(s string) int {
	var strlen int // 字串长度
	// str := "bbbbb"
	if len(s) > 0 && len(strings.TrimSpace(s)) == 0 || len(s) == 1 {
		return 1
	} else {
		if len(s) == 0 {
			return 0
		}
	}
	strslice := strings.Split(s, "") // [a b c a b c b b]
	fmt.Println("切分之后的数据：", strslice, len(strslice))
	// newslice := make([]string, 0)
	newmap := make(map[string]string, 0)
	for i := 0; i < len(strslice); i++ {
		for j := i + 1; j < len(strslice); j++ {
			if strmap_exist(strslice[j], newmap) {
				if strlen < len(newmap) {
					newmap = map_delete(newmap)
				}
				newmap[strslice[j]] = strslice[j]
				strlen = len(strslice)
				break
			}
			newmap[strslice[j]] = strslice[j]
			// newslice = append(newslice, strslice[i])
			strlen = len(strslice)
			// fmt.Printf("当前新字符串是%s，长度是：%d", newslice, strlen)
		}
	}
	fmt.Println("新的字符串切片是：", newmap)
	return len(newmap)
	// 利用map特性存放数据
	/* 	strmap := make(map[string]string, 0)
	   	for i := 0; i < len(strslice); i++ {
	   		strmap[strslice[i]] = strslice[i]
	   	} */
	// fmt.Println("新存储的map是：", strmap, "长度是：", len(strmap))
}

// 判断字符串是否存在  切片
func strslice_exist(s string, strslice []string) bool {
	for _, str := range strslice {
		if str == s {
			return true
		}
	}
	return false
}

// 判断字符串是否存在 map
func strmap_exist(s string, strmap map[string]string) bool {
	fmt.Println("当前的map数据：", strmap)
	for _, str := range strmap {
		fmt.Println("当前比较的数据是：", str, s)
		if str == s {
			return true
		}
	}
	return false
}

// 删除不合适的元素
func map_delete(m map[string]string) map[string]string {
	for _, str := range m {
		delete(m, str)
	}
	return m
}

// 上题
/* func demo1_1_1(s string) {
	hash := make([]int, 0)
	var i, j, max int

} */

func lengthOfLongestSubstring(s string) int {
	if len(s) > 0 && len(strings.TrimSpace(s)) == 0 || len(s) == 1 {
		return 1
	} else {
		if len(s) == 0 {
			return 0
		}
	}
	var strlen int // 字串长度
	strslice := strings.Split(s, "")
	newmap := make(map[string]string, 0)
	for i := 0; i < len(strslice); i++ {
		for j := i + 1; j < len(strslice); j++ {
			for _, str := range newmap {
				if str == strslice[j] {
					if strlen < len(newmap) {
						for _, str := range newmap {
							delete(newmap, str)
						}
					}
					newmap[strslice[j]] = strslice[j]
					strlen = len(strslice)
					break
				}
			}
			newmap[strslice[j]] = strslice[j]
			strlen = len(strslice)
		}
	}
	return len(newmap)
}

/*
 实现strStr()
给定一个 haystack 字符串和一个 needle 字符串，
在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如果不存在，则返回  -1。

示例 1:
输入: haystack = "hello", needle = "ll"
输出: 2

示例 2:
输入: haystack = "aaaaa", needle = "bba"
输出: -1
*/

func demo1_strStr_1() {
	fmt.Println(strStr("a", "a"))
}

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	i := strings.Index(haystack, needle)
	fmt.Println("i", i)
	return i
}

/*
	最后一个单词的长度
*/
func demo1_lengthOfLastWord_main() {
	// s := "Hello World"
	// s := "a "
	s := "b   a    "
	fmt.Println(lengthOfLastWord(s))
}

func lengthOfLastWord(s string) int {
	if len(s) == 0 {
		return 0
	}
	length := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' && length == 0 {
			continue
		} else if s[i] == ' ' {
			break
		} else {
			length++
		}
	}
	return length
}

/*
	反转字符串
编写一个函数，其作用是将输入的字符串反转过来。输入字符串以字符数组 char[] 的形式给出。

不要给另外的数组分配额外的空间，你必须原地修改输入数组、使用 O(1) 的额外空间解决这一问题。

你可以假设数组中的所有字符都是 ASCII 码表中的可打印字符。

*/

func demo1_reverseString_main() {
	var s = []byte("hello")
	reverseString(s)
}

func reverseString(s []byte) {
	if len(s) < 2 {
		fmt.Println(s)
	}
	var i int = len(s) - 1
	fmt.Println("当前的字符串：", s, len(s), i)
	var tmp byte
	for j := 0; j < len(s); j++ {
		if j > len(s)/2 {
			break
		}
		tmp = s[j]
		fmt.Println(s[j], s[i])
		fmt.Printf("第%d次处理后的%s：", j, s)
		s[j] = s[i]
		s[i] = tmp
		i--
	}
	fmt.Println("处理之后的字符串：", string(s))
}
