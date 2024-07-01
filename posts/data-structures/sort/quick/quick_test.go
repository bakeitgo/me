package quick

import (
	"fmt"
	"testing"
)

func TestSwappingInPartitionFunc(t *testing.T) {
	actual := []int{1, 15, 3, 12, 2, 3, 1, 10}
	expect := []int{1, 3, 2, 3, 1, 10, 15, 12}
	partition(actual)

	for i := 0; i < len(actual); i++ {
		if actual[i] != expect[i] {
			t.Fatalf("values havent been swapped correctly")
		}
	}
}

func TestPartitionReturnsPivotIndex(t *testing.T) {
	actual := []int{1, 15, 3, 12, 2, 3, 1, 10}
	// expect := []int{1, 3, 2, 3, 1, 10, 15, 12}
	pivotIdx := partition(actual)

	if pivotIdx.index() != 5 {
		t.Fatalf(fmt.Sprintf("Wrong pivot index returned, actual is: %d, should be: %d", pivotIdx, 5))
	}
}

type quicksorttest struct {
	actual   []int
	expected []int
}

var quicksort = []quicksorttest{
	{actual: []int{1, 15, 3, 12, 2, 3, 1, 10, 15, 15}, expected: []int{1, 1, 2, 3, 3, 10, 12, 15, 15, 15}},
	{actual: []int{1}, expected: []int{1}},
	{actual: []int{}, expected: []int{}},
	{actual: []int{10, 9, 8, 7}, expected: []int{7, 8, 9, 10}},
	{actual: []int{1, 1, 1, 1}, expected: []int{1, 1, 1, 1}},
	{actual: []int{1, 2, 2, 3, 1}, expected: []int{1, 1, 2, 2, 3}},
}

func TestQuickSort(t *testing.T) {
	for _, testcase := range quicksort {
		sorted := QuickSort(testcase.actual)
		for i := 0; i < len(testcase.actual); i++ {
			if sorted[i] != testcase.expected[i] {
				t.Fatalf("Array has not been sorted correctly, at index %d expected value: %d, actual value: %d", i, testcase.expected[i], sorted[i])
			}
		}
	}
}
