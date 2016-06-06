package routers

import (
    "github.com/gin-gonic/gin"
    "academiasig-api/controllers"
)

func AddRoutesUsuario(r *gin.RouterGroup) *gin.RouterGroup {

	r.GET("/usuarios", controllers.GetUsuarios)
    r.POST("/usuarios/auth", controllers.AuthenticationUser)
    
    return r
}