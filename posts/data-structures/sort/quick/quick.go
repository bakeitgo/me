package quick

/*

## Algorithm lookup

	* Lets examine we have an array: []int{1,9,7,6,3,8,4,2,5}

	* We need to choose pivot [We will choose the last element of the array for the sake of simplicity (5)]

	* Then we need to circle around the pivot with rest of the values where:
		*  <= 5 will go to the left side,
		*  > 5 will go to the right side
	It will look like this: [---LESS THAN OR EQUAL <---- PIVOT ----> GREATER THAN -------]

	* After init "sort" where we put [-- LESS -- PIVOT -- GREATER --], we gonna PARTITION array onto 2 sub arrays where

		* 1st array will be [0 -- PIVOT INDEX - 1] (EXCLUSIVE)

		* 2st array [PIVOT INDEX + 1 -- part.len -1] (EXCLUSIVE)

		* And in every single array we gonna run operation again

	* When every array is partitioned already to place where its length is 0 // 1, it means its sorted already and we can return all arrays up to the parent


*/

/*
	swap

	* iterate over array [1,15,3,12,2,3,1,10]

	* pivot - 10

	* if v <= pivot swap

	* on swap incr idx

	* [1,15,] idx++ [0]

	* [1,3,15] idx++ [1] swap arr[i] arr[idx]

	* [1,3,2,15,12] idx++ [2] swap arr[i] arr[idx]


*/

// wrapping onto struct to encapsulate index and make it immutable
type pivotIndex struct {
	i int
}

func (p pivotIndex) index() int {
	return p.i
}

func partition(arr []int) pivotIndex {
	pivotIdx := pivotIndex{len(arr) - 1}
	swapIdx := -1

	for i := 0; i < len(arr); i++ {
		if arr[i] <= arr[pivotIdx.index()] {
			swapIdx++
			arr[i], arr[swapIdx] = arr[swapIdx], arr[i]
		}
	}

	return pivotIndex{swapIdx}
}

func qs(arr []int) {
	if len(arr) <= 1 {
		return
	}
	pivot := partition(arr)
	qs(arr[:pivot.index()])
	qs(arr[pivot.index():])
}

func QuickSort(arr []int) []int {
	qs(arr)
	return arr
}
