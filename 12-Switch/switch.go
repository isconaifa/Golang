package main

import "fmt"

func dia(num int) string {
	switch num {
	case 1:

		return "Domingo"
	case 2:

		return "Segunda"
	case 3:

		return "Terça"
	case 4:

		return "Quarta"
	case 5:

		return "Quinta"
	case 6:

		return "Sexta"
	case 7:
		return "Sábado"
	default:

		return "Ohh bicho abestalhado, você está digitando o número errado!"
	}
}
func main() {
	fmt.Println("Trabalhando com a estrutura Switch")

	fmt.Println(dia(9))
}
