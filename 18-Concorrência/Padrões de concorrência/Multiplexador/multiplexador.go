package main

import (
	"fmt"
	"time"
)

func main() {
	canal := multiplexar(escrever("Canal 1"), escrever("Canal 2"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}

func multiplexar(canalEntrada1, canalEntrada2 <-chan string) <-chan string {
	canalSaida := make(chan string)
	go func() {
		for {
			select {
			case msg1 := <-canalEntrada1:
				canalSaida <- msg1
			case msg2 := <-canalEntrada2:
				canalSaida <- msg2
			}
		}
	}()
	return canalSaida
}

func escrever(texto string) <-chan string {
	canal := make(chan string)
	go func() {
		for {
			canal <- fmt.Sprintf("valor recebido: %s", texto)
			time.Sleep(time.Millisecond * 500)
		}
	}()
	return canal

}
