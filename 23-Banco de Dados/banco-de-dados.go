package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	stringConexao := "golang:golang@/devbook?charset=utf8&parseTime=True&loc=Local"

	db, err := sql.Open("mysql", stringConexao)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Conexão aberta!")

	linhas, err := db.Query("SELECT * FROM usuarios")
	if err != nil {
		log.Fatal(err)
	}
	defer linhas.Close()
	fmt.Println(linhas)

}
