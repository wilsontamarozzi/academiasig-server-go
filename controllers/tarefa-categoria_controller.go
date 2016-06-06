package controllers

import (
	"github.com/gin-gonic/gin"
	"academiasig-api/services"
	"academiasig-api/services/models"
)

func GetTarefaCategorias(c *gin.Context) {
	
	descricao		:= c.Query("descricao")
	fullTextSearch 	:= c.Query("search")

	var content models.TarefaCategorias

	if fullTextSearch == "" {
		content = services.GetTarefaCategorias(descricao)
	} else {
		content = services.GetTarefaCategoriasByFullTextSearch(fullTextSearch)
	}

	if len(content) <= 0 {
		c.JSON(404, gin.H{"Error": "404", "message": "Registros não encontrado."})
	} else {
		c.JSON(200, content)
	}
}

func GetTarefaCategoria(c *gin.Context) {

	tarefaCategoriaId := c.Params.ByName("id")

	content := services.GetTarefaCategoria(tarefaCategoriaId)

	if content == (models.TarefaCategoria{}) {
		c.JSON(404, gin.H{"Status": "404", "message": "Registro não encontrado."})
	} else {
		c.JSON(200, content)
	}	
}

func DeleteTarefaCategoria(c *gin.Context) {

	tarefaCategoriaId := c.Params.ByName("id")

	tarefaCategoria := services.GetTarefaCategoria(tarefaCategoriaId)

	if tarefaCategoria == (models.TarefaCategoria{}) {
		c.JSON(404, gin.H{"Status": "404", "message": "Registro não encontrado."})
	} else {
		err := services.DeleteTarefaCategoria(tarefaCategoriaId)

		if err == nil {
			c.Writer.WriteHeader(204)
		} else {
			c.JSON(500, gin.H{"Status": "500", "Message": "Houve um erro no servidor."})
		}
	}
}

func CreateTarefaCategoria(c *gin.Context) {

	var tarefaCategoria models.TarefaCategoria
	c.Bind(&tarefaCategoria)

	err := tarefaCategoria.IsValid()

	if len(err) > 0 {
		c.JSON(422, gin.H{"errors" : err})
	} else {
		tarefaCategoria = services.CreateTarefaCategoria(tarefaCategoria)
		
		if len(tarefaCategoria.UUID) > 0 {
			c.JSON(201, tarefaCategoria)
		} else {
			c.JSON(500, gin.H{"Status": "500", "Message": "Houve um erro no servidor."})
		}
	}
}

func UpdateTarefaCategoria(c *gin.Context) {

	tarefaCategoriaId := c.Params.ByName("id")
	
	err := services.GetTarefaCategoria(tarefaCategoriaId)

	if err == (models.TarefaCategoria{}) {
		c.JSON(404, gin.H{"Status": "404", "message": "Registro não encontrado."})
	} else {
		var tarefaCategoria models.TarefaCategoria
		c.Bind(&tarefaCategoria)

		err := tarefaCategoria.IsValid()

		if len(err) > 0 {
			c.JSON(422, gin.H{"errors" : err})
		} else {
			err := services.UpdateTarefaCategoria(tarefaCategoria)

			if err == nil {
				c.JSON(201, tarefaCategoria)
			} else {
				c.JSON(500, gin.H{"Status": "500", "Message": "Houve um erro no servidor."})
			}
		}
	}
}