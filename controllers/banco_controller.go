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

	Método responsável por buscar todas as Bancos

	Rota: /bancos.json
*/
func GetBancos(w http.ResponseWriter, r *http.Request) {

	con := database.GetConnection()

	var bancos models.Bancos	

	con.Table("banco").Find(&bancos)

	json.NewEncoder(w).Encode(bancos)
}

/*	@autor: Wilson T.J.

	Método responsável por buscar uma Banco especifica pelo ID

	Rota: /bancos/{id:[0-9]+}.json
*/
func GetBanco(w http.ResponseWriter, r *http.Request) {

	vars 	:= mux.Vars(r)
	id, _ 	:= strconv.ParseInt(vars["id"], 0, 64)

	con := database.GetConnection()

	var banco models.Banco

	con.Table("banco").First(&banco, id)

	json.NewEncoder(w).Encode(banco)
}

/*	@autor: Wilson T.J.

	Método responsável por buscar uma Banco pelo parametros da queryString

	Rota: /bancos/pesquisa.json?id=ID&nome=NOME
*/
func GetBancoPesquisa(w http.ResponseWriter, r *http.Request) {

	id, _ 	:= strconv.ParseInt(r.FormValue("id"), 0, 64)
	nome 	:= r.FormValue("nome")
	numero 	:= r.FormValue("numero")

	if nome != "" {
		nome = "%" + nome + "%"
	}

	con := database.GetConnection()

	var bancos models.Bancos

	con.Table("banco").
		Where("banco_id = ?", id).
			Or("nome LIKE ?", nome).
			Or("numero = ?", numero).
		Find(&bancos)

	json.NewEncoder(w).Encode(bancos)
}