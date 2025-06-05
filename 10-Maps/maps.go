package main

import "fmt"

func main() {
	fmt.Println(" Abordando a estrutura Maps em Go")

	usuario := map[string]string{
		"nome":      "Ani",
		"sobrenome": "silva",
	}

	fmt.Println(usuario["nome"])
}
