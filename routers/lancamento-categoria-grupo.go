package routers

import (
    "github.com/gin-gonic/gin"
    "academiasig-api/controllers"
)

func AddRoutesLancamentoCategoriaGrupo(r *gin.RouterGroup) *gin.RouterGroup {

    r.GET("/lancamento-categoria-grupos", controllers.GetGruposCategoria)
    r.GET("/lancamento-categoria-grupos/:id", controllers.GetGrupoCategoria)
    r.DELETE("/lancamento-categoria-grupos/:id", controllers.DeleteGrupoCategoria)
    r.POST("/lancamento-categoria-grupos", controllers.CreateGrupoCategoria)
    r.PUT("/lancamento-categoria-grupos/:id", controllers.UpdateGrupoCategoria)

    return r
}