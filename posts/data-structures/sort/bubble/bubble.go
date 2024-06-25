package bubble_sort

func BubbleSort(arr []int) []int {
	count := len(arr)
	for {
		for i:= 0; i < count - 1; i++ {
			if arr[i] > arr[i+1] {
				arr[i],arr[i+1] = arr[i+1],arr[i]
			}
		}
	count--
	if count == 0 {
		break
	}
	}
	return arr
}

func BubbleSortWithLabel(arr []int) []int {
	count := 0
	iteration:
	for i:=0;i < len(arr)-1;i++ {
		if arr[i] > arr[i+1] {
			arr[i],arr[i+1] = arr[i+1],arr[i]
		}
	}
	count++
	if count == len(arr) -1 {
		return arr	
	}
	goto iteration
}




/*

Iterate over array and on each iteration we:

check current processed value array[i] > array[i+1]

	yes: swap array[i] with array[i+1] 
	no: omit this value, go to the next one

the result of each iteration is the greatest value is always pushed to the end,
then we splice array[:len(array)] and iterate again without the last value (we nknow its the greatest one)

instead of splicing im using while loop and the count variable which is array len at first init.
after each while iteration we decrement count, thus we base on index in iteration, we omit the last array value

*/
