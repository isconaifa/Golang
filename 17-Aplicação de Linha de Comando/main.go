package main

import (
	"fmt"
	"linha-de-comando/app"
	"log"
	"os"
)

func main() {
	fmt.Println("Ponto de partida da minha aplicação")

	aplicacao := app.Gerar()
	if err := aplicacao.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
