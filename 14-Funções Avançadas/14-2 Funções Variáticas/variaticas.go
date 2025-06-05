package main

import "fmt"

func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

func main() {
	fmt.Println("Funções Variáticas")
	// Funções Variáticas funcionam de modo semelhante ao parametro rest do javascript
	calculo := soma(100, 100, 100)
	fmt.Println(calculo)
}
