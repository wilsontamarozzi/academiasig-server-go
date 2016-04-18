package routers

import (
    "github.com/gin-gonic/gin"
    "academiasig-api/controllers"
)

func AddRoutesUsuario(r *gin.RouterGroup) *gin.RouterGroup {

    r.POST("/usuarios/auth", controllers.AuthenticationUser)
    
    return r
}