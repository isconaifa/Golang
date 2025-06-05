package main

import "fmt"

func recuperarExecucao() {
	if r := recover(); r != nil {
		fmt.Println("Execucao recuperada com sucesso", r)
	}
}
func alunoEstaAprovado(n1, n2 float64) bool {
	defer recuperarExecucao()
	if n1 < 0 || n2 < 0 {
		panic("A nota do aluno nao pode ser negativa")
	}

	media := (n1 + n2) / 2
	if media >= 6 {
		return true
	}
	return false

}
func main() {

	fmt.Println("Estudando o Panic e o Recover")

	fmt.Println(alunoEstaAprovado(4, -1))
}
