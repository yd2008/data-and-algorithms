package main

import (
	"example.com/m/v2/src/utils"
	"fmt"
)

func main() {

	randomArray := utils.RandomIntArray(400, 15)
	fmt.Println(randomArray)
	sortArr := Sort(randomArray)
	fmt.Println(sortArr)
}

var (
	array []int
)

func Sort(randomArr []int) []int {
	array = randomArr
	sort(0, len(array))
	return array
}

func sort(begin, end int) {
	if end-begin < 2 {
		return
	}

	point := pointIndex(begin, end)
	sort(begin, point)
	sort(point+1, end)
}

// pointIndex 获取当前轴点的位置
func pointIndex(begin, end int) int {
	// 备份轴点元素
	pointE := array[begin]

	// end 指向最后元素
	end--

	for begin < end {
		for begin < end {
			if pointE < array[end] {
				end--
			} else {
				array[begin] = array[end]
				begin++
				break
			}
		}

		for begin < end {
			if pointE > array[begin] {
				begin++
			} else {
				array[end] = array[begin]
				end--
				break
			}
		}
	}

	// 将轴点元素放到最终位置
	array[begin] = pointE

	return begin
}
