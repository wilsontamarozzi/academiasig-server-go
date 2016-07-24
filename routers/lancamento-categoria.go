package routers

import (
    "github.com/gin-gonic/gin"
    "academiasig-api/controllers"
)

func AddRoutesLancamentoCategoria(r *gin.RouterGroup) *gin.RouterGroup {

    r.GET("/lancamento-categorias", controllers.GetLancamentoCategorias)
    r.GET("/lancamento-categorias/:id", controllers.GetLancamentoCategoria)
    r.DELETE("/lancamento-categorias/:id", controllers.DeleteLancamentoCategoria)
    r.POST("/lancamento-categorias", controllers.CreateLancamentoCategoria)
    r.PUT("/lancamento-categorias/:id", controllers.UpdateLancamentoCategoria)

    return r
}