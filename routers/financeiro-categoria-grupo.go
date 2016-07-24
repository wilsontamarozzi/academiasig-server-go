package routers

import (
    "github.com/gin-gonic/gin"
    "academiasig-api/controllers"
)

func AddRoutesFinanceiroCategoriaGrupo(r *gin.RouterGroup) *gin.RouterGroup {

    r.GET("/financeiro-categoria-grupos", controllers.GetGruposCategoria)
    r.GET("/financeiro-categoria-grupos/:id", controllers.GetGrupoCategoria)
    r.DELETE("/financeiro-categoria-grupos/:id", controllers.DeleteGrupoCategoria)
    r.POST("/financeiro-categoria-grupos", controllers.CreateGrupoCategoria)
    r.PUT("/financeiro-categoria-grupos/:id", controllers.UpdateGrupoCategoria)

    return r
}