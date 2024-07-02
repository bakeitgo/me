package tree

type node[T any] struct {
	value T
}

type tree[T any] struct {
	data  *node[T]
	leafs []*tree[T]
}

func NewTree[T any](data T) *tree[T] {
	return &tree[T]{&node[T]{data}, nil}
}

func (t *tree[T]) AddLeaf(leaf *tree[T]) *tree[T] {
	t.leafs = append(t.leafs, leaf)
	return t
}

func preOrderTraverse(tree *tree[int], leafsData *[]int) {
	if tree.data == nil {
		return
	}
	*leafsData = append(*leafsData, tree.data.value)
	for _, leaf := range tree.leafs {
		preOrderTraverse(leaf, leafsData)
	}
}

func PreOrderTraverse(tree *tree[int]) (leafsData []int) {
	preOrderTraverse(tree, &leafsData)
	return
}

func postOrderTraverse(tree *tree[int], leafsData *[]int) {
	if tree.data == nil {
		return
	}
	for _, leaf := range tree.leafs {
		postOrderTraverse(leaf, leafsData)
	}
	*leafsData = append(*leafsData, tree.data.value)
}

func PostOrderTraverse(tree *tree[int]) (leafsData []int) {
	postOrderTraverse(tree, &leafsData)
	return
}
