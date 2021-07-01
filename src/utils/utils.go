package utils

import "math/rand"

// RandomIntArray 返回一个最大数为 max， 长度length的随机int数组
func RandomIntArray(max, length int) []int {
	// 固定seed
	rand.Seed(666)
	slice := make([]int, length)
	for i := range slice {
		slice[i] = rand.Intn(max)
	}
	return slice
}