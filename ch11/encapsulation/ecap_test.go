package encapsulation

import (
	"fmt"
	"testing"
	"unsafe"
)

// 定义一个结构体
type Employee struct {
	Id   string
	Name string
	Age  int
}

func (e Employee) String() string {
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
}

func TestCreateEmployeeObj(t *testing.T) {
	e := Employee{"0", "Bob", 20}
	t.Log(e)
	e1 := Employee{Name: "Mike", Age: 30}
	t.Log(e1)
	t.Log(e.Id)
	e2 := new(Employee) //返回指针
	t.Log(e2)
	t.Logf("e is %T", e)
	t.Logf("e2 is %T", e2)
}

func TestStructOperations(t *testing.T) {
	e := Employee{"0", "Bob", 20}
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
	t.Log(e.String())
}
