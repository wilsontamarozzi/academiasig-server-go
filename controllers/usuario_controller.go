package controllers

import (
	"github.com/gin-gonic/gin"
	"academiasig-api/services"
	"academiasig-api/services/models"
)

/*	@autor: Wilson T.J.

	Método responsável por buscar um Usuarios pelo login e senha

	Rota: /usuarios.json
*/
func AuthenticationUser(c *gin.Context) {

	var usuario models.Usuario
	c.Bind(&usuario)

	content := services.AuthenticationUser(usuario.Login, usuario.Senha)

	c.JSON(200, content)
}