package main

import (
	"fmt"
	"runtime"
	"time"
)

var size = 100000
var m = make(map[int]*Data, size)

type Data struct {
	Num  int
	Name string
}

func setMap() {
	for i := 1; i <= size; i++ {
		data := Data{Num: i, Name: "name"}
		m[i] = &data
	}
}

func delMap() {
	for i := 1; i <= size; i++ {
		delete(m, i)
	}
}

func resetMap() {
	m = make(map[int]*Data)
}

func main() {
	var m runtime.MemStats
	setMap()
	runtime.ReadMemStats(&m)
	fmt.Printf("memory = %v KB\n", m.Alloc/1024)

	// delete map key
	delMap()
	runtime.GC()
	time.Sleep(100 * time.Millisecond)
	runtime.ReadMemStats(&m)
	fmt.Printf("memory = %v KB\n", m.Alloc/1024)

	resetMap()
	runtime.GC()
	time.Sleep(100 * time.Millisecond)
	runtime.ReadMemStats(&m)
	fmt.Printf("memory = %v KB\n", m.Alloc/1024)
}
