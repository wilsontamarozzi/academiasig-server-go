package routers

import (
    "github.com/gin-gonic/gin"
    "academiasig-api/controllers"
)

func AddRoutesLogradouro(r *gin.RouterGroup) *gin.RouterGroup {

    r.GET("/logradouros/:cep", controllers.GetLogradouro)

    return r
}