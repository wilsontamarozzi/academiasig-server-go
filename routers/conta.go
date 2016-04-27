package routers

import (
    "github.com/gin-gonic/gin"
    "academiasig-api/controllers"
)

func AddRoutesConta(r *gin.RouterGroup) *gin.RouterGroup {

    r.GET("/contas", controllers.GetContas)
    r.GET("/contas/:id", controllers.GetConta)
    r.DELETE("/contas/:id", controllers.DeleteConta)
    r.POST("/contas", controllers.CreateConta)
    r.PUT("/contas/:id", controllers.UpdateConta)

    return r
}