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
	//Recebe a conexão do banco
	con := database.GetConnection()

	//Cria a array de pessoas
	var pessoasFisica models.PessoasFisica

	//Faz a busca no banco de dados
	con.Table("pessoa_fisica").Find(&pessoasFisica)

	/*con.Table("pessoa_fisica").
		Select("pessoa_fisica.pessoa_fisica_id, pessoa_fisica.cpf, pessoa_fisica.rg, pessoa.pessoa_id, pessoa.nome").
		Joins("INNER JOIN pessoa ON pessoa_fisica.pessoa_id = pessoa.pessoa_id").
		Find(&pessoasFisica)*/

	//Converte e exibe o objeto Pessoas in JSON
	json.NewEncoder(w).Encode(pessoasFisica)
}

/*	@autor: Wilson T.J.

	Método responsável por buscar uma Pessoa especifica pelo ID

	Rota: /pessoas/{id:[0-9]+}.json
*/
func GetPessoaFisica(w http.ResponseWriter, r *http.Request) {
	//Recebe toda a queryString
	vars 	:= mux.Vars(r)
	//Pega o parametro ID da queryString
	//Converte a variavel ID de String para Int
	id, _ 	:= strconv.ParseInt(vars["id"], 0, 64)

	//Recebe a conexão do banco
	con := database.GetConnection()

	//Cria o objeto Pessoa
	var pessoaFisica models.PessoaFisica

	//Faz a Busca no banco de dados
	con.Table("pessoa_fisica").First(&pessoaFisica, id)

	//Converte e exibe o objeto Pessoa in JSON
	json.NewEncoder(w).Encode(pessoaFisica)
}

/*	@autor: Wilson T.J.

	Método responsável por buscar uma Pessoa pelo parametros da queryString

	Rota: /pessoas/pesquisa.json?id=ID&nome=NOME
*/
func GetPessoaFisicaPesquisa(w http.ResponseWriter, r *http.Request) {

	//Recebe os parametros de ID e NOME da queryString
	id, _ 	:= strconv.ParseInt(r.FormValue("id"), 0, 64)
	nome 	:= r.FormValue("nome")

	//Verificar se o nome está vázio para colocar o % do LIKE
	if nome != "" {
		nome = "%" + nome + "%"
	}

	//Recebe a conexão do banco
	con := database.GetConnection()

	//Cria o objeto Pessoa
	var pessoaFisica models.PessoaFisica

	//Faz a Busca no banco de dados
	con.Table("pessoa_fisica").
		Where("pessoa_id = ?", id).
		First(&pessoaFisica)

	//Converte e exibe o objeto Pessoa in JSON
	json.NewEncoder(w).Encode(pessoaFisica)
}