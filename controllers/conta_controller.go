package controllers

import (
	"github.com/gin-gonic/gin"
	"academiasig-api/services"
	"academiasig-api/services/models"
)

func GetContas(c *gin.Context) {
	
	descricao 		:= c.Query("descricao")
	tipoConta 		:= c.Query("tipo_conta")
	ativo 			:= c.Query("ativo")
	fullTextSearch 	:= c.Query("search")
	
	var content models.Contas

	if fullTextSearch == "" {
		content = services.GetContas(descricao, tipoConta, ativo)
	} else {
		content = services.GetContasByFullTextSearch(fullTextSearch, tipoConta, ativo)
	}

	if len(content) <= 0 {
		c.JSON(404, gin.H{"Error": "404", "message": "Registros não encontrado."})
	} else {
		c.JSON(200, content)
	}
}

func GetConta(c *gin.Context) {

	contaId := c.Params.ByName("id")

	content := services.GetConta(contaId)

	if content == (models.Conta{}) {
		c.JSON(404, gin.H{"Status": "404", "message": "Registro não encontrado."})
	} else {
		c.JSON(200, content)
	}	
}

func DeleteConta(c *gin.Context) {

	contaId := c.Params.ByName("id")

	conta := services.GetConta(contaId)

	if conta == (models.Conta{}) {
		c.JSON(404, gin.H{"Status": "404", "message": "Registro não encontrado."})
	} else {
		err := services.DeleteConta(contaId)

		if err == nil {
			c.Writer.WriteHeader(204)
		} else {
			c.JSON(500, gin.H{"Status": "500", "Message": "Houve um erro no servidor."})
		}
	}
}

func CreateConta(c *gin.Context) {

	var conta models.Conta
	c.Bind(&conta)

	err := conta.IsValid()

	if len(err) > 0 {
		c.JSON(422, gin.H{"errors" : err})
	} else {
		conta = services.CreateConta(conta)
		
		if len(conta.UUID) > 0 {
			c.JSON(201, conta)
		} else {
			c.JSON(500, gin.H{"Status": "500", "Message": "Houve um erro no servidor."})
		}
	}
}

func UpdateConta(c *gin.Context) {

	contaId := c.Params.ByName("id")
	
	err := services.GetConta(contaId)

	if err == (models.Conta{}) {
		c.JSON(404, gin.H{"Status": "404", "message": "Registro não encontrado."})
	} else {
		var conta models.Conta
		c.Bind(&conta)

		err := conta.IsValid()

		if len(err) > 0 {
			c.JSON(422, gin.H{"errors" : err})
		} else {
			err := services.UpdateConta(conta)

			if err == nil {
				c.JSON(201, conta)
			} else {
				c.JSON(500, gin.H{"Status": "500", "Message": "Houve um erro no servidor."})
			}
		}
	}
}