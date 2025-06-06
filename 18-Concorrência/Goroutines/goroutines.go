package main

import "time"

func main() {
	go escrever("SCP")
	escrever("SLB")
}
func escrever(texto string) {
	for {
		println(texto)
		time.Sleep(1000 * time.Millisecond)
	}
}
