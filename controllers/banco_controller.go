package controllers

import (
	"net/http"
	"encoding/json"
	"strconv"

	"github.com/gorilla/mux"
	"academiasig-api/services"
)

/*	@autor: Wilson T.J.

	Método responsável por buscar todas as Bancos

	Rota: /bancos.json
*/
func GetBancos(w http.ResponseWriter, r *http.Request) {

	bancos := services.GetBancos()

	json.NewEncoder(w).Encode(bancos)
}

/*	@autor: Wilson T.J.

	Método responsável por buscar uma Banco especifica pelo ID

	Rota: /bancos/{id:[0-9]+}.json
*/
func GetBanco(w http.ResponseWriter, r *http.Request) {

	vars 		:= mux.Vars(r)
	bancoId, _ 	:= strconv.ParseInt(vars["id"], 0, 64)

	banco := services.GetBanco(bancoId)

	json.NewEncoder(w).Encode(banco)
}

/*	@autor: Wilson T.J.

	Método responsável por buscar uma Banco pelo parametros da queryString

	Rota: /bancos/pesquisa.json?id=ID&nome=NOME
*/
func GetBancoPesquisa(w http.ResponseWriter, r *http.Request) {

	bancoId, _ 	:= strconv.ParseInt(r.FormValue("id"), 0, 64)
	nome 		:= r.FormValue("nome")
	numero 		:= r.FormValue("numero")

	bancos := services.GetBancoPesquisa(bancoId, nome, numero)

	json.NewEncoder(w).Encode(bancos)
}