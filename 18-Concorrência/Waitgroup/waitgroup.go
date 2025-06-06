package main

import (
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		escrever("SCP")
		wg.Done()
	}()

	go func() {
		escrever("SLB")
		wg.Done()
	}()

	wg.Wait()
}
func escrever(texto string) {
	for i := 0; i < 5; i++ {
		println(texto)
		time.Sleep(1000 * time.Millisecond)
	}
}
