package controllers

import (
	"net/http"
	"encoding/json"
	"strconv"

	"github.com/gorilla/mux"
	"AcademiaSIG-API/database"
	//"AcademiaSIG-API/services"
	"AcademiaSIG-API/services/models"
)

/*	@autor: Wilson T.J.

	Método responsável por buscar todas as Contas

	Rota: /contas.json
*/
func GetContas(w http.ResponseWriter, r *http.Request) {

	con := database.GetConnection()

	var contas models.Contas	

	con.Preload("Banco").
		Preload("Titular").
		Preload("Operadora").
		Find(&contas)

	json.NewEncoder(w).Encode(contas)
}

/*	@autor: Wilson T.J.

	Método responsável por buscar uma Conta especifica pelo ID

	Rota: /contas/{id:[0-9]+}.json
*/
func GetConta(w http.ResponseWriter, r *http.Request) {

	vars 	:= mux.Vars(r)
	id, _ 	:= strconv.ParseInt(vars["id"], 0, 64)

	con := database.GetConnection()

	var conta models.Conta

	con.Preload("Banco").
		Preload("Titular").
		Preload("Operadora").
		First(&conta, id)

	json.NewEncoder(w).Encode(conta)
}

/*	@autor: Wilson T.J.

	Método responsável por buscar uma Conta pelo parametros da queryString

	Rota: /contas/pesquisa.json?id=ID&nome=NOME
*/
func GetContaPesquisa(w http.ResponseWriter, r *http.Request) {

	id, _ 	:= strconv.ParseInt(r.FormValue("id"), 0, 64)
	nome 	:= r.FormValue("nome")
	numero 	:= r.FormValue("numero")

	if nome != "" {
		nome = "%" + nome + "%"
	}

	con := database.GetConnection()

	var contas models.Contas

	con.Preload("Banco").
		Preload("Titular").
		Preload("Operadora").
		Where("conta_id = ?", id).
			Or("nome LIKE ?", nome).
			Or("numero = ?", numero).
		Find(&contas)

	json.NewEncoder(w).Encode(contas)
}