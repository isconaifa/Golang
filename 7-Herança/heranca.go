package main

import "fmt"

type pessoa struct {
	nome           string
	nacionalidade  string
	país           string
	dataNascimento string
}

type estudante struct {
	pessoa
	curso           string
	numeroMatricula string
}

func main() {
	fmt.Println("Falando de herança em golang")
	fmt.Println("------------------------")

	p1 := pessoa{
		nome:           "João da Silva",
		nacionalidade:  "Brasileiro",
		país:           "Brasil",
		dataNascimento: "01/01/2000",
	}

	e1 := estudante{
		pessoa:          p1,
		curso:           "Engenharia da Computação",
		numeroMatricula: "123456",
	}

	fmt.Println(e1.nome)
	fmt.Println(e1.curso)
	fmt.Println(e1.numeroMatricula)
	fmt.Println(e1.país)
	fmt.Println(e1.dataNascimento)
	fmt.Println(e1.nacionalidade)
}
