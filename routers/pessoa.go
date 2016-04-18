package routers

import (
    "github.com/gin-gonic/gin"
    "academiasig-api/controllers"
)

func AddRoutesPessoa(r *gin.RouterGroup) *gin.RouterGroup {

    r.GET("/pessoas", controllers.GetPessoas)
    r.GET("/pessoas/:id", controllers.GetPessoa)
    r.DELETE("/pessoas/:id", controllers.DeletePessoa)
    r.POST("/pessoas", controllers.CreatePessoa)
    r.PUT("/pessoas/:id", controllers.UpdatePessoa)

    return r
}