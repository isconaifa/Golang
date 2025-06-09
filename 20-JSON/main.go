package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint8  `json:"idade"`
}

func main() {
	c := Cachorro{"Rex", "Dalmata", 3}

	cachorroJSON, err := json.Marshal(c)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(cachorroJSON))
}
