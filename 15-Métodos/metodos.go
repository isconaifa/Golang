package main

import "fmt"

type Pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
}

func (p Pessoa) Salvar() {
	fmt.Printf("Salvando dados de %s no banco de dados!\n", p.nome)
}

func (p *Pessoa) FazerAniversario() string {
	p.idade++
	return fmt.Sprintf("%s fez aniversário e agora tem %d anos", p.nome, p.idade)
}

func (p Pessoa) MaiorIdade() bool {
	return p.idade >= 18
}
func main() {
	pessoa1 := Pessoa{"Ani", "Silva", 17}
	pessoa1.Salvar()

	fmt.Println(pessoa1.FazerAniversario())

	maiorIdade := pessoa1.MaiorIdade()
	fmt.Println(maiorIdade)
}
