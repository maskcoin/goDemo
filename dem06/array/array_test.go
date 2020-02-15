package array

import "testing"

func TestArrayInit(t *testing.T) {
	var arr [3]int
	arr1 := [4]int{1, 2, 3, 4}
	arr3 := [...]int{1, 3, 4, 5, 8}
	arr1[1] = 5
	t.Log(arr[1], arr[2])
	t.Log(arr1, arr3)
}

func TestArrayTravel(t *testing.T) {
	arr := [...]int{1, 3, 4, 5, 8}
	for i := 0; i < len(arr); i++ {
		t.Log(arr[i])
	}
	for i, v := range arr {
		t.Log(i, v)
	}
}

func TestArraySection(t *testing.T) {
	arr := [...]int{1, 2, 3, 4, 5, 6}
	arrSec := arr[:3]
	t.Log(arrSec)
}
