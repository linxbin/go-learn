package func_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

type MyFunc func(op int) int

func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

func TestFn(t *testing.T) {
	a, b := returnMultiValues()
	t.Log(a, b)

	tsSF := timeSpent(slowFun)
	t.Log(tsSF(10))
}

func timeSpent(inner MyFunc) MyFunc {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

func slowFun(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

func Sum(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret += op
	}
	return ret
}

func TestVarParams(t *testing.T) {
	t.Log(Sum(1, 2, 3, 4, 5, 6, 7))
	t.Log(Sum(1, 2, 3, 4))
}

func Clear() {
	fmt.Println("Clear resources.")
}

func TestDefer(t *testing.T) {
	defer Clear()
	t.Log("Start")
}
