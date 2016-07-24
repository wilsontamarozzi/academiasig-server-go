package controllers

import (
	"github.com/gin-gonic/gin"
	"academiasig-api/services"
	"academiasig-api/services/models"
)

func GetLancamentos(c *gin.Context) {
	
	descricao		:= c.Query("descricao")
	fullTextSearch 	:= c.Query("search")

	var content models.Lancamentos

	if fullTextSearch == "" {
		content = services.GetLancamentos(descricao)
	} else {
		content = services.GetLancamentosByFullTextSearch(fullTextSearch)
	}

	if len(content) <= 0 {
		c.JSON(404, gin.H{"Error": "404", "message": "Registros não encontrado."})
	} else {
		c.JSON(200, content)
	}
}

func GetLancamento(c *gin.Context) {

	lancamentoId := c.Params.ByName("id")

	content := services.GetLancamento(lancamentoId)

	if content.Descricao == "" {
		c.JSON(404, gin.H{"Status": "404", "message": "Registro não encontrado."})
	} else {
		c.JSON(200, content)
	}	
}

func DeleteLancamento(c *gin.Context) {

	lancamentoId := c.Params.ByName("id")

	lancamento := services.GetLancamento(lancamentoId)

	if lancamento.Descricao == "" {
		c.JSON(404, gin.H{"Status": "404", "message": "Registro não encontrado."})
	} else {
		err := services.DeleteLancamento(lancamentoId)

		if err == nil {
			c.Writer.WriteHeader(204)
		} else {
			c.JSON(500, gin.H{"Status": "500", "Message": "Houve um erro no servidor."})
		}
	}
}

func CreateLancamento(c *gin.Context) {

	var lancamento models.Lancamento
	c.Bind(&lancamento)

	err := lancamento.IsValid()

	if len(err) > 0 {
		c.JSON(422, gin.H{"errors" : err})
	} else {
		lancamento = services.CreateLancamento(lancamento)
		
		if len(lancamento.UUID) > 0 {
			c.JSON(201, lancamento)
		} else {
			c.JSON(500, gin.H{"Status": "500", "Message": "Houve um erro no servidor."})
		}
	}
}

func UpdateLancamento(c *gin.Context) {

	lancamentoId := c.Params.ByName("id")
	
	err := services.GetLancamento(lancamentoId)

	if err.Descricao == "" {
		c.JSON(404, gin.H{"Status": "404", "message": "Registro não encontrado."})
	} else {
		var lancamento models.Lancamento
		c.Bind(&lancamento)

		err := lancamento.IsValid()

		if len(err) > 0 {
			c.JSON(422, gin.H{"errors" : err})
		} else {
			err := services.UpdateLancamento(lancamento)

			if err == nil {
				c.JSON(201, lancamento)
			} else {
				c.JSON(500, gin.H{"Status": "500", "Message": "Houve um erro no servidor."})
			}
		}
	}
}