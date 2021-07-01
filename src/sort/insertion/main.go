package main

import (
	"example.com/m/v2/src/utils"
	"fmt"
)

func main() {
	array := utils.RandomIntArray(400, 32)
	fmt.Println(array)
	sort := Sort2(array)
	fmt.Println(sort)
}

func Sort(randomArr []int) []int {
	for begin := 1; begin < len(randomArr); begin++ {
		for cur := begin; cur > 0; cur-- {
			if randomArr[cur] < randomArr[cur-1] {
				randomArr[cur-1], randomArr[cur] = randomArr[cur], randomArr[cur-1]
			}
		}
	}
	return randomArr
}

// Sort1 挪动实现
func Sort1(randomArr []int) []int {
	for begin := 1; begin < len(randomArr); begin++ {

	}
	return randomArr
}

// Sort2 添加二分插入
func Sort2(randomArr []int) []int {
	for begin := 1; begin < len(randomArr); begin++ {
		e := randomArr[begin]
		index := searchIndex(e, randomArr[:begin])
		// 将[index, begin) 朝右边挪
		for i := begin; i > index; i-- {
			randomArr[i] = randomArr[i-1]
		}
		randomArr[index] = e
	}
	return randomArr
}

// 二分搜索 有就返回 没有则返回-1
func search(e int) int {
	ints := []int{2, 5, 6, 7, 8, 9, 10, 18, 33}

	// 查找中间元素
	begin := 0
	end := len(ints)
	for begin < end {
		mid := (begin + end) >> 1
		if e < ints[mid] {
			end = mid
		} else if e > ints[mid] {
			begin = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

// 二分搜索确定取药插入的位置
func searchIndex(e int, ints []int) int {
	begin := 0
	end := len(ints)
	for begin < end {
		mid := (begin + end) >> 1
		if e < ints[mid] {
			end = mid
		} else {
			begin = mid + 1
		}
	}
	return begin
}