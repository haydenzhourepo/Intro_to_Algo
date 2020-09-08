package chap2

// merge sort
func MergeSort(arr []int, start, end int)  {
	if start < end {
		mid := (start + end) / 2
		MergeSort(arr, start, mid)
		MergeSort(arr, mid + 1, end)
		Merge(arr, start, mid, end)
	}
}

func Merge(arr []int, start, mid, end int)  {
	leftArr := make([]int, mid - start + 1)
	copy(leftArr, arr[start:mid + 1])

	rightArr := make([]int, end - mid)
	copy(rightArr, arr[(mid + 1):end + 1])

	i := 0
	j := 0
	k := start
	for {
		// left array is empty
		if i == len(leftArr) {
			if j < len(rightArr) {
				for _, i2 := range rightArr[j:] {
					arr[k] = i2
					k++
				}
			}
			break
		}
		// right array is empty
		if j == len(rightArr) {
			if i < len(leftArr) {
				for _, i2 := range leftArr[i:] {
					arr[k] = i2
					k++
				}
			}

			break
		}

		if leftArr[i] > rightArr[j] {
			arr[k] = rightArr[j]
			j++
		}else {
			arr[k] = leftArr[i]
			i++
		}
		k++
	}

}
