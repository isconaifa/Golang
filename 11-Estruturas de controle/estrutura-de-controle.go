package main

import "fmt"

func main() {
	fmt.Println("Estruturas de Controle")

	if 7%2 == 0 {
		fmt.Println("7 e par")
	} else {
		fmt.Println("7 e impar")
	}

	if 8%4 == 0 {
		fmt.Println("8 e divisivel por 4")
	}
	numero := 9
	if novoNumero := numero; novoNumero != 9 {
		fmt.Println("Tem algo errado por aqui!")
	} else {
		fmt.Println("Tudo certo por aqui!")
	}
}
