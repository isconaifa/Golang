// TESTE DE UNIDADE
package enderecos

import "testing"

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoEndereco(t *testing.T) {

	cenariosDeTeste := []cenarioDeTeste{
		{"Rua ABC", "Rua"},
		{"RUA NAIFA", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Rodovia dos Imigrantes", "Rodovia"},
		{"Praceta da Liberdade", "tipo de endereco invalido"},
		{"Estrada dos Imigrantes", "Estrada"},
		{"", "tipo de endereco invalido"},
	}

	for _, cenario := range cenariosDeTeste {
		retornoRecebido := TipoEndereco(cenario.enderecoInserido)
		if retornoRecebido != cenario.retornoEsperado {
			t.Errorf("O valor recebido %s e diferente do esperado %s", retornoRecebido, cenario.retornoEsperado)
		}
	}
}
