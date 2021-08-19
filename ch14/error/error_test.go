package error

import (
	"errors"
	"fmt"
	"testing"
)

var LessThanTwoError = errors.New("n should be not less than 2")

var LargerThenHundredError = errors.New("n should be not larger than 100")

func GetFibonacci(n int) ([]int, error) {
	if n < 2 {
		return nil, LessThanTwoError
	}
	if n > 100 {
		return nil, LargerThenHundredError
	}
	fibList := []int{1, 1}
	for i := 2; /*短变量声明 := */ i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil
}

func TestGetFibonacci(t *testing.T) {
	if v, err := GetFibonacci(4); err != nil {
		if err == LessThanTwoError {
			fmt.Println("It is less.")
		}
		if err == LargerThenHundredError {
			fmt.Println("It is larger.")
		}
		t.Error(err)
	} else {
		t.Log(v)
	}
}
