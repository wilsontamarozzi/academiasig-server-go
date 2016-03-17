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
func GetPessoas(w http.ResponseWriter, r *http.Request) {
	//Recebe a conexão do banco
	con := database.GetConnection()

	//Cria a array de pessoas
	var pessoas models.Pessoas	

	//Faz a busca no banco de dados
	con.Preload("PessoaFisica.Usuario").Preload("PessoaJuridica").Find(&pessoas)

	//Converte e exibe o objeto Pessoas in JSON
	json.NewEncoder(w).Encode(pessoas)
}

/*	@autor: Wilson T.J.

	Método responsável por buscar uma Pessoa especifica pelo ID

	Rota: /pessoas/{id:[0-9]+}.json
*/
func GetPessoa(w http.ResponseWriter, r *http.Request) {
	//Recebe toda a queryString
	vars 	:= mux.Vars(r)
	//Pega o parametro ID da queryString
	//Converte a variavel ID de String para Int
	id, _ 	:= strconv.ParseInt(vars["id"], 0, 64)

	//Recebe a conexão do banco
	con := database.GetConnection()

	//Cria o objeto Pessoa
	var pessoa models.Pessoa

	//Faz a Busca no banco de dados
	con.First(&pessoa, id)

	//Converte e exibe o objeto Pessoa in JSON
	json.NewEncoder(w).Encode(pessoa)
}

/*	@autor: Wilson T.J.

	Método responsável por buscar uma Pessoa pelo parametros da queryString

	Rota: /pessoas/pesquisa.json?id=ID&nome=NOME
*/
func GetPessoaPesquisa(w http.ResponseWriter, r *http.Request) {

	//Recebe os parametros de ID e NOME da queryString
	id, _ 		:= strconv.ParseInt(r.FormValue("id"), 0, 64)
	nome 		:= r.FormValue("nome")
	ativo 		:= r.FormValue("ativo")
	email 		:= r.FormValue("email")
	tipoPessoa 	:= r.FormValue("tipoPessoa")

	//Verificar se o nome está vázio para colocar o % do LIKE
	if nome != "" {
		nome = "%" + nome + "%"
	}

	//Recebe a conexão do banco
	con := database.GetConnection()

	//Cria o objeto Pessoa
	var pessoas models.Pessoas

	//Faz a Busca no banco de dados
	con.Where("pessoa_id = ?", id).
			Or("nome LIKE ?", nome).
			Or("ativo = ?", ativo).
			Or("email = ?", email).
			Or("tipo_pessoa = ?", tipoPessoa).
		Find(&pessoas)

	//Converte e exibe o objeto Pessoa in JSON
	json.NewEncoder(w).Encode(pessoas)
}