package main

import "fmt"

func generica(inter interface{}) {
	fmt.Println(inter)
}
func main() {
	generica("Ani")
	generica(1)
	generica(true)

	mapa := map[string]interface{}{
		"Nome":         "Ani",
		"Idade":        17,
		"Estado Civil": "Solteira",
		"Altura":       1.67,
		"Cidade":       "Curitiba",
		"Estado":       "PR",
		"Curso":        "Admistração Pública",
	}

	for chave, valor := range mapa {
		fmt.Println(chave, valor)
	}
}
