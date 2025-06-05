package main

import (
	"fmt"
)

func main() {
	fmt.Println("trabalhando com loops em Go")

	/*
		i := 0
		for i < 10 {
			i++
			fmt.Println("Incrementando i", i)
			time.Sleep(time.Second)

		}

	*/

	fmt.Println("For com a clausula range")
	nomes := [5]string{"Ani", "Priscila", "Ana", "Márcia", "Esther"}
	for _, nome := range nomes {
		fmt.Println(nome)
	}
	fmt.Println("--------for com range para percorrer string--------------")

	for _, letra := range "FRANCISCO" {
		fmt.Println(string(letra))
	}

	fmt.Println("--------for com range para percorrer maps--------------")

	usuario := map[string]string{
		"nome":      "Ani",
		"nick":      "DJ Baixinha",
		"sobrenome": "silva",
	}
	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}
}
