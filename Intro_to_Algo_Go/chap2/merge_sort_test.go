package chap2

import (
	"fmt"
	"intro_to_algo_go/util"
	"testing"
)

func TestMergeSort(t *testing.T) {
	arr := []int{4, 2, 6, -2, -6, 0, 9, 55}
	MergeSort(arr, 0, len(arr) - 1)
	fmt.Println(arr)


	actual_res := []int{-6, -2, 0, 2, 4, 6, 9, 55}

	if !util.SliceEqual(arr, actual_res) {
		t.Fail()
	}
}

