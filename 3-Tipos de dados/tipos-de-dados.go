package main

import (
	"errors"
	"fmt"
)

func main() {
	var numero int16 = 10055
	fmt.Println(numero)

	//alias
	//INT32 = RUNE
	var numero1 rune = 10055
	fmt.Println(numero1)

	var numero2 uint32 = 875
	fmt.Println(numero2)

	//BYTE = UINT8
	var numero3 byte = 100
	fmt.Println(numero3)

	char := 'A'
	fmt.Println(char)

	var erro error = errors.New("Erro interno")
	fmt.Println(erro)
}
