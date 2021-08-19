package empty_interface

import (
	"fmt"
	"testing"
)

func DoSomething(p interface{}) {
	switch v := p.(type) {
	case int:
		fmt.Println("Integer", v)
	case string:
		fmt.Println("String", v)
	default:

		fmt.Println("Unknown Type", v)
	}
}

func TestEmptyInterface(t *testing.T) {
	DoSomething(10)
	DoSomething("10")
	param := [...]int{1, 2, 3, 4}
	DoSomething(param)
}
