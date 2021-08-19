package type_test

import (
	"testing"
)

type MyInt int64

func TestType(t *testing.T) {
	var a int32 = 1
	var b int64
	var c MyInt
	c = MyInt(b)
	t.Log(a, b, c)
}

func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	t.Log(a, aPtr)
	t.Logf("%T, %T", a, aPtr)
}

func TestString(t *testing.T) {
	s := " "
	t.Log("*" + s + "*") //初始化零值是“”
	t.Log(len(s))

}
