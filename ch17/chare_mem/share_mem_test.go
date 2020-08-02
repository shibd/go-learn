package chare_mem

import (
	"sync"
	"testing"
	"time"
)

func TestCounter(t *testing.T) {

	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			counter++
		}()
	}

	time.Sleep(time.Millisecond * 1000)
	t.Log(counter)

}

func TestCounterThreadSafe(t *testing.T) {

	var mut sync.Mutex
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			defer mut.Unlock()
			mut.Lock()
			counter++
		}()
	}

	time.Sleep(time.Millisecond * 1000)
	t.Logf("counter = %d", counter)
}

func TestCounterWaitGroup(t *testing.T) {

	var mut sync.Mutex
	var wg sync.WaitGroup
	counter := 0

	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				mut.Unlock()
				wg.Done()
			}()
			mut.Lock()
			counter++

		}()
	}
	wg.Wait()
	t.Logf("counter = %d", counter)

}
