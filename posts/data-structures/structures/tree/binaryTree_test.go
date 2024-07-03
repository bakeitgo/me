package tree

import (
	"testing"
)

type binarytreetestCase struct {
	name     string
	actual   *binaryTree[int]
	expected []int
}

var binTree = NewBinaryTree(111)
var leftLeaf = NewBinaryNode(1, NewBinaryNode(10, NewBinaryNode(30, nil, nil), NewBinaryNode(40, nil, nil)), NewBinaryNode(20, NewBinaryNode(50, nil, nil), NewBinaryNode(60, nil, nil)))
var rightLeaf = NewBinaryNode(2, NewBinaryNode(12, NewBinaryNode(32, nil, nil), NewBinaryNode(42, nil, nil)), NewBinaryNode(22, NewBinaryNode(52, nil, nil), NewBinaryNode(62, nil, nil)))

var binaryTreeTraverseCases = []binarytreetestCase{
	{
		name:   "preorder",
		actual: binTree.AddLeftLeaf(leftLeaf).AddRightLeaf(rightLeaf),
		/*
						     111
					   /           \
					 1               2
				   /    \         /     \
			     10      20      12      22
			   /    \   /  \    /  \    /  \
			 30     40 50   60 32   42 52   62
		*/
		expected: []int{111, 1, 10, 30, 40, 20, 50, 60, 2, 12, 32, 42, 22, 52, 62},
	},
	{
		name:   "postorder",
		actual: binTree.AddLeftLeaf(leftLeaf).AddRightLeaf(rightLeaf),
		/*
						     111
					   /           \
					 1               2
				   /    \         /     \
			     10      20      12      22
			   /    \   /  \    /  \    /  \
			 30     40 50   60 32   42 52   62
		*/
		expected: []int{30, 40, 10, 50, 60, 20, 1, 32, 42, 12, 52, 62, 22, 2, 111},
	},
	{
		name:   "inorder",
		actual: binTree.AddLeftLeaf(leftLeaf).AddRightLeaf(rightLeaf),
		/*
						     111
					   /           \
					 1               2
				   /    \         /     \
			     10      20      12      22
			   /    \   /  \    /  \    /  \
			 30     40 50   60 32   42 52   62
		*/
		expected: []int{30, 10, 40, 1, 50, 20, 60, 111, 32, 12, 42, 2, 52, 22, 62},
	},
	{
		name:   "bfs",
		actual: binTree.AddLeftLeaf(leftLeaf).AddRightLeaf(rightLeaf),
		/*
						     111
					   /           \
					 1               2
				   /    \         /     \
			     10      20      12      22
			   /    \   /  \    /  \    /  \
			 30     40 50   60 32   42 52   62
		*/
		expected: []int{111, 1, 2, 10, 20, 12, 22, 30, 40, 50, 60, 32, 42, 52, 62},
	},
}

func TestPostOrderBinaryTreeTraverse(t *testing.T) {
	treeTraverse := binaryTreeTraverseCases[1]
	treeValues := treeTraverse.actual.PostOrderTraverse()
	if actual, expected := len(treeValues), len(treeTraverse.expected); actual != expected {
		t.Fatalf("Invalid length of data retrieved during tree traverse, actual len: %d, expected len: %d", actual, expected)
	}
	for i, expected := range treeTraverse.expected {
		if expected != treeValues[i] {
			t.Fatalf("Missing data after traversal: actual: %d, expected: %d, at index: %d", treeValues[i], expected, i)
		}
	}
}

func TestPreOrderBinaryTreeTraverse(t *testing.T) {
	treeTraverse := binaryTreeTraverseCases[0]
	treeValues := treeTraverse.actual.PreOrderTraverse()
	if actual, expected := len(treeValues), len(treeTraverse.expected); actual != expected {
		t.Fatalf("Invalid length of data retrieved during tree traverse, actual len: %d, expected len: %d", actual, expected)
	}
	for i, expected := range treeTraverse.expected {
		if expected != treeValues[i] {
			t.Fatalf("Missing data after traversal: actual: %d, expected: %d, at index: %d", treeValues[i], expected, i)
		}
	}
}

