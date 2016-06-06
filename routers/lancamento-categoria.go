package routers

import (
    "github.com/gin-gonic/gin"
    "academiasig-api/controllers"
)

func AddRoutesLancamentoCategoria(r *gin.RouterGroup) *gin.RouterGroup {

    r.GET("/lancamento-categorias", controllers.GetCategorias)
    r.GET("/lancamento-categorias/:id", controllers.GetCategoria)
    r.DELETE("/lancamento-categorias/:id", controllers.DeleteCategoria)
    r.POST("/lancamento-categorias", controllers.CreateCategoria)
    r.PUT("/lancamento-categorias/:id", controllers.UpdateCategoria)

    return r
}