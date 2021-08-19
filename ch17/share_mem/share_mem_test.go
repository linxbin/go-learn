package share_mem

import (
	"sync"
	"testing"
	"time"
)

func TestCounter(t *testing.T) {
	counter := 0
	for i := 1; i < 10; i++ {
		go func() {
			counter++
		}()
	}
	time.Sleep(1 * time.Second)
	t.Logf("counter = %d", counter)
}

func TestCounterThreadSafe(t *testing.T) {
	var mut sync.Mutex
	counter := 0
	for i := 0; i < 10; i++ {
		go func() {
			defer mut.Unlock()
			mut.Lock()
			counter++
		}()
	}
	time.Sleep(time.Second * 1)
	t.Logf("counter = %d", counter)
}

func TestCounterWaitGroup(t *testing.T) {
	var mut sync.Mutex
	var waitGroup sync.WaitGroup
	counter := 0
	for i := 0; i < 10; i++ {
		waitGroup.Add(1)
		go func() {
			defer mut.Unlock()
			mut.Lock()
			counter++
			waitGroup.Done()
		}()
	}
	waitGroup.Wait()
	t.Logf("counter = %d", counter)
}
