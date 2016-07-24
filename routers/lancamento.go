package routers

import (
    "github.com/gin-gonic/gin"
    "academiasig-api/controllers"
)

func AddRoutesLancamento(r *gin.RouterGroup) *gin.RouterGroup {

    r.GET("/lancamentos", controllers.GetLancamentos)
    r.GET("/lancamentos/:id", controllers.GetLancamento)
    r.DELETE("/lancamentos/:id", controllers.DeleteLancamento)
    r.POST("/lancamentos", controllers.CreateLancamento)
    r.PUT("/lancamentos/:id", controllers.UpdateLancamento)

    return r
}