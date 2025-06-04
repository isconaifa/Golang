package main

import (
	"fmt"
)

type usuario struct {
	nome     string
	idade    int8
	endereco Endereco
}

type Endereco struct {
	Logradouro  string
	Numero      int8
	Complemento string
	Cidade      string
	Estado      string
}

func main() {
	fmt.Println("Structs")
	dadosEndereco := Endereco{
		Logradouro:  "\nPaís de Gales\n",
		Numero:      93,
		Complemento: "Apartamento\n",
		Cidade:      "Cuiabá\n",
		Estado:      "Mato Grosso\n",
	}

	/*var u usuario
	u.nome = "Márcia"
	u.idade = 25
	fmt.Println(u)
	*/

	fmt.Println("----------------")
	usuario2 := usuario{"Esther Oliveira\n", +21, dadosEndereco}
	fmt.Println(usuario2)
	fmt.Println("----------------")
	usuario3 := usuario{nome: "Ani Silva"}
	fmt.Println(usuario3)
}
