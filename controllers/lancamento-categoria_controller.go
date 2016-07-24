package controllers

import (
	"github.com/gin-gonic/gin"
	"academiasig-api/services"
	"academiasig-api/services/models"
)

func GetLancamentoCategorias(c *gin.Context) {
	
	lancamentoUUID := c.Query("lancamento-uuid")

	var content models.LancamentoCategorias

	content = services.GetLancamentoCategorias(lancamentoUUID)

	if len(content) <= 0 {
		c.JSON(404, gin.H{"Error": "404", "message": "Registros não encontrado."})
	} else {
		c.JSON(200, content)
	}
}

func GetLancamentoCategoria(c *gin.Context) {

	lancamentoCategoriaId := c.Params.ByName("id")

	content := services.GetLancamentoCategoria(lancamentoCategoriaId)

	if content == (models.LancamentoCategoria{}) {
		c.JSON(404, gin.H{"Status": "404", "message": "Registro não encontrado."})
	} else {
		c.JSON(200, content)
	}	
}

func DeleteLancamentoCategoria(c *gin.Context) {

	lancamentoCategoriaId := c.Params.ByName("id")

	lancamentoCategoria := services.GetLancamentoCategoria(lancamentoCategoriaId)

	if lancamentoCategoria == (models.LancamentoCategoria{}) {
		c.JSON(404, gin.H{"Status": "404", "message": "Registro não encontrado."})
	} else {
		err := services.DeleteLancamentoCategoria(lancamentoCategoriaId)

		if err == nil {
			c.Writer.WriteHeader(204)
		} else {
			c.JSON(500, gin.H{"Status": "500", "Message": "Houve um erro no servidor."})
		}
	}
}

func CreateLancamentoCategoria(c *gin.Context) {

	var lancamentoCategoria models.LancamentoCategoria
	c.Bind(&lancamentoCategoria)

	err := lancamentoCategoria.IsValid()

	if len(err) > 0 {
		c.JSON(422, gin.H{"errors" : err})
	} else {
		lancamentoCategoria = services.CreateLancamentoCategoria(lancamentoCategoria)
		
		if len(lancamentoCategoria.UUID) > 0 {
			c.JSON(201, lancamentoCategoria)
		} else {
			c.JSON(500, gin.H{"Status": "500", "Message": "Houve um erro no servidor."})
		}
	}
}

func UpdateLancamentoCategoria(c *gin.Context) {

	lancamentoCategoriaId := c.Params.ByName("id")
	
	err := services.GetLancamentoCategoria(lancamentoCategoriaId)

	if err == (models.LancamentoCategoria{}) {
		c.JSON(404, gin.H{"Status": "404", "message": "Registro não encontrado."})
	} else {
		var lancamentoCategoria models.LancamentoCategoria
		c.Bind(&lancamentoCategoria)

		err := lancamentoCategoria.IsValid()

		if len(err) > 0 {
			c.JSON(422, gin.H{"errors" : err})
		} else {
			err := services.UpdateLancamentoCategoria(lancamentoCategoria)

			if err == nil {
				c.JSON(201, lancamentoCategoria)
			} else {
				c.JSON(500, gin.H{"Status": "500", "Message": "Houve um erro no servidor."})
			}
		}
	}
}