package main

import (
	"crud/servidor"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/usuarios", servidor.CriarUsuarios).Methods("POST")
	router.HandleFunc("/usuarios", servidor.BuscarUsuarios).Methods("GET")
	router.HandleFunc("/usuarios/{id}", servidor.BuscarUsuarioPorId).Methods("GET")
	router.HandleFunc("/usuarios/{id}", servidor.AtualizarUsuarios).Methods("PUT")
	router.HandleFunc("/usuarios/{id}", servidor.DeletarUsuarioPorId).Methods("DELETE")

	fmt.Println("Servidor rodando na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
