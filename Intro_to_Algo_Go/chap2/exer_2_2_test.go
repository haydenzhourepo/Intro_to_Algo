package chap2

import (
	"fmt"
	"intro_to_algo_go/util"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	arr := []int{4, 2, 6, -2, -6, 0, 9, 55}
	SelectionSort(arr)
	actual_res := []int{-6, -2, 0, 2, 4, 6, 9, 55}

	fmt.Println(arr)
	if !util.SliceEqual(arr, actual_res) {
		t.Fail()
	}

	fmt.Println(arr)
}
