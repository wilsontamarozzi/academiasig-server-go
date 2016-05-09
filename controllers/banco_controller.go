package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"academiasig-api/services"
	"academiasig-api/services/models"
)

func GetBancos(c *gin.Context) {
	
	nome 			:= c.Query("nome")
	numero 			:= c.Query("numero")
	fullTextSearch 	:= c.Query("search")

	var content models.Bancos

	if fullTextSearch == "" {
		content = services.GetBancos(nome, numero)
	} else {
		content = services.GetBancosByFullTextSearch(fullTextSearch)
	}

	if len(content) <= 0 {
		c.JSON(404, gin.H{"Error": "404", "message": "Registros não encontrado."})
	} else {
		c.JSON(200, content)
	}
}

func GetBanco(c *gin.Context) {

	bancoId, _ := strconv.ParseInt(c.Params.ByName("id"), 0, 64)

	content := services.GetBanco(bancoId)

	if content == (models.Banco{}) {
		c.JSON(404, gin.H{"Status": "404", "message": "Registro não encontrado."})
	} else {
		c.JSON(200, content)
	}	
}

func DeleteBanco(c *gin.Context) {

	bancoId, _ := strconv.ParseInt(c.Params.ByName("id"), 0, 64)

	banco := services.GetBanco(bancoId)

	if banco == (models.Banco{}) {
		c.JSON(404, gin.H{"Status": "404", "message": "Registro não encontrado."})
	} else {
		err := services.DeleteBanco(bancoId)

		if err == nil {
			c.Writer.WriteHeader(204)
		} else {
			c.JSON(500, gin.H{"Status": "500", "Message": "Houve um erro no servidor."})
		}
	}
}

func CreateBanco(c *gin.Context) {

	var banco models.Banco
	c.Bind(&banco)

	err := banco.IsValid()

	if len(err) > 0 {
		c.JSON(422, gin.H{"errors" : err})
	} else {
		banco = services.CreateBanco(banco)
		
		if banco.Id > 0 {
			c.JSON(201, banco)
		} else {
			c.JSON(500, gin.H{"Status": "500", "Message": "Houve um erro no servidor."})
		}
	}
}

func UpdateBanco(c *gin.Context) {

	bancoId, _ := strconv.ParseInt(c.Params.ByName("id"), 0, 64)
	
	err := services.GetBanco(bancoId)

	if err == (models.Banco{}) {
		c.JSON(404, gin.H{"Status": "404", "message": "Registro não encontrado."})
	} else {
		var banco models.Banco
		c.Bind(&banco)

		err := banco.IsValid()

		if len(err) > 0 {
			c.JSON(422, gin.H{"errors" : err})
		} else {
			err := services.UpdateBanco(banco)

			if err == nil {
				c.JSON(201, banco)
			} else {
				c.JSON(500, gin.H{"Status": "500", "Message": "Houve um erro no servidor."})
			}
		}
	}
}