package enderecos

import "strings"

func TipoEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}

	enderecoEmLetraMinuscula := strings.ToLower(endereco)
	primeiraPalvraDoendereco := strings.Split(enderecoEmLetraMinuscula, " ")[0]

	enderecoTemUmTipoValido := false
	for _, tipo := range tiposValidos {
		if tipo == primeiraPalvraDoendereco {
			enderecoTemUmTipoValido = true
			break
		}
	}

	if enderecoTemUmTipoValido {
		return strings.Title(primeiraPalvraDoendereco)
	}
	return "tipo de endereco invalido"
}
