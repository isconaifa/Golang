package main

import "fmt"

func alunoEstaAprovado(n1, n2 float64) bool {
	defer fmt.Println("Media calculada. Resultado:")
	media := (n1 + n2) / 2
	if media >= 6 {
		return true
	}
	return false
}
func main() {
	fmt.Println(" Estudando o Defer")
	fmt.Println(alunoEstaAprovado(4, 0))
}
