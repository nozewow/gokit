package main

import (
	"log"
	"sync"
	"sync/atomic"
)

// atomic load and store
func main() {
	var at atomic.Value
	var num int
	at.Store(num)
	var wg sync.WaitGroup
	wg.Add(2)
	var mu sync.Mutex

	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			mu.Lock()
			v := at.Load().(int)
			v++
			at.Store(v)
			mu.Unlock()
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			mu.Lock()
			v := at.Load().(int)
			v++
			at.Store(v)
			mu.Unlock()
		}
	}()

	wg.Wait()
	val := at.Load().(int)
	log.Println(val)
}
