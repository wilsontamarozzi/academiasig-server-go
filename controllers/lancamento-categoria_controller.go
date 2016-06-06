package controllers

import (
	"github.com/gin-gonic/gin"
	"academiasig-api/services"
	"academiasig-api/services/models"
)

func GetCategorias(c *gin.Context) {
	
	nome 			:= c.Query("nome")
	fullTextSearch 	:= c.Query("search")

	var content models.LancamentoCategorias

	if fullTextSearch == "" {
		content = services.GetCategorias(nome)
	} else {
		content = services.GetCategoriasByFullTextSearch(fullTextSearch)
	}

	if len(content) <= 0 {
		c.JSON(404, gin.H{"Error": "404", "message": "Registros não encontrado."})
	} else {
		c.JSON(200, content)
	}
}

func GetCategoria(c *gin.Context) {

	categoriaId := c.Params.ByName("id")

	content := services.GetCategoria(categoriaId)

	if content == (models.LancamentoCategoria{}) {
		c.JSON(404, gin.H{"Status": "404", "message": "Registro não encontrado."})
	} else {
		c.JSON(200, content)
	}	
}

func DeleteCategoria(c *gin.Context) {

	categoriaId := c.Params.ByName("id")

	categoria := services.GetCategoria(categoriaId)

	if categoria == (models.LancamentoCategoria{}) {
		c.JSON(404, gin.H{"Status": "404", "message": "Registro não encontrado."})
	} else {
		err := services.DeleteCategoria(categoriaId)

		if err == nil {
			c.Writer.WriteHeader(204)
		} else {
			c.JSON(500, gin.H{"Status": "500", "Message": "Houve um erro no servidor."})
		}
	}
}

func CreateCategoria(c *gin.Context) {

	var categoria models.LancamentoCategoria
	c.Bind(&categoria)

	err := categoria.IsValid()

	if len(err) > 0 {
		c.JSON(422, gin.H{"errors" : err})
	} else {
		categoria = services.CreateCategoria(categoria)
		
		if len(categoria.UUID) > 0 {
			c.JSON(201, categoria)
		} else {
			c.JSON(500, gin.H{"Status": "500", "Message": "Houve um erro no servidor."})
		}
	}
}

func UpdateCategoria(c *gin.Context) {

	categoriaId := c.Params.ByName("id")
	
	err := services.GetCategoria(categoriaId)

	if err == (models.LancamentoCategoria{}) {
		c.JSON(404, gin.H{"Status": "404", "message": "Registro não encontrado."})
	} else {
		var categoria models.LancamentoCategoria
		c.Bind(&categoria)

		err := categoria.IsValid()

		if len(err) > 0 {
			c.JSON(422, gin.H{"errors" : err})
		} else {
			err := services.UpdateCategoria(categoria)

			if err == nil {
				c.JSON(201, categoria)
			} else {
				c.JSON(500, gin.H{"Status": "500", "Message": "Houve um erro no servidor."})
			}
		}
	}
}