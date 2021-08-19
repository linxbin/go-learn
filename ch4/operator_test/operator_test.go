package operator_test

import "testing"

func TestOperator(t *testing.T) {
	const (
		Readable   = 1 << iota // 1
		Writable               // 2
		Executable             // 3
	)
}

func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4, 5}
	b := [...]int{1, 2, 3, 4, 5}
	c := [...]int{1, 2, 3, 4, 4}
	if a == b {
		t.Log("a b equal")
	} else {
		t.Log("a b not equal")
	}
	t.Log(a == c)
}

const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestBitClear(t *testing.T) {
	t.Logf("%b", Readable) // 0001
	t.Logf("%b", Writable)
	t.Logf("%b", Executable)
	a := 7
	t.Logf("%b", a&^Readable) // 0111 &^ 0001 = 0110 按位异或运算符
	t.Logf("%b", a&Readable)  // 0111 & 0001 = 0001 按位与运算符

	a = a &^ Readable
	a = a &^ Executable
	t.Logf("%b", a)

	//
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
}
