package main

import "fmt"

func somar(n1, n2 int) int {
	return n1 + n2

}

func main() {
	soma := somar(10, 20)
	println(soma)

	var f = func() {
		fmt.Println("Função anonima")
	}
	f()

	fmt.Println(calculosMatematicos(10, 20))
}

func calculosMatematicos(n1, n2 float64) (float64, float64, float64, float64) {
	soma := n1 + n2
	subtracao := n1 - n2
	multiplicacao := n1 * n2
	divisao := n1 / n2
	return soma, subtracao, multiplicacao, divisao
}
