package panic_recover

import (
	"errors"
	"fmt"
	"testing"
)

func TestPanicRecover(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recovered from ", err)
		}
	}()

	fmt.Println("Start")
	panic(errors.New("Something wrong!"))
}
