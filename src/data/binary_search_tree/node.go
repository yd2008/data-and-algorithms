package binary_search_tree

type Node struct {
	value int
	left *Node
	right *Node
	parent *Node
}

func (n *Node) Left() *Node {
	return n.left
}

func (n *Node) Right() *Node {
	return n.right
}

func (n *Node) IsLeaf() bool {
	return n.left == nil && n.right == nil
}

func (n *Node) HasTwoChild() bool {
	return n.left != nil && n.right != nil
}

// SWAP 交换节点的左右子节点
func (n *Node) SWAP() {
	n.left, n.right = n.right, n.left
}