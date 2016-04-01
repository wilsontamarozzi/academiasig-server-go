package controllers

import (
	"net/http"
	"encoding/json"
	"strconv"

	"github.com/gorilla/mux"
	"academiasig-api/services"
)

/*	@autor: Wilson T.J.

	Método responsável por buscar todas as Cidades

	Rota: /cidades.json
*/
func GetCidades(w http.ResponseWriter, r *http.Request) {

	cidades := services.GetCidades()

	json.NewEncoder(w).Encode(cidades)
}

/*	@autor: Wilson T.J.

	Método responsável por buscar uma Cidade especifica pelo ID

	Rota: /cidades/{id:[0-9]+}.json
*/
func GetCidade(w http.ResponseWriter, r *http.Request) {

	vars 			:= mux.Vars(r)
	cidadeId, _ 	:= strconv.ParseInt(vars["id"], 0, 64)

	cidade := services.GetCidade(cidadeId)

	json.NewEncoder(w).Encode(cidade)
}

/*	@autor: Wilson T.J.

	Método responsável por buscar uma Cidade pelo parametros da queryString

	Rota: /cidades/pesquisa.json?id=ID&nome=NOME
*/
func GetCidadePesquisa(w http.ResponseWriter, r *http.Request) {

	cidadeId, _ 	:= strconv.ParseInt(r.FormValue("id"), 0, 64)
	nome 			:= r.FormValue("nome")

	cidades := services.GetCidadePesquisa(cidadeId, nome)

	json.NewEncoder(w).Encode(cidades)
}