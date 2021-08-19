package array_test

import "testing"

func TestArrayInt(t *testing.T) {
	var arr = [3]int{2, 3, 4}
	arr[2] = 4
	t.Log(arr)
}

func TestArrayTravel(t *testing.T) {
	arr3 := [...]int{1, 3, 4, 5}
	for i := 0; i < len(arr3); i++ {
		t.Log(arr3[i])
	}
	for _, e := range arr3 {
		t.Log(e)
	}
}

func TestArraySection(t *testing.T) {
	arr3 := [...]int{1, 2, 3, 4, 5}
	arr3Sec := arr3[2:]
	t.Log(arr3Sec)
}
