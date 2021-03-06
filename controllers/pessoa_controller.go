package controllers

import (
	"github.com/gin-gonic/gin"
	"academiasig-api/services"
	"academiasig-api/services/models"
)

/*	@autor: Wilson T.J.

	Método responsável por buscar todas as Pessoas

	Method: GET
	Rota: /pessoas
*/
func GetPessoas(c *gin.Context) {
	
	pessoaId		:= c.Query("id")
	ativo			:= c.Query("ativo")
	nome 			:= c.Query("nome")	
	email 			:= c.Query("email")
	tipoPessoa 		:= c.Query("tipo_pessoa")
	usuarioSistema 	:= c.Query("usuario_sistema")
	fullTextSearch 	:= c.Query("search")

	var content models.Pessoas

	if fullTextSearch == "" {
		content = services.GetPessoas(pessoaId, ativo, nome, email, tipoPessoa, usuarioSistema)
	} else {
		content = services.GetPessoasByFullTextSearch(fullTextSearch, tipoPessoa, ativo)
	}

	if len(content) <= 0 {
		c.JSON(404, gin.H{"Error": "404", "message": "Registros não encontrado."})
	} else {
		c.JSON(200, content)
	}
}

/*	@autor: Wilson T.J.

	Método responsável por buscar uma Pessoa especifica pelo ID

	Method: GET
	Rota: /pessoas/{id:[0-9]+}
*/
func GetPessoa(c *gin.Context) {

	pessoaId := c.Params.ByName("id")

	content := services.GetPessoa(pessoaId)

	if content == (models.Pessoa{}) {
		c.JSON(404, gin.H{"Status": "404", "message": "Registro não encontrado."})
	} else {
		c.JSON(200, content)
	}	
}

/*	@autor: Wilson T.J.

	Método responsável por deletar uma Pessoa especifica pelo ID

	Method: DELETE
	Rota: /pessoas/{id:[0-9]+}
*/
func DeletePessoa(c *gin.Context) {

	pessoaId := c.Params.ByName("id")

	pessoa := services.GetPessoa(pessoaId)

	if pessoa == (models.Pessoa{}) {
		c.JSON(404, gin.H{"Status": "404", "message": "Registro não encontrado."})
	} else {
		err := services.DeletePessoa(pessoaId)

		if err == nil {
			c.Writer.WriteHeader(204)
		} else {
			c.JSON(500, gin.H{"Status": "500", "Message": "Houve um erro no servidor."})
		}
	}
}

/*	@autor: Wilson T.J.

	Método responsável por cadastrar uma Pessoa

	Method: POST
	Rota: /pessoas
*/
func CreatePessoa(c *gin.Context) {

	var pessoa models.Pessoa
	c.Bind(&pessoa)

	err := pessoa.IsValid()

	if len(err) > 0 {
		c.JSON(422, gin.H{"errors" : err})
	} else {
		pessoa = services.CreatePessoa(pessoa)
		
		if len(pessoa.UUID) > 0 {
			c.JSON(201, pessoa)
		} else {
			c.JSON(500, gin.H{"Status": "500", "Message": "Houve um erro no servidor."})
		}
	}
}

/*	@autor: Wilson T.J.

	Método responsável por alterar uma Pessoa

	Method: PUT
	Rota: /pessoas/{id:[0-9]+}
*/
func UpdatePessoa(c *gin.Context) {
	
	pessoaId := c.Params.ByName("id")

	err := services.GetPessoa(pessoaId)

	if err == (models.Pessoa{}) {
		c.JSON(404, gin.H{"Status": "404", "message": "Registro não encontrado."})
	} else {
		var pessoa models.Pessoa
		c.Bind(&pessoa)

		err := pessoa.IsValid()

		if len(err) > 0 {
			c.JSON(422, gin.H{"errors" : err})
		} else {
			err := services.UpdatePessoa(pessoa)

			if err == nil {
				c.JSON(201, pessoa)
			} else {
				c.JSON(500, gin.H{"Status": "500", "Message": "Houve um erro no servidor."})
			}
		}
	}
}