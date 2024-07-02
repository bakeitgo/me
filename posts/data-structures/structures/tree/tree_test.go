package tree

import "testing"

type treetestCase struct {
	name     string
	actual   *tree[int]
	expected []int
}

var treeTraverseCases = []treetestCase{
	{
		name:   "preorder",
		actual: NewTree(1).AddLeaf(NewTree(3).AddLeaf(NewTree(10)).AddLeaf(NewTree(15))).AddLeaf(NewTree(5).AddLeaf(NewTree(3)).AddLeaf(NewTree(8)).AddLeaf(NewTree(99))),
		/*
					1
				  /   \
				3       5
			  /   \    / | \
			10    15  3  8  99
		*/
		expected: []int{1, 3, 10, 15, 5, 3, 8, 99},
	},
	{
		name:   "postorder",
		actual: NewTree(1).AddLeaf(NewTree(3).AddLeaf(NewTree(10)).AddLeaf(NewTree(15))).AddLeaf(NewTree(5).AddLeaf(NewTree(3)).AddLeaf(NewTree(8)).AddLeaf(NewTree(99))),
		/*
					1
				  /   \
				3       5
			  /   \    / | \
			10    15  3  8  99
		*/
		expected: []int{10, 15, 3, 3, 8, 99, 5, 1},
	},
}

func TestPreOrderTraverse(t *testing.T) {
	treeTraverse := treeTraverseCases[0]
	actual := PreOrderTraverse(treeTraverse.actual)
	if actual, expected := len(actual), len(treeTraverse.expected); actual != expected {
		t.Fatalf("Invalid length of data retrieved during tree traverse, actual len: %d, expected len: %d", actual, expected)
	}
	for i, expected := range treeTraverse.expected {
		if expected != actual[i] {
			t.Fatalf("Missing data after traversal: actual: %d, expected: %d, at index: %d", actual[i], expected, i)
		}
	}
}

func TestPostOrderTraverse(t *testing.T) {
	treeTraverse := treeTraverseCases[1]
	actual := PostOrderTraverse(treeTraverse.actual)
	if actual, expected := len(actual), len(treeTraverse.expected); actual != expected {
		t.Fatalf("Invalid length of data retrieved during tree traverse, actual len: %d, expected len: %d", actual, expected)
	}
	for i, expected := range treeTraverse.expected {
		if expected != actual[i] {
			t.Fatalf("Missing data after traversal: actual: %d, expected: %d, at index: %d", actual[i], expected, i)
		}
	}
}
