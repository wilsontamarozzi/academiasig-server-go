package controllers

import (
	"github.com/gin-gonic/gin"
	"academiasig-api/services"
	"academiasig-api/services/models"
)

func GetGruposCategoria(c *gin.Context) {
	
	nome 			:= c.Query("nome")
	fullTextSearch 	:= c.Query("search")

	var content models.FinanceiroCategoriaGrupos

	if fullTextSearch == "" {
		content = services.GetGruposCategoria(nome)
	} else {
		content = services.GetGruposCategoriaByFullTextSearch(fullTextSearch)
	}

	if len(content) <= 0 {
		c.JSON(404, gin.H{"Error": "404", "message": "Registros não encontrado."})
	} else {
		c.JSON(200, content)
	}
}

func GetGrupoCategoria(c *gin.Context) {

	gruposCategoriaId := c.Params.ByName("id")

	content := services.GetGrupoCategoria(gruposCategoriaId)

	if content.Nome == "" {
		c.JSON(404, gin.H{"Status": "404", "message": "Registro não encontrado."})
	} else {
		c.JSON(200, content)
	}	
}

func DeleteGrupoCategoria(c *gin.Context) {

	grupoCategoriaId := c.Params.ByName("id")

	grupoCategoria := services.GetGrupoCategoria(grupoCategoriaId)

	if grupoCategoria.Nome == "" {
		c.JSON(404, gin.H{"Status": "404", "message": "Registro não encontrado."})
	} else {
		err := services.DeleteGrupoCategoria(grupoCategoriaId)

		if err == nil {
			c.Writer.WriteHeader(204)
		} else {
			c.JSON(500, gin.H{"Status": "500", "Message": "Houve um erro no servidor."})
		}
	}
}

func CreateGrupoCategoria(c *gin.Context) {

	var grupoCategoria models.FinanceiroCategoriaGrupo
	c.Bind(&grupoCategoria)

	err := grupoCategoria.IsValid()

	if len(err) > 0 {
		c.JSON(422, gin.H{"errors" : err})
	} else {
		grupoCategoria = services.CreateGrupoCategoria(grupoCategoria)
		
		if len(grupoCategoria.UUID) > 0 {
			c.JSON(201, grupoCategoria)
		} else {
			c.JSON(500, gin.H{"Status": "500", "Message": "Houve um erro no servidor."})
		}
	}
}

func UpdateGrupoCategoria(c *gin.Context) {

	grupoCategoriaId := c.Params.ByName("id")
	
	err := services.GetGrupoCategoria(grupoCategoriaId)

	if err.Nome == "" {
		c.JSON(404, gin.H{"Status": "404", "message": "Registro não encontrado."})
	} else {
		var grupoCategoria models.FinanceiroCategoriaGrupo
		c.Bind(&grupoCategoria)

		err := grupoCategoria.IsValid()

		if len(err) > 0 {
			c.JSON(422, gin.H{"errors" : err})
		} else {
			err := services.UpdateGrupoCategoria(grupoCategoria)

			if err == nil {
				c.JSON(201, grupoCategoria)
			} else {
				c.JSON(500, gin.H{"Status": "500", "Message": "Houve um erro no servidor."})
			}
		}
	}
}