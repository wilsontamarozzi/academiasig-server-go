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
func GetPessoasFisica(w http.ResponseWriter, r *http.Request) {

	pessoasFisica := services.GetPessoasFisica()

	json.NewEncoder(w).Encode(pessoasFisica)
}

/*	@autor: Wilson T.J.

	Método responsável por buscar uma Pessoa especifica pelo ID

	Rota: /pessoas/{id:[0-9]+}.json
*/
func GetPessoaFisica(w http.ResponseWriter, r *http.Request) {

	vars 				:= mux.Vars(r)
	pessoaFisicaId, _ 	:= strconv.ParseInt(vars["id"], 0, 64)

	pessoaFisica := services.GetPessoaFisica(pessoaFisicaId)

	json.NewEncoder(w).Encode(pessoaFisica)
}

/*	@autor: Wilson T.J.

	Método responsável por buscar uma Pessoa pelo parametros da queryString

	Rota: /pessoas/pesquisa.json?id=ID&nome=NOME
*/
func GetPessoasFisicaPesquisa(w http.ResponseWriter, r *http.Request) {

	pessoaFisicaId, _ 	:= strconv.ParseInt(r.FormValue("id"), 0, 64)
	//nome 				:= r.FormValue("nome")

	pessoasFisica := services.GetPessoasFisicaPesquisa(pessoaFisicaId)

	json.NewEncoder(w).Encode(pessoasFisica)
}