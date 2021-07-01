package main

import (
	"example.com/m/v2/src/utils"
	"fmt"
)
import "example.com/m/v2/src/data/heap"

func main() {

	array := utils.RandomIntArray(300, 11)
	fmt.Println(array)
	heap := heap.MyHeap{
		Elements: array,
		Size:     len(array),
	}
	heap.Heapify()
	for i := len(array)-1; i > 0; i-- {
		array[i] = heap.Remove()
	}
	fmt.Println(array)
}
