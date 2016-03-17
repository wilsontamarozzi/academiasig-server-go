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

	Método responsável por buscar todas as Cidades

	Rota: /cidades.json
*/
func GetCidades(w http.ResponseWriter, r *http.Request) {

	con := database.GetConnection()

	var cidades models.Cidades	

	con.Preload("Estado").Find(&cidades)

	json.NewEncoder(w).Encode(cidades)
}

/*	@autor: Wilson T.J.

	Método responsável por buscar uma Cidade especifica pelo ID

	Rota: /cidades/{id:[0-9]+}.json
*/
func GetCidade(w http.ResponseWriter, r *http.Request) {

	vars 	:= mux.Vars(r)
	id, _ 	:= strconv.ParseInt(vars["id"], 0, 64)

	con := database.GetConnection()

	var cidade models.Cidade

	con.Preload("Estado").First(&cidade, id)

	json.NewEncoder(w).Encode(cidade)
}

/*	@autor: Wilson T.J.

	Método responsável por buscar uma Cidade pelo parametros da queryString

	Rota: /cidades/pesquisa.json?id=ID&nome=NOME
*/
func GetCidadePesquisa(w http.ResponseWriter, r *http.Request) {

	id, _ 	:= strconv.ParseInt(r.FormValue("id"), 0, 64)
	nome 	:= r.FormValue("nome")

	if nome != "" {
		nome = "%" + nome + "%"
	}

	con := database.GetConnection()

	var cidades models.Cidades

	con.Preload("Estado").
		Where("cidade_id = ?", id).
			Or("cidade_nome LIKE ?", nome).
		Find(&cidades)

	json.NewEncoder(w).Encode(cidades)
}