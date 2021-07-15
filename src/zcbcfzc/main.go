package main

import "fmt"

// 最长不重复字串
// 例 "abcd" -> 4
// 例 "abcdabcd" -> 4
// 例 "bbbb" -> 1
// 例 "yes我爱golang你爱吗" -> 10
func main() {

	fmt.Println(maxDifSubStr("yes我爱golang你爱吗"))
	fmt.Println(maxDifSubStr("bbbbb"))
	fmt.Println(maxDifSubStr("abcd"))
}

// 返回长度最长的不重复子串 支持中文
func maxDifSubStr(str string) int {
	// 创建记录字典
	lastOccurIndexMap := make(map[rune]int)
	start := 0
	maxLength := 0

	for i, char := range []rune(str) {
		// 如果有相同字母出现并且index大于等于start则更新start
		if index, ok := lastOccurIndexMap[char]; ok && index >= start {
			start = index + 1
		}

		// 更新最长长度
		if i - start + 1 > maxLength {
			maxLength = i - start + 1
		}

		// 更新最后出现字典
		lastOccurIndexMap[char] = i
	}

	return maxLength
}