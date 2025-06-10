package servidor

import (
	"crud/banco"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Usuario struct {
	Id    uint32 `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

// CriarUsuarios funcao cria usuarios
func CriarUsuarios(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("Ocorreu um erro"))
		return
	}
	//
	var usuario Usuario
	if err = json.Unmarshal(corpoRequisicao, &usuario); err != nil {
		w.Write([]byte("Ocorreu um erro ao converter o json para struct"))
		return
	}
	db, err := banco.Conectar()
	if err != nil {
		w.Write([]byte("Ocorreu um erro ao conectar no banco de dados"))
		return
	}
	defer db.Close()
	stmt, err := db.Prepare("insert into usuarios (nome, email) values (?,?)")
	if err != nil {
		w.Write([]byte("Ocorreu um erro ao criar o statement"))
		return
	}
	defer stmt.Close()

	insercao, err := stmt.Exec(usuario.Nome, usuario.Email)
	if err != nil {
		w.Write([]byte("Ocorreu um erro ao executar o statement"))
		return
	}
	idInserido, err := insercao.LastInsertId()
	if err != nil {
		w.Write([]byte("Ocorreu um erro ao obter o id"))
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuario inserido com sucesso! Id: %d", idInserido)))

	//fmt.Println(usuario)
}

// BuscarUsuarios funcao busca usuarios
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	db, err := banco.Conectar()
	if err != nil {
		w.Write([]byte("Ocorreu um erro ao conectar no banco de dados"))
		return
	}
	defer db.Close()

	linhas, err := db.Query("select * from usuarios")
	if err != nil {
		w.Write([]byte("Ocorreu um erro ao buscar os usuarios"))
		return
	}
	defer linhas.Close()

	var usuarios []Usuario
	for linhas.Next() {
		var usuario Usuario
		if err = linhas.Scan(&usuario.Id, &usuario.Nome, &usuario.Email); err != nil {
			w.Write([]byte("Ocorreu um erro ao escanear o usuario"))
			return
		}
		usuarios = append(usuarios, usuario)
	}
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(usuarios); err != nil {
		w.Write([]byte("Ocorreu um erro ao converter os usuarios para json"))
		return
	}
}

// BuscarUsuario funcao busca usuario especifico
func BuscarUsuarioPorId(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	ID, err := strconv.ParseUint(parametros["id"], 10, 32)
	if err != nil {
		w.Write([]byte("Ocorreu um erro ao converter o id"))
		return
	}
	db, err := banco.Conectar()
	if err != nil {
		w.Write([]byte("Ocorreu um erro ao conectar no banco de dados"))
		return
	}
	defer db.Close()
	linha, err := db.Query("select * from usuarios where id = ?", ID)
	if err != nil {
		w.Write([]byte("Ocorreu um erro ao buscar o usuario"))
		return
	}
	defer linha.Close()
	var usuario Usuario
	if linha.Next() {
		if err = linha.Scan(&usuario.Id, &usuario.Nome, &usuario.Email); err != nil {
			w.Write([]byte("Ocorreu um erro ao escanear o usuario"))
			return
		}
	}
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(usuario); err != nil {
		w.Write([]byte("Ocorreu um erro ao converter o usuario para json"))
		return
	}

}

// AtualizarUsuarios funcao atualiza usuarios
func AtualizarUsuarios(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	ID, err := strconv.ParseUint(parametros["id"], 10, 32)
	if err != nil {
		w.Write([]byte("Ocorreu um erro ao converter o id"))
		return
	}
	corpoRequisicao, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("Ocorreu um erro ao ler o corpo da requisicao"))
		return
	}
	var usuario Usuario
	if err = json.Unmarshal(corpoRequisicao, &usuario); err != nil {
		w.Write([]byte("Ocorreu um erro ao converter o json para struct"))
		return
	}
	db, err := banco.Conectar()
	if err != nil {
		w.Write([]byte("Ocorreu um erro ao conectar no banco de dados"))
		return
	}
	defer db.Close()
	stmt, err := db.Prepare("update usuarios set nome = ?, email = ? where id = ?")
	if err != nil {
		w.Write([]byte("Ocorreu um erro ao criar o statement"))
		return
	}
	defer stmt.Close()
	if _, err = stmt.Exec(usuario.Nome, usuario.Email, ID); err != nil {
		w.Write([]byte("Ocorreu um erro ao executar o statement"))
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// DeletarUsuario funcao deleta usuario especificado
func DeletarUsuarioPorId(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	ID, err := strconv.ParseUint(parametros["id"], 10, 32)
	if err != nil {
		w.Write([]byte("Ocorreu um erro ao converter o id"))
		return
	}
	db, err := banco.Conectar()
	if err != nil {
		w.Write([]byte("Ocorreu um erro ao conectar no banco de dados"))
		return
	}
	defer db.Close()
	stmt, err := db.Prepare("delete from usuarios where id = ?")
	if err != nil {
		w.Write([]byte("Ocorreu um erro ao criar o statement"))
		return
	}
	defer stmt.Close()
	if _, err = stmt.Exec(ID); err != nil {
		w.Write([]byte("Ocorreu um erro ao executar o statement"))
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
