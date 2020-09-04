package chap2

// 2.2-2 selection sort
func SelectionSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}

		// change if i != minIndex
		if i != minIndex {
			tem := arr[i]
			arr[i] = arr[minIndex]
			arr[minIndex] = tem
		}
	}
}
