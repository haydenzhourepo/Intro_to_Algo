package chap2

import (
	"fmt"
	"intro_to_algo_go/util"
	"testing"
)

func TestInsertionSortReverse(t *testing.T) {
	arr := []int{4, 2, 6, -2, -6, 0, 9, 55}
	InsertionSortReverse(arr)
	actual_res := []int{55, 9, 6, 4, 2, 0, -2, -6}

	fmt.Println(arr)
	if !util.SliceEqual(arr, actual_res) {
		t.Fail()
	}

	fmt.Println(arr)
}

func TestLinearSearch(t *testing.T) {
	arr := []int{4, 2, 6, -2, -6, 0, 9, 55}

	if LinearSearch(arr, 1) != -1 {
		t.Fail()
	}

	if LinearSearch(arr, 2) != 1 {
		t.Fail()
	}
}

func TestSumBinaryInteger(t *testing.T) {
	a := []int{1, 0, 1, 0, 1}
	b := []int{1, 1, 0, 1, 0}

	res := SumBinaryInteger(a, b)
	actual_res := []int{1, 0, 1, 1, 1, 1}

	fmt.Println(res)
	if !util.SliceEqual(res, actual_res) {
		t.Fail()
	}
}
