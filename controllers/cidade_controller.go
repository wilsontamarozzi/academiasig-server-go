package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"academiasig-api/services"
	"academiasig-api/services/models"
)

/*	@autor: Wilson T.J.

	Método responsável por buscar todas as Cidades

	Method: GET
	Rota: /cidades
*/
func GetCidades(c *gin.Context) {

	content := services.GetCidades()

	if len(content) <= 0 {
		c.JSON(404, gin.H{"Error": "404", "message": "Registros não encontrado."})
	} else {
		c.JSON(200, content)
	}
}

/*	@autor: Wilson T.J.

	Método responsável por buscar um Cidade especifica pelo CEP

	Method: GET
	Rota: /cidades/{id:[0-9]+}
*/
func GetCidade(c *gin.Context) {

	cidadeCEP, _ := strconv.ParseInt(c.Params.ByName("cep"), 0, 64)

	content := services.GetCidade(cidadeCEP)

	if content == (models.Cidade{}) {
		c.JSON(404, gin.H{"Error": "404", "message": "Registro não encontrado."})
	} else {
		c.JSON(200, content)
	}	
}