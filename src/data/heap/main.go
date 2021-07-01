package heap

type Heap interface {
	IsEmpty() bool
	Clear()
	Add(e int)
	Get() int    // 获得堆顶元素
	Remove() int // 删除堆顶元素
	Replace(e int) int
}

type MyHeap struct {
	Elements []int
	Size     int
}

func (m *MyHeap) IsEmpty() bool {
	return m.Size == 0
}

func (m *MyHeap) Clear() {
	m.Elements = []int{}
	m.Size = len(m.Elements)
}

func (m *MyHeap) Add(e int) {
	m.Elements = append(m.Elements, e)
	m.Size = len(m.Elements)
	m.siftUp(m.Size - 1)
}

func (m *MyHeap) siftUp(index int) {
	// index = 0 就是根节点
	if index == 0 {
		return
	}

	// 优化
	e := m.Elements[index]
	for index > 0 {
		// 获取父节点
		pindex := (index - 1) >> 1
		p := m.Elements[pindex]
		if e <= p {
			break
		}

		m.Elements[index] = p
		index = pindex
	}
	m.Elements[index] = e
}

func (m *MyHeap) Remove() int {
	if m.Size == 0 {
		panic("size is 0!")
	}
	root := m.Elements[0]
	m.Elements[0] = m.Elements[m.Size-1]
	m.Elements = m.Elements[:m.Size-1]
	m.Size = len(m.Elements)
	m.siftDown(0)
	return root
}

func (m *MyHeap) siftDown(index int) {
	// 循环退出的条件是当前index没有叶子结点
	// 完全二叉树第一个叶子结点序号是 size >> 1
	e := m.Elements[index]
	half := m.Size >> 1
	for index < half {
		// 先取左子节点
		cindex := index<<1 + 1
		child := m.Elements[cindex]

		// 获取右子节点
		rcindex := cindex + 1

		// 选出最大子节点 并且比子节点大
		if (m.Size > rcindex && m.Elements[rcindex] > child) {
			child = m.Elements[rcindex]
			cindex = rcindex
		}

		if child <= e {
			break
		}

		m.Elements[index] = child
		index = cindex
	}
	// 把e放到合适的位置
	m.Elements[index] = e
}

func (m *MyHeap) Get() int {
	return m.Elements[0]
}

func (m *MyHeap) Replace(e int) int {
	if m.Size == 0 {
		panic("size is 0!")
	}
	root := m.Elements[0]
	m.Elements[0] = e
	m.siftDown(0)
	return root
}

// Heapify 批量建堆方法
func (m *MyHeap) Heapify() {
	// 从最后一个非叶子结点 自下而上的下滤
	for i := m.Size >> 1 - 1; i >= 0; i-- {
		m.siftDown(i)
	}
}

//func main() {
	//heap := &MyHeap{}
	//heap.Add(5)
	//heap.Add(4)
	//heap.Add(3)
	//heap.Add(2)
	//heap.Add(1)
	//heap.Add(10)
	//heap.Add(8)
	//fmt.Println(heap.elements)
	//fmt.Println(heap.size)
	//heap.Remove()
	//fmt.Println(heap.elements)
	//heap.Replace(6)
	//fmt.Println(heap.elements)
	//heap.Remove()
	//fmt.Println(heap.elements)
	//fmt.Println(heap.size)

	//fmt.Println(heap.elements)
	//ints := []int{4, 23, 6, 55, 5, 65, 546, 2123, 64, 62}
	//heap := MyHeap{elements: ints, size: len(ints)}
	//heap.heapify()
	//fmt.Println(heap.elements)
//}
