package main

import (
	"example.com/m/v2/src/utils"
	"fmt"
)

func main() {
	randomArray := utils.RandomIntArray(600, 32)
	sortedArr := Sort(randomArray)
	fmt.Println(sortedArr)
}

var (
	array []int
)

func Sort(randomArr []int) []int {
	array = randomArr
	shellArr := generateShellSteps(len(randomArr))
	for _, step := range shellArr {
		sort(step)
	}

	return array
}

func sort(step int) {
	for col := 0; col < step; col++ {
		for begin := col + step; begin < len(array); begin += step {
			// cur> col 的原因 col, col+step, col+2*step ...
			for cur := begin; cur > col; cur -= step {
				if array[cur] < array[cur-step] {
					array[cur], array[cur-step] = array[cur-step], array[cur]
				}
			}
		}
	}
}

func generateShellSteps(length int) []int {
	var arr []int
	for (length >> 1) > 0 {
		length >>= 1
		arr = append(arr, length)
	}
	return arr
}
