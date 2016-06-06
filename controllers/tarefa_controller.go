package controllers

import (
	"github.com/gin-gonic/gin"
	"academiasig-api/services"
	"academiasig-api/services/models"
)

func GetTarefas(c *gin.Context) {
	
	descricao		:= c.Query("descricao")
	concluida		:= c.Query("concluida")
	fullTextSearch 	:= c.Query("search")
	responsavelUUID	:= c.Query("responsavel")
	categoriaUUID	:= c.Query("categoria")
	dataInicio 		:= c.Query("dataInicio")
	dataFim 		:= c.Query("dataFim")
	tipoData		:= c.Query("tipoData")

	var content models.Tarefas

	if fullTextSearch == "" {
		content = services.GetTarefas(descricao, concluida, responsavelUUID, categoriaUUID, tipoData, dataInicio, dataFim)
	} else {
		content = services.GetTarefasByFullTextSearch(fullTextSearch, concluida, responsavelUUID, categoriaUUID, tipoData, dataInicio, dataFim)
	}

	if len(content) <= 0 {
		c.JSON(404, gin.H{"Error": "404", "message": "Registros não encontrado."})
	} else {
		c.JSON(200, content)
	}
}

func GetTarefa(c *gin.Context) {

	tarefaId := c.Params.ByName("id")

	content := services.GetTarefa(tarefaId)

	if content.Descricao == "" {
		c.JSON(404, gin.H{"Status": "404", "message": "Registro não encontrado."})
	} else {
		c.JSON(200, content)
	}	
}

func DeleteTarefa(c *gin.Context) {

	tarefaId := c.Params.ByName("id")

	tarefa := services.GetTarefa(tarefaId)

	if tarefa.Descricao == "" {
		c.JSON(404, gin.H{"Status": "404", "message": "Registro não encontrado."})
	} else {
		err := services.DeleteTarefa(tarefaId)

		if err == nil {
			c.Writer.WriteHeader(204)
		} else {
			c.JSON(500, gin.H{"Status": "500", "Message": "Houve um erro no servidor."})
		}
	}
}

func CreateTarefa(c *gin.Context) {

	var tarefa models.Tarefa
	c.Bind(&tarefa)

	err := tarefa.IsValid()

	if len(err) > 0 {
		c.JSON(422, gin.H{"errors" : err})
	} else {
		tarefa = services.CreateTarefa(tarefa)
		
		if len(tarefa.UUID) > 0 {
			c.JSON(201, tarefa)
		} else {
			c.JSON(500, gin.H{"Status": "500", "Message": "Houve um erro no servidor."})
		}
	}
}

func UpdateTarefa(c *gin.Context) {
	
	tarefaId := c.Params.ByName("id")

	err := services.GetTarefa(tarefaId)

	if err.Descricao == "" {
		c.JSON(404, gin.H{"Status": "404", "message": "Registro não encontrado."})
	} else {
		var tarefa models.Tarefa
		c.Bind(&tarefa)

		err := tarefa.IsValid()

		if len(err) > 0 {
			c.JSON(422, gin.H{"errors" : err})
		} else {
			err := services.UpdateTarefa(tarefa)

			if err == nil {
				c.JSON(201, tarefa)
			} else {
				c.JSON(500, gin.H{"Status": "500", "Message": "Houve um erro no servidor."})
			}
		}
	}
}