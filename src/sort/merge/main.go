package main

import (
	"example.com/m/v2/src/utils"
	"fmt"
)

func main() {
	randomArr := utils.RandomIntArray(68, 3)
	fmt.Println(randomArr)
	sortArr := Sort(randomArr)
	fmt.Println(sortArr)
}

var (
	array     []int
	leftArray []int
)

// Sort 归并排序
func Sort(randomArr []int) []int {
	array = randomArr
	leftArray = make([]int, len(array)>>1)
	sort(0, len(array))
	return array
}

// sort 对 [begin,end) 范围的数据进行归并排序
func sort(begin, end int) {
	if end-begin < 2 {
		return
	}

	mid := (begin + end) >> 1
	sort(begin, mid)
	sort(mid, end)
	merge(begin, mid, end)
}

func merge(begin, mid, end int) {
	li, le := 0, mid-begin
	ri, re := mid, end
	ai := begin

	// 备份左边数组
	for i := li; i < le; i++ {
		leftArray[i] = array[begin+i]
	}

	// 左边还没有完
	for li < le {
		if ri < re && array[ri] < leftArray[li] {
			array[ai] = array[ri]
			ri++
			ai++
		} else {
			array[ai] = leftArray[li]
			li++
			ai++
		}
	}
}