func TestInOrderBinaryTreeTraverse(t *testing.T) {
	treeTraverse := binaryTreeTraverseCases[2]
	treeValues := treeTraverse.actual.InOrderTraverse()
	if actual, expected := len(treeValues), len(treeTraverse.expected); actual != expected {
		t.Fatalf("Invalid length of data retrieved during tree traverse, actual len: %d, expected len: %d", actual, expected)
	}
	for i, expected := range treeTraverse.expected {
		if expected != treeValues[i] {
			t.Fatalf("Missing data after traversal: actual: %d, expected: %d, at index: %d", treeValues[i], expected, i)
		}
	}
}

func TestBFSBinaryTreeTraverse(t *testing.T) {
	treeTraverse := binaryTreeTraverseCases[3]
	treeValues := BFSTraverse(treeTraverse.actual)
	if actual, expected := len(treeValues), len(treeTraverse.expected); actual != expected {
		t.Fatalf("Invalid length of data retrieved during tree traverse, actual len: %d, expected len: %d", actual, expected)
	}
	for i, expected := range treeTraverse.expected {
		if expected != treeValues[i] {
			t.Fatalf("Missing data after traversal: actual: %d, expected: %d, at index: %d", treeValues[i], expected, i)
		}
	}
}

var binaryTreeComparison = []struct {
	leftTree  *binaryTree[int]
	rightTree *binaryTree[int]
	expected  bool
}{{
	leftTree:  NewBinaryTree(111).AddLeftLeaf(leftLeaf).AddRightLeaf(rightLeaf),
	rightTree: NewBinaryTree(111).AddLeftLeaf(leftLeaf).AddRightLeaf(rightLeaf),
	/*
						 111
				   /           \
				 1               2
			   /    \         /     \
			 10      20      12      22
		   /    \   /  \    /  \    /  \
		 30     40 50   60 32   42 52   62
	*/
	expected: true,
}, {
	leftTree:  NewBinaryTree(111).AddLeftLeaf(leftLeaf).AddRightLeaf(rightLeaf),
	rightTree: NewBinaryTree(111).AddRightLeaf(leftLeaf),
	/*
						 111
				   /           \
				 1               2
			   /    \         /     \
			 10      20      12      22
		   /    \   /  \    /  \    /  \
		 30     40 50   60 32   42 52   62
	*/
	expected: false,
}}

func TestBinaryTreeComparison(t *testing.T) {
	for _, tc := range binaryTreeComparison {
		equal := Compare(tc.leftTree, tc.rightTree)
		if equal != tc.expected {
			t.Fatalf("Tree are not equal")
		}
	}
}

var binaryTreeFind = []struct {
	tree      *binaryTree[int]
	findValue int
	expected  bool
}{{
	tree:      NewBinaryTree(111).AddLeftLeaf(leftLeaf).AddRightLeaf(rightLeaf),
	findValue: 111,
	/*
						 111
				   /           \
				 1               2
			   /    \         /     \
			 10      20      12      22
		   /    \   /  \    /  \    /  \
		 30     40 50   60 32   42 52   62
	*/
	expected: true,
}, {
	tree:      NewBinaryTree(111).AddLeftLeaf(leftLeaf).AddRightLeaf(rightLeaf),
	findValue: 15,
	/*
						 111
				   /           \
				 1               2
			   /    \         /     \
			 10      20      12      22
		   /    \   /  \    /  \    /  \
		 30     40 50   60 32   42 52   62
	*/
	expected: false,
},
	{
		tree:      NewBinaryTree(111).AddLeftLeaf(leftLeaf).AddRightLeaf(rightLeaf),
		findValue: 52,
		/*
							 111
					   /           \
					 1               2
				   /    \         /     \
				 10      20      12      22
			   /    \   /  \    /  \    /  \
			 30     40 50   60 32   42 52   62
		*/
		expected: true,
	},

	{
		tree:      NewBinaryTree(111).AddLeftLeaf(leftLeaf).AddRightLeaf(rightLeaf),
		findValue: 40,
		/*
							 111
					   /           \
					 1               2
				   /    \         /     \
				 10      20      12      22
			   /    \   /  \    /  \    /  \
			 30     40 50   60 32   42 52   62
		*/
		expected: true,
	},
}

func TestBinaryTreeFindValue(t *testing.T) {
	for idx, tc := range binaryTreeFind {
		found := Find(tc.tree, tc.findValue)
		if found != tc.expected {
			t.Fatalf("At test case %d, expected find result: %v, actual: %v", idx, tc.expected, found)
		}
	}
}
