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

	Método responsável por buscar todas as Pessoas

	Rota: /pessoas.json
*/
func GetPessoasFisica(w http.ResponseWriter, r *http.Request) {

	con := database.GetConnection()

	var pessoasFisica models.PessoasFisica

	con.Table("pessoa_fisica").Find(&pessoasFisica)

	json.NewEncoder(w).Encode(pessoasFisica)
}

/*	@autor: Wilson T.J.

	Método responsável por buscar uma Pessoa especifica pelo ID

	Rota: /pessoas/{id:[0-9]+}.json
*/
func GetPessoaFisica(w http.ResponseWriter, r *http.Request) {

	vars 	:= mux.Vars(r)
	id, _ 	:= strconv.ParseInt(vars["id"], 0, 64)

	con := database.GetConnection()

	var pessoaFisica models.PessoaFisica

	con.Table("pessoa_fisica").First(&pessoaFisica, id)

	json.NewEncoder(w).Encode(pessoaFisica)
}

/*	@autor: Wilson T.J.

	Método responsável por buscar uma Pessoa pelo parametros da queryString

	Rota: /pessoas/pesquisa.json?id=ID&nome=NOME
*/
func GetPessoaFisicaPesquisa(w http.ResponseWriter, r *http.Request) {

	id, _ 	:= strconv.ParseInt(r.FormValue("id"), 0, 64)
	nome 	:= r.FormValue("nome")

	if nome != "" {
		nome = "%" + nome + "%"
	}

	con := database.GetConnection()

	var pessoaFisica models.PessoaFisica

	con.Table("pessoa_fisica").
		Where("pessoa_id = ?", id).
		First(&pessoaFisica)

	json.NewEncoder(w).Encode(pessoaFisica)
}