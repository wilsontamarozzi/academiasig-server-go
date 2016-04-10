package controllers

import (
	"net/http"
	"encoding/json"
	"strconv"

	"github.com/gorilla/mux"
	"academiasig-api/services"
)

/*	@autor: Wilson T.J.

	Método responsável por buscar todas as Pessoas

	Rota: /pessoas.json
*/
func GetPessoas(w http.ResponseWriter, r *http.Request) {
	
	pessoaId	:= r.FormValue("id")
	ativo		:= r.FormValue("ativo")
	nome 		:= r.FormValue("nome")	
	email 		:= r.FormValue("email")
	tipoPessoa 	:= r.FormValue("tipo_pessoa")

	pessoas := services.GetPessoas(pessoaId, ativo, nome, email, tipoPessoa)

	json.NewEncoder(w).Encode(pessoas)
}

/*	@autor: Wilson T.J.

	Método responsável por buscar uma Pessoa especifica pelo ID

	Rota: /pessoas/{id:[0-9]+}.json
*/
func GetPessoa(w http.ResponseWriter, r *http.Request) {

	vars 			:= mux.Vars(r)
	pessoaId, _ 	:= strconv.ParseInt(vars["id"], 0, 64)

	pessoa := services.GetPessoa(pessoaId)

	json.NewEncoder(w).Encode(pessoa)
}