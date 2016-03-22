package controllers

import (
	"net/http"
	"encoding/json"
	"strconv"

	"github.com/gorilla/mux"
	"AcademiaSIG-API/services"
)

/*	@autor: Wilson T.J.

	Método responsável por buscar todas as Contas

	Rota: /contas.json
*/
func GetContas(w http.ResponseWriter, r *http.Request) {

	contas := services.GetContas()	

	json.NewEncoder(w).Encode(contas)
}

/*	@autor: Wilson T.J.

	Método responsável por buscar uma Conta especifica pelo ID

	Rota: /contas/{id:[0-9]+}.json
*/
func GetConta(w http.ResponseWriter, r *http.Request) {

	vars 	:= mux.Vars(r)
	contaId, _ 	:= strconv.ParseInt(vars["id"], 0, 64)

	conta := services.GetConta(contaId)	

	json.NewEncoder(w).Encode(conta)
}

/*	@autor: Wilson T.J.

	Método responsável por buscar uma Conta pelo parametros da queryString

	Rota: /contas/pesquisa.json?id=ID&nome=NOME
*/
func GetContaPesquisa(w http.ResponseWriter, r *http.Request) {

	contaId, _ 	:= strconv.ParseInt(r.FormValue("id"), 0, 64)
	descricao	:= r.FormValue("descricao")

	contas := services.GetContaPesquisa(contaId, descricao)

	json.NewEncoder(w).Encode(contas)
}