package main

import (
	"fmt"
	"time"
)

func main() {

	canal := make(chan string)
	go escrever("SCP", canal)

	fmt.Println("Depois da função escrever começar a ser executada")

	for message := range canal {
		fmt.Println(message)
	}
}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto
		time.Sleep(1000 * time.Millisecond)
	}
	close(canal)
}
