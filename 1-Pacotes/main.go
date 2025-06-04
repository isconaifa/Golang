package main

import (
	"fmt"
	"github.com/badoux/checkmail"
	"modulo/auxiliar"
)

func main() {
	fmt.Println("Usando arquivo de outro diretório no main")
	auxiliar.Escrever()

	erro := checkmail.ValidateFormat("abc@gmail.com")

	fmt.Println(erro)
}
