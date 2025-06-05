package main

import "fmt"

func fibona(pos uint) uint {
	if pos <= 1 {
		return pos
	}
	return fibona(pos-1) + fibona(pos-2)
}
func main() {
	fmt.Println("Funções Recursivas")
	posicao := uint(10)
	fmt.Println(fibona(posicao))
}
