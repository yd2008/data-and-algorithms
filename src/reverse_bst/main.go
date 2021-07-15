package main

import (
	"example.com/m/v2/src/data/binary_search_tree"
	"fmt"
)

// 反转一颗二叉树
func main() {
	bst := binary_search_tree.CreateBst([]int{10, 5, 20, 1, 6, 11, 22})
	reverseBst(bst)
	bst.LevelTraversal(func(value int) {
		fmt.Println(value)
	})
}


func reverseBst(bst *binary_search_tree.BinarySearchTree) {
	reverse(bst.Root())
}

func reverse(node *binary_search_tree.Node)  {
	if node.Left() == nil && node.Right() == nil {
		return
	}

	// 用前序遍历实现
	node.SWAP()
	reverse(node.Left())
	reverse(node.Right())
}


