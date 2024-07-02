package tree

type binaryNode[T any] struct {
	data  *T
	left  *binaryNode[T]
	right *binaryNode[T]
	size  int
}

type binaryTree[T any] struct {
	root *binaryNode[T]
	size int
}

func NewBinaryTree(v int) *binaryTree[int] {
	return &binaryTree[int]{&binaryNode[int]{&v, nil, nil, 0}, 1}
}

func NewBinaryNode(v int, left, right *binaryNode[int]) *binaryNode[int] {
	size := 1
	if left != nil {
		size++
	}
	if right != nil {
		size++
	}
	return &binaryNode[int]{&v, left, right, size}
}

func (t *binaryTree[T]) AddLeftLeaf(leaf *binaryNode[T]) *binaryTree[T] {
	t.size++
	t.size += leaf.size
	t.root.left = leaf
	return t
}

func (t *binaryTree[T]) AddRightLeaf(leaf *binaryNode[T]) *binaryTree[T] {
	t.size++
	t.size += leaf.size
	t.root.right = leaf
	return t
}

func (t *binaryTree[T]) preOrderTraverse(node *binaryNode[T], leafsData *[]T) {
	if node == nil {
		return
	}
	*leafsData = append(*leafsData, *node.data)
	t.preOrderTraverse(node.left, leafsData)
	t.preOrderTraverse(node.right, leafsData)
}

func (t *binaryTree[T]) PreOrderTraverse() (leafsData []T) {
	t.preOrderTraverse(t.root, &leafsData)
	return leafsData
}

func (t *binaryTree[T]) postOrderTraverse(node *binaryNode[T], leafsData *[]T) {
	if node == nil {
		return
	}
	t.postOrderTraverse(node.left, leafsData)
	t.postOrderTraverse(node.right, leafsData)
	*leafsData = append(*leafsData, *node.data)
}

func (t *binaryTree[T]) PostOrderTraverse() (leafsData []T) {
	t.postOrderTraverse(t.root, &leafsData)
	return leafsData
}

func (t *binaryTree[T]) inOrderTraverse(node *binaryNode[T], leafsData *[]T) {
	if node == nil {
		return
	}
	t.inOrderTraverse(node.left, leafsData)
	*leafsData = append(*leafsData, *node.data)
	t.inOrderTraverse(node.right, leafsData)
}

func (t *binaryTree[T]) InOrderTraverse() (leafsData []T) {
	t.inOrderTraverse(t.root, &leafsData)
	return leafsData
}

/*

Upper are DFS traversal (Depth first search)
* We go from root to the deepest nested children, and then go backwards, through siblings (if they exist ofc), to the parent
* And We do it for every branch

*/

/*

BFS (Breadth-first search) according to the DFS is an opposite

* We dont go to the deepest children,
instead we go to the first children,
then all of its siblings,
then to the second children and all of its siblings,
then to children of children and all of its siblings etc etc

* It is more like "row" based

*/

type qNode struct {
	data *binaryNode[int]
	next *qNode
}

type queue struct {
	tail *qNode
	head *qNode
	len  int
}

func (q *queue) enqueue(data *binaryNode[int]) {
	node := &qNode{data, nil}
	if q.len == 0 {
		q.head = node
		q.tail = node
	} else {
		q.tail.next = node
		q.tail = node
	}
	q.len++
}

func (q *queue) dequeue() *binaryNode[int] {
	if q.len == 0 {
		return nil
	}
	node := q.head.data
	q.head = q.head.next
	q.len--
	return node
}

func NewQueue() *queue {
	return &queue{}
}

func bfsTraverse(queue *queue, traversePath *[]int) {
	n := queue.dequeue()
	for n != nil {
		*traversePath = append(*traversePath, *n.data)
		queue.enqueue(n.left)
		queue.enqueue(n.right)
		n = queue.dequeue()
	}
}

func BFSTraverse(tree *binaryTree[int]) (traversedValues []int) {
	q := NewQueue()
	q.enqueue(tree.root)
	bfsTraverse(q, &traversedValues)
	return traversedValues
}
