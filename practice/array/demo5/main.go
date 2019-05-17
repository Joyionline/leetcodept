package main

import (
	"fmt"
	"sort"
)

func main() {
	// demo5_singleNumber_main()
	demo5_maxProfit_main()
}

/*
	只出现一次的数字
	给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
说明：
你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？
示例 1:
输入: [2,2,1]
输出: 1
示例 2:
输入: [4,1,2,1,2]
输出: 4

*/

func demo5_singleNumber_main() {
	// 测试用例{2, 2, 1} ,{4, 1, 2, 1, 2}
	nums := []int{1, 2, 3, 3, 1}
	fmt.Println("只出现一次的数据：", singleNumber(nums))
}

func singleNumber(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	sort.Ints(nums)
	var lastnum int
	for i := 0; i < len(nums); i++ {
		if i != 0 && i != len(nums)-1 {
			if nums[i] != nums[i-1] && nums[i] != nums[i+1] {
				lastnum = nums[i]
			}
		}
		if i == 0 {
			lastnum = nums[i]
		}
		if i == len(nums)-1 && nums[i] != nums[i-1] {
			lastnum = nums[i]
		}
	}
	return lastnum
}

/*
	买卖股票的最佳时机 II
给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。
设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成更多的交易（多次买卖一支股票）。
注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。

示例 1:
输入: [7,1,5,3,6,4]
输出: 7
解释: 在第 2 天（股票价格 = 1）的时候买入，在第 3 天（股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4 。
     随后，在第 4 天（股票价格 = 3）的时候买入，在第 5 天（股票价格 = 6）的时候卖出, 这笔交易所能获得利润 = 6-3 = 3 。
示例 2:
输入: [1,2,3,4,5]
输出: 4
解释: 在第 1 天（股票价格 = 1）的时候买入，在第 5 天 （股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4 。
     注意你不能在第 1 天和第 2 天接连购买股票，之后再将它们卖出。
     因为这样属于同时参与了多笔交易，你必须在再次购买前出售掉之前的股票。
示例 3:
输入: [7,6,4,3,1]
输出: 0
解释: 在这种情况下, 没有交易完成, 所以最大利润为 0。

*/

func demo5_maxProfit_main() {
	// 测试用例  {7, 1, 5, 3, 6, 4}，{1, 2, 3, 4, 5},{7, 6, 4, 3, 1},{1},{2, 1, 2, 0, 1},{3, 3, 5, 0, 0, 3, 1, 4}
	nums := []int{5, 2, 3, 2, 6, 6, 2, 9, 1, 0, 7, 4, 5, 0}
	// fmt.Println("该次股票交易所获取的利润是：", maxProfit(nums))
	fmt.Println("该次股票交易所获取的利润是：", maxProfit_2_1(nums))
}

// TODO 暴力法
func maxProfit_2(prices []int) int {
	var get, in, out, incount, outcount int
	for i := 0; i < len(prices); i++ {
		if i != len(prices)-1 {
			if prices[i] <= prices[i+1] && incount == outcount {
				in = prices[i]
				incount++
			}
		}
		if i != 0 && i != len(prices)-1 {
			if prices[i] > prices[i-1] && prices[i+1] < prices[i] && incount != 0 {
				out = prices[i]
				get += out - in
				outcount++
			}
		}
		if i == len(prices)-1 && i != 0 {
			if prices[i] > prices[i-1] && incount != 0 {
				out = prices[i]
				get += out - in
				outcount++
			}
		}
		fmt.Println("当前买入的价格是：", in)
		fmt.Println("当前卖出的价格是：", out)
	}
	return get
}

/*
简单遍历
	该种方式采用了同一天买入和卖出，
	当今天的价格比昨天高时就卖出昨天的，买入今天的。
*/
func maxProfit_2_1(prices []int) int {
	var maxprofit, in, out int
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			out = prices[i]
			in = prices[i]
			maxprofit += prices[i] - prices[i-1]
		}
		fmt.Println("当前买入的价格是：", in)
		fmt.Println("当前卖出的价格是：", out)
	}
	return maxprofit
}
