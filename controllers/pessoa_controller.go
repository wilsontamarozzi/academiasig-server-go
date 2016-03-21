package controllers

import (
	"net/http"
	"encoding/json"
	"strconv"

	"github.com/gorilla/mux"
	"AcademiaSIG-API/services"
)

/*	@autor: Wilson T.J.

	Método responsável por buscar todas as Pessoas

	Rota: /pessoas.json
*/
func GetPessoas(w http.ResponseWriter, r *http.Request) {
	
	pessoas := services.GetPessoas()

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

/*	@autor: Wilson T.J.

	Método responsável por buscar uma Pessoa pelo parametros da queryString

	Rota: /pessoas/pesquisa.json?id=ID&nome=NOME
*/
func GetPessoaPesquisa(w http.ResponseWriter, r *http.Request) {

	pessoaId, _ 	:= strconv.ParseInt(r.FormValue("id"), 0, 64)
	nome 			:= r.FormValue("nome")
	ativo, _		:= strconv.ParseBool(r.FormValue("ativo"))
	email 			:= r.FormValue("email")
	tipoPessoa 		:= r.FormValue("tipoPessoa")

	pessoas := services.GetPessoaPesquisa(pessoaId, nome, ativo, email, tipoPessoa)

	json.NewEncoder(w).Encode(pessoas)
}