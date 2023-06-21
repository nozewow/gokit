package bench

import (
	"sync"
	"testing"
)

var (
	m     sync.Map
	times int = 100000
)

func initMap(num int) {
	for i := 0; i < num; i++ {
		m.Store(i, i)
	}
}

func loadMap(num int) {
	//for i := 0; i < num; i++ {
	//	m.Load(i)
	//}
	var wg sync.WaitGroup
	for i := 0; i < num; i++ {
		wg.Add(1)
		go func(n int) {
			m.Load(n * num)
			wg.Done()
		}(i)
	}

	wg.Wait()
}

func BenchmarkMap(b *testing.B) {
	initMap(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		loadMap(times)
	}
}

// BenchmarkLockMap-8   	     572	   2135636 ns/op
func BenchmarkLockMap(b *testing.B) {
	mm := make(map[int]int)
	var mu sync.RWMutex
	for q := 0; q < 10000; q++ {
		mm[q] = q
	}

	var wg sync.WaitGroup
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for q := 0; q < times; q++ {
			wg.Add(1)
			go func(num int) {
				mu.RLock()
				_, _ = mm[num*10000]
				mu.RUnlock()
				wg.Done()
			}(q)
		}
		wg.Wait()
	}
}
