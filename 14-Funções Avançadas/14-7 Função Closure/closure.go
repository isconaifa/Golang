package main

import "fmt"

func closure() func() {
	text := "Ani"
	funcao := func() {
		fmt.Println(text)
	}
	return funcao
}
func main() {
	text := "ola mundo"
	fmt.Println(text)
	impressao := closure()
	impressao()
}
