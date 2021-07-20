package binary_search_tree

// BinarySearchTree 二叉搜索树
type BinarySearchTree struct {
	size int
	root *Node
}

func CreateBst(values []int) *BinarySearchTree {
	tree := BinarySearchTree{
		size: 0,
		root: nil,
	}
	for _, value := range values {
		tree.Add(value)
	}
	return &tree
}

func (bst *BinarySearchTree) Root() *Node {
	return bst.root
}

func (bst *BinarySearchTree) Size() int {
	return bst.size
}

func (bst *BinarySearchTree) IsEmpty() bool {
	return bst.size == 0
}

func (bst *BinarySearchTree) Clear() {
	bst.root = nil
	bst.size = 0
}

func (bst *BinarySearchTree) Height() int {
	//return recursionH(bst.root)
	return height(bst.root)
}

// 通过递归的方式计算高度
func recursionH(node *Node) int {
	if node == nil {
		return 0
	}
	if recursionH(node.left) >= recursionH(node.right) {
		return 1 + recursionH(node.left)
	} else {
		return 1 + recursionH(node.right)
	}
}

// 普通方式计算高度
func height(node *Node) int {
	leftH := 1
	rightH := 1

	queue := []*Node{node}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		if cur.left != nil {
			leftH++
			queue = append(queue, cur.left)
		}

		if cur.right != nil {
			rightH++
			queue = append(queue, cur.right)
		}
	}

	if leftH >= rightH {
		return leftH
	}

	return rightH
}

// PreorderTravel 前序遍历 先父节点 后左节点 再右节点
func (bst *BinarySearchTree) PreorderTravel(callback func(value int)) {
	if bst.IsEmpty() {
		return
	}

	node := bst.root
	preorderTravel(node, callback)
}

func preorderTravel(node *Node, callback func(value int)) {
	if node == nil {
		return
	}

	callback(node.value)
	preorderTravel(node.left, callback)
	preorderTravel(node.right, callback)
}

// InorderTravel 中序遍历 先左节点 后父节点 再右节点
func (bst *BinarySearchTree) InorderTravel(callback func(value int)) {
	if bst.IsEmpty() {
		return
	}

	node := bst.root
	inorderTravel(node, callback)
}

func inorderTravel(node *Node, callback func(value int)) {
	if node == nil {
		return
	}

	inorderTravel(node.left, callback)
	callback(node.value)
	inorderTravel(node.right, callback)
}

// PostorderTravel 后序遍历 先左节点 后右节点 再父节点
func (bst *BinarySearchTree) PostorderTravel(callback func(value int)) {
	if bst.IsEmpty() {
		return
	}

	node := bst.root
	postorderTravel(node, callback)
}

func postorderTravel(node *Node, callback func(value int)) {
	if node == nil {
		return
	}

	postorderTravel(node.left, callback)
	postorderTravel(node.right, callback)
	callback(node.value)
}

// LevelTraversal 层序遍历 从root节点开始 从上往下 从左往右遍历节点
func (bst *BinarySearchTree) LevelTraversal(callback func(value int)) {
	if bst.IsEmpty() {
		return
	}

	queue := []*Node{bst.root}

	for len(queue) > 0 {
		// 取出第一个节点
		cur := queue[0]
		queue = queue[1:]

		// 执行遍历操作
		callback(cur.value)

		if cur.left != nil {
			queue = append(queue, cur.left)
		}

		if cur.right != nil {
			queue = append(queue, cur.right)
		}
	}
}

// IsComplete 判断这颗二叉树是否是完全二叉树
func (bst *BinarySearchTree) IsComplete() bool {
	if bst.IsEmpty() {
		return false
	}
	// 层序遍历
	queue := []*Node{bst.root}

	// 剩余节点是否都为子节点
	isLeaf := false
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		if cur.right != nil && cur.left == nil {
			return false
		}

		if isLeaf && !cur.IsLeaf() {
			return false
		}

		if cur.left != nil {
			queue = append(queue, cur.left)
		} else {
			isLeaf = true
		}

		if cur.right != nil {
			queue = append(queue, cur.right)
		} else {
			isLeaf = true
		}
	}

	return true
}


func (bst *BinarySearchTree) Add(value int) {

	if bst.root == nil {
		bst.root = &Node{
			value: value,
		}
		bst.size++
		return
	}

	node := bst.root
	parent := bst.root
	// 记录方向 true 是左边 false 是右边
	var dir bool
	// Todo: 支持泛型后替换为更通用的实现
	for node != nil {
		parent = node
		if node.value > value { // node.value 比较大 朝左边找
			node = node.left
			dir = true
		} else if node.value < value { // value 比较大 朝右边找
			node = node.right
			dir = false
		} else { // 相等 替换
			node.value = value
			return
		}
	}

	// 把当前value放到合适位置
	if dir { // 放左边
		parent.left = &Node{
			value:  value,
			parent: parent,
		}
	} else {
		parent.right = &Node{
			value:  value,
			parent: parent,
		}
	}

	bst.size++
}

func (bst *BinarySearchTree) remove(n *Node) {
	if bst.size == 1 { // 只有一个根节点
		bst.root = nil
		bst.size = 0
	}

	// 度为2的情况
	if n.HasTwoChild() { // 用前驱或者后置替换都可以
		// 找到前驱并修改value
		precursor := bst.precursor(n)
		n.value = precursor.value
		// 设置需要删除的节点
		n = precursor
	}

	// 删除度为1的节点
	if !n.IsLeaf() {
		var child *Node
		if n.left != nil {
			child = n.left
		} else {
			child = n.right
		}
		// 改变指向
		if n == n.parent.left {
			n.parent.left = child
		} else {
			n.parent.right = child
		}
		child.parent = n.parent
		return
	}

	// 最后一定是度为0的叶子结点
	if n == n.parent.left {
		n.parent.left = nil
	} else {
		n.parent.right = nil
	}
}

// 获取一个节点的前驱节点
func (bst *BinarySearchTree) precursor(n *Node) *Node {
	// 如果左子节点不为空 则 left->right->right....
	if n.left != nil {
		node := n.left
		for node.right != nil {
			node = node.right
		}

		return node
	}

	// 如果左子节点为空并且parent不为空 则 parent->parent->parent... 直到第一次出现在右子节点
	if n.left == nil && n.parent != nil {
		pNode := n.parent
		for pNode != nil {
			if n == pNode.right {
				return pNode
			}

			if n == pNode.left { // 继续往上找
				n = pNode
				pNode = pNode.parent
			}
		}
	}

	return nil
}

// 获取一个节点的后继节点
func (bst *BinarySearchTree) successor(n *Node) *Node {
	// 实现方法镜像前驱
	if n.right != nil {
		node := n.right
		for node.left != nil {
			node = node.left
		}

		return node
	}

	if n.right == nil && n.parent != nil {
		pNode := n.parent
		for pNode != nil {
			if n == pNode.left {
				return pNode
			}

			if n == pNode.right {
				n = pNode
				pNode = pNode.parent
			}
		}
	}

	return nil
}

// 通过value获取node
func (bst *BinarySearchTree) node(value int) *Node {
	node := bst.root
	for node != nil {
		if node.value > value { // 朝左找
			node = node.left
		} else if node.value < value { // 朝右找
			node = node.right
		} else { // 相等
			return node
		}
	}

	return nil
}

func (bst *BinarySearchTree) Remove(value int) {
	bst.remove(bst.node(value))
}

func (bst *BinarySearchTree) Contains(value int) bool {

	flag := false
	bst.LevelTraversal(func(e int) {
		if e == value {
			flag = true
		}
	})

	return flag
}





































