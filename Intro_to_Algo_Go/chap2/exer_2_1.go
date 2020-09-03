package chap2

// 2.1-2 insertion sort reverse
func InsertionSortReverse(arr []int) {
	for j := 1; j < len(arr); j++ {
		key := arr[j]

		// insert tem into sorted sequnce *arr[0, i-1]
		i := j - 1
		for i >= 0 && arr[i] < key {
			arr[i+1] = arr[i]
			i--
		}

		//  insert key if move
		if i != j-1 {
			arr[i+1] = key
		}
	}
}

// 2.1-3 search arr linearity by value v
func LinearSearch(arr []int, v int) int {
	for i, val := range arr {
		if val == v {
			return i
		}
	}

	return -1
}

// 2.1-4 sum two equal length binary integer array
func SumBinaryInteger(a, b []int) []int {
	if len(a) != len(b) {
		return nil
	}

	// carry flag
	CF := 0
	res := []int{}

	for i := len(a) - 1; i >= 0; i-- {
		sum := a[i] + b[i] + CF
		if sum >= 2 {
			res = append(res, sum-2)
			CF = 1
			continue
		} else {
			res = append(res, sum)
			CF = 0
		}
	}

	if CF == 1 {
		res = append(res, 1)
	}

	// reverse slice
	tem := []int{}
	for j := len(res) - 1; j >= 0; j-- {
		tem = append(tem, res[j])
	}

	return tem
}
