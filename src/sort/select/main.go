package main

import (
	"example.com/m/v2/src/utils"
	"fmt"
)

func main() {
	array := utils.RandomIntArray(300, 25)
	fmt.Println(array)
	sort1 := Sort1(array)
	fmt.Println(sort1)
}

func Sort1(randomArr []int) []int {
	for end := len(randomArr) - 1; end > 0; end-- {
		maxIndex := 0
		for begin := 1; begin <= end; begin++ {
			if randomArr[maxIndex] < randomArr[begin] {
				maxIndex = begin
			}
		}
		randomArr[maxIndex], randomArr[end] = randomArr[end], randomArr[maxIndex]
	}
	return randomArr
}
