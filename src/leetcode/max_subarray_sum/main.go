package main

import "fmt"

// 最大连续子序列和
// 例： -2,1,-3,4,-1,2,1,-5,4
// 子序列可以不连续
// 子串，子数组，子区间 必须连续
func main() {
	maxSum := maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
	fmt.Println(maxSum)
}

// 返回最长公共子序列的和
func maxSubArray(nums []int) int {
	// 定义dp[i]是 nums[i]为结尾的最大值
	// 如果 dp[i] < 0 则直接抛弃
	// 反之dp[i] 的最大值就是 dp[i - 1] + nums[i]
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	maxSum := 0
	for i := 1; i < len(nums); i++ {
		if dp[i-1] < 0 {
			dp[i] = nums[i]
		} else {
			dp[i] = dp[i-1] + nums[i]
		}
		if maxSum < dp[i] {
			maxSum = dp[i]
		}
	}
	fmt.Println(dp)
	return maxSum
}
