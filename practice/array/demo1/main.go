package main

import "fmt"

func main() {
	nums := []int{1, 1, 2}
	fmt.Println(DelDuplicates(nums))
}

// 从排序数组中删除重复项
func DelDuplicates(nums []int) int {
	newnums := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		repeat := false
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newnums = append(newnums, nums[i])
		}
	}
	fmt.Println("新的数组数据：", newnums)
	return len(newnums)
}
