package main

import (
	"example.com/m/v2/src/data/binary_search_tree"
	"fmt"
)

func main() {

	//bst := binary_search_tree.CreateBst([]int{4,11,2,5,6,9,8,10,3})
	completeBst := binary_search_tree.CreateBst([]int{10,5,20,1,6,11,22})
	//bst := binary_search_tree.CreateBst([]int{2,1,3})

	//bst.LevelTraversal(func(value int) {
	//	fmt.Println(value)
	//})

	//fmt.Println(bst.Contains(4))

	//fmt.Println(bst.Height())
	//fmt.Println(bst.IsComplete())
	//fmt.Println(completeBst.IsComplete())
	//fmt.Println(completeBst.TestPrecursor(11))
	completeBst.Remove(22)
	completeBst.InorderTravel(func(value int) {
		fmt.Println(value)
	})
}
