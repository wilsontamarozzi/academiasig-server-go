package routers

import (
    "github.com/gin-gonic/gin"
    "academiasig-api/controllers"
)

func AddRoutesTarefa(r *gin.RouterGroup) *gin.RouterGroup {

    r.GET("/tarefas", controllers.GetTarefas)
    r.GET("/tarefas/:id", controllers.GetTarefa)
    r.DELETE("/tarefas/:id", controllers.DeleteTarefa)
    r.POST("/tarefas", controllers.CreateTarefa)
    r.PUT("/tarefas/:id", controllers.UpdateTarefa)

    return r
}