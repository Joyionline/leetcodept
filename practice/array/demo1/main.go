package main

import (
	"fmt"
)

func main() {
	// del_demo1_main()
	// removeElement_demo1_main()
	moveZeroes_demo1_main()
	// pivotIndex_demo1_main()/
	// dominantIndex_demo1_main()
}

// 从排序数组中删除重复项
// TODO 待修改 （可达到效果但不满足题意）

func del_demo1_main() {
	nums := []int{1, 1, 2}
	// fmt.Println(DelDuplicates_demo1(nums))
	fmt.Println(DelDuplicates_demo2(nums))
}
func DelDuplicates_demo1(nums []int) int {
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

// 从排序数组中删除重复项
/*
	双指针法：
		通过定义一个变量在一个循环中不断增加值，可以控制时间复杂度

复杂度分析
	时间复杂度：O(n)， 假设数组的长度是 n，那么 i 和 j 分别最多遍历 n 步。
	空间复杂度：O(1)。

*/

// TODO 返回值
func DelDuplicates_demo2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var i int = 0
	for j := 0; j < len(nums); j++ {
		fmt.Println("当前nums:", nums)
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}

// 移除元素
func removeElement_demo1_main() {
	nums := []int{3, 2, 2, 3}
	var val int = 3
	fmt.Println(removeElement_demo1(nums, val))
}

/*
	题目：
	给定一个数组 nums 和一个值 val，你需要原地移除所有数值等于 val 的元素，返回移除后数组的新长度。
	不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
	元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。
	示例 1:
	给定 nums = [3,2,2,3], val = 3,
	函数应该返回新的长度 2, 并且 nums 中的前两个元素均为 2。
	你不需要考虑数组中超出新长度后面的元素。
*/

func removeElement_demo1(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	var i int = 0
	for j := 0; j < len(nums); j++ {
		fmt.Println("当前nums:", nums)
		if nums[j] != val {
			nums[i] = nums[j]
			i++
		}
	}
	return i
}

/*
	移动零
给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
示例:
输入: [0,1,0,3,12]
输出: [1,3,12,0,0]
说明:
必须在原数组上操作，不能拷贝额外的数组。
尽量减少操作次数
*/

func moveZeroes_demo1_main() {
	nums := []int{0, 1, 0, 3, 12}
	// moveZeroes_demo1(nums)
	moveZeroes_demo2(nums)
}

// 移动零  188 ms	9 MB
func moveZeroes_demo1(nums []int) {
	var i int = 0
	var tmp int = 0
	for j := 0; j < len(nums); j++ {
		if nums[j] != 0 {
			tmp = nums[i]
			nums[i] = nums[j]
			nums[j] = tmp
			i++
		}
	}
	fmt.Println("输出数组：", nums)
}

// 移动零 196 ms	12.3 MB
func moveZeroes_demo2(nums []int) {
	var i int
	for j := 0; j < len(nums); j++ {
		if nums[j] != 0 {
			if i != j {
				nums[i] = nums[j]
				nums[j] = 0
			}
			i++
		}
	}
	fmt.Println("输出的数组：", nums)
}

// 寻找数组的中心索引
func pivotIndex_demo1_main() {
	// 测试用例  {1, 7, 3, 6, 5, 6}  {1, 2, 3}
	nums := []int{-1, -1, -1, 0, 1, 1}
	fmt.Println("数组nums的中心索引是：", pivotIndex_demo1(nums))
}

func pivotIndex_demo1(nums []int) int {
	if len(nums) <= 2 {
		return -1
	}
	var leftsum int
	for i := 0; i < len(nums); i++ {
		var rightsum int
		for j := i + 1; j < len(nums); j++ {
			rightsum = rightsum + nums[j]
		}
		if i != 0 {
			leftsum = leftsum + nums[i-1]
		}
		fmt.Println("当前左边和是:", leftsum)
		fmt.Println("当前右边和是:", rightsum)
		if leftsum == rightsum {
			fmt.Println("当前索引：", i)
			return i
			break
		} else {
			continue
		}
	}
	return -1
}

func dominantIndex_demo1_main() {
	nums := []int{3, 6, 1, 0}
	dominantIndex_demo1_1(nums)
}

// 寻找数组中至少是其它数字两倍的最大数
func dominantIndex_demo1_1(nums []int) int {
	var max, min int
	for j := 0; j < len(nums); j++ {
		fmt.Println("增加后的数据：", j)
		for k := 1; k < len(nums); k++ {
			if nums[k] > nums[j] {
				max = nums[k]
				min = nums[j]
			}
			if min < nums[j] || max < nums[k] {
				min = nums[j]
				max = nums[k]
			}
		}
	}
	fmt.Println("最大值和最小值分别是：", min, max)
	return 1
}
