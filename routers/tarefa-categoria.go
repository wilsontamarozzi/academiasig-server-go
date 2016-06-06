package routers

import (
    "github.com/gin-gonic/gin"
    "academiasig-api/controllers"
)

func AddRoutesTarefaCategoria(r *gin.RouterGroup) *gin.RouterGroup {

    r.GET("/tarefa-categorias", controllers.GetTarefaCategorias)
    r.GET("/tarefa-categorias/:id", controllers.GetTarefaCategoria)
    r.DELETE("/tarefa-categorias/:id", controllers.DeleteTarefaCategoria)
    r.POST("/tarefa-categorias", controllers.CreateTarefaCategoria)
    r.PUT("/tarefa-categorias/:id", controllers.UpdateTarefaCategoria)

    return r
}