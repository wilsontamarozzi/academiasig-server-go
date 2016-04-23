package controllers

import (
	"strconv"

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
	
	pessoaId	:= c.Query("id")
	ativo		:= c.Query("ativo")
	nome 		:= c.Query("nome")	
	email 		:= c.Query("email")
	tipoPessoa 	:= c.Query("tipo_pessoa")

	content := services.GetPessoas(pessoaId, ativo, nome, email, tipoPessoa)

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

	pessoaId, _ := strconv.ParseInt(c.Params.ByName("id"), 0, 64)

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

	pessoaId, _ := strconv.ParseInt(c.Params.ByName("id"), 0, 64)

	pessoa := services.GetPessoa(pessoaId)

	if pessoa == (models.Pessoa{}) {
		c.JSON(404, gin.H{"Status": "404", "message": "Registro não encontrado."})
	} else {
		err := services.DeletePessoa(pessoaId)

		if err == nil {
			c.JSON(204, gin.H{"Status": "204", "Message": "Registro excluido com sucesso."})
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
		
		if pessoa.Id > 0 {
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

	pessoaId, _ := strconv.ParseInt(c.Params.ByName("id"), 0, 64)
	
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