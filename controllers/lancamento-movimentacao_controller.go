package controllers

import (
	"github.com/gin-gonic/gin"
	"academiasig-api/services"
	"academiasig-api/services/models"
)

func GetLancamentoMovimentacaos(c *gin.Context) {
	
	lancamentoUUID := c.Query("lancamento-uuid")

	var content models.LancamentoMovimentacoes

	content = services.GetLancamentoMovimentacaos(lancamentoUUID)

	if len(content) <= 0 {
		c.JSON(404, gin.H{"Error": "404", "message": "Registros não encontrado."})
	} else {
		c.JSON(200, content)
	}
}

func GetLancamentoMovimentacao(c *gin.Context) {

	lancamentoMovimentacaoId := c.Params.ByName("id")

	content := services.GetLancamentoMovimentacao(lancamentoMovimentacaoId)

	if content == (models.LancamentoMovimentacao{}) {
		c.JSON(404, gin.H{"Status": "404", "message": "Registro não encontrado."})
	} else {
		c.JSON(200, content)
	}	
}

func DeleteLancamentoMovimentacao(c *gin.Context) {

	lancamentoMovimentacaoId := c.Params.ByName("id")

	lancamentoMovimentacao := services.GetLancamentoMovimentacao(lancamentoMovimentacaoId)

	if lancamentoMovimentacao == (models.LancamentoMovimentacao{}) {
		c.JSON(404, gin.H{"Status": "404", "message": "Registro não encontrado."})
	} else {
		err := services.DeleteLancamentoMovimentacao(lancamentoMovimentacaoId)

		if err == nil {
			c.Writer.WriteHeader(204)
		} else {
			c.JSON(500, gin.H{"Status": "500", "Message": "Houve um erro no servidor."})
		}
	}
}

func CreateLancamentoMovimentacao(c *gin.Context) {

	var lancamentoMovimentacao models.LancamentoMovimentacao
	c.Bind(&lancamentoMovimentacao)

	err := lancamentoMovimentacao.IsValid()

	if len(err) > 0 {
		c.JSON(422, gin.H{"errors" : err})
	} else {
		lancamentoMovimentacao = services.CreateLancamentoMovimentacao(lancamentoMovimentacao)
		
		if len(lancamentoMovimentacao.UUID) > 0 {
			c.JSON(201, lancamentoMovimentacao)
		} else {
			c.JSON(500, gin.H{"Status": "500", "Message": "Houve um erro no servidor."})
		}
	}
}

func UpdateLancamentoMovimentacao(c *gin.Context) {

	lancamentoMovimentacaoId := c.Params.ByName("id")
	
	err := services.GetLancamentoMovimentacao(lancamentoMovimentacaoId)

	if err == (models.LancamentoMovimentacao{}) {
		c.JSON(404, gin.H{"Status": "404", "message": "Registro não encontrado."})
	} else {
		var lancamentoMovimentacao models.LancamentoMovimentacao
		c.Bind(&lancamentoMovimentacao)

		err := lancamentoMovimentacao.IsValid()

		if len(err) > 0 {
			c.JSON(422, gin.H{"errors" : err})
		} else {
			err := services.UpdateLancamentoMovimentacao(lancamentoMovimentacao)

			if err == nil {
				c.JSON(201, lancamentoMovimentacao)
			} else {
				c.JSON(500, gin.H{"Status": "500", "Message": "Houve um erro no servidor."})
			}
		}
	}
}