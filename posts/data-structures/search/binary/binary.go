package binary_search


func BinarySearch(arr []int, v int) bool {
	if len(arr) == 0 {
		return false
	}
	mid := len(arr)/2
	if arr[mid] == v {
		return true
	}
	left,right := arr[:mid],arr[mid+1:]
	if left[divide-1] > v {
		return BinarySearch(left,v)
	} else {
		return BinarySearch(right,v)
	}
}
