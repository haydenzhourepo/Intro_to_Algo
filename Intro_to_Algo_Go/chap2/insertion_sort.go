package chap2

// insertion sort
func InsertionSort(arr []int) {
	for j := 1; j < len(arr); j++ {
		key := arr[j]

		// insert tem into sorted sequnce *arr[0, i-1]
		i := j - 1
		for i >= 0 && arr[i] > key {
			arr[i+1] = arr[i]
			i--
		}

		//  insert key if move
		if i != j-1 {
			arr[i+1] = key
		}
	}
}
