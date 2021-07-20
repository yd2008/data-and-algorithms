package main

import (
	"example.com/m/v2/src/utils"
	"fmt"
)

// lis longest increasing subsequence (lis)
// 最长上升子序列（可以不连续）
// 10,2,2,5,1,7,101,18
func main() {
	fmt.Println(lis([]int{10, 2, 2, 5, 1, 7, 101, 18}))
}

func lis(nums []int) int {
	// 定义状态 所有dp[x]的初始值为1
	// dp[i] 就是以第nums[i]结束的最长子序列
	dp := make([]int, len(nums))
	dp[0] = 1
	maxLength := 0
	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ { // 遍历之前遍历过的数组
			if nums[j] >= nums[i] {
				continue
			}
			dp[i] = utils.MaxInt(dp[i], dp[j]+1)
		}
		maxLength = utils.MaxInt(maxLength, dp[i])
	}
	fmt.Println(dp)
	return maxLength
}
