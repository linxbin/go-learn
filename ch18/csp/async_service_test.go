package concurrency

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond * 50)
	return "Done"
}

func otherTask() {
	fmt.Println("working on something else")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("Task is done.")
}

func TestService(t *testing.T) {
	fmt.Println(service())
	otherTask()
}

func AsyncService(wg *sync.WaitGroup) chan string {
	retCh := make(chan string, 2)
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func(i int) {
			ret := service()
			fmt.Printf("returned result as  %d. \n", i)
			retCh <- ret
			fmt.Println("service exited.")
			wg.Done()
		}(i)
	}

	return retCh
}

//
func TestAsyncService(t *testing.T) {
	var wg sync.WaitGroup
	retCh := AsyncService(&wg)
	otherTask()
	for i := 0; i < 2; i++ {
		fmt.Println(<-retCh)
	}
	wg.Wait()
}
