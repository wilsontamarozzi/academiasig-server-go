package routers

import (
    "github.com/gin-gonic/gin"
    "academiasig-api/controllers"
)

func AddRoutesFinanceiroCategoria(r *gin.RouterGroup) *gin.RouterGroup {

    r.GET("/financeiro-categorias", controllers.GetCategorias)
    r.GET("/financeiro-categorias/:id", controllers.GetCategoria)
    r.DELETE("/financeiro-categorias/:id", controllers.DeleteCategoria)
    r.POST("/financeiro-categorias", controllers.CreateCategoria)
    r.PUT("/financeiro-categorias/:id", controllers.UpdateCategoria)

    return r
}