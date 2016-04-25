package routers

import (
    "github.com/gin-gonic/gin"
    "academiasig-api/controllers"
)

func AddRoutesBanco(r *gin.RouterGroup) *gin.RouterGroup {

    r.GET("/bancos", controllers.GetBancos)
    r.GET("/bancos/:id", controllers.GetBanco)
    r.DELETE("/bancos/:id", controllers.DeleteBanco)
    r.POST("/bancos", controllers.CreateBanco)
    r.PUT("/bancos/:id", controllers.UpdateBanco)

    return r
}