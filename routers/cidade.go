package routers

import (
    "github.com/gin-gonic/gin"
    "academiasig-api/controllers"
)

func AddRoutesCidade(r *gin.RouterGroup) *gin.RouterGroup {

    r.GET("/cidades", controllers.GetCidades)

    return r
}