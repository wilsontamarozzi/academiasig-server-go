package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"academiasig-api/services"
	"academiasig-api/services/models"
)

/*	@autor: Wilson T.J.

	Método responsável por buscar todos os Logradouros

	Method: GET
	Rota: /logradouros
*/
func GetLogradouros(c *gin.Context) {

	content := services.GetLogradouros()

	if len(content) <= 0 {
		c.JSON(404, gin.H{"Error": "404", "message": "Registros não encontrado."})
	} else {
		c.JSON(200, content)
	}
}

/*	@autor: Wilson T.J.

	Método responsável por buscar um Logradouro especifica pelo CEP

	Method: GET
	Rota: /logradouros/{id:[0-9]+}
*/
func GetLogradouro(c *gin.Context) {

	logradouroCEP, _ := strconv.ParseInt(c.Params.ByName("cep"), 0, 64)

	content := services.GetLogradouro(logradouroCEP)

	if content == (models.Logradouro{}) {
		c.JSON(404, gin.H{"Error": "404", "message": "Registro não encontrado."})
	} else {
		c.JSON(200, content)
	}	
}