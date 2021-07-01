package main

import (
	"example.com/m/v2/src/utils"
	"fmt"
)

func main() {
	array := utils.RandomIntArray(300, 11)
	fmt.Println(array)
	sort1 := Sort1(array)
	fmt.Println(sort1)
}

func Sort1(randomArr []int) []int {
	for end := len(randomArr) - 1; end > 0; end-- {
		for begin := 1; begin <= end; begin++ {
			//fmt.Printf("end is %d, begin is %d\n", end, begin)
			if randomArr[begin] < randomArr[begin-1] {
				randomArr[begin], randomArr[begin-1] = randomArr[begin-1], randomArr[begin]
			}
		}
	}
	return randomArr
}
