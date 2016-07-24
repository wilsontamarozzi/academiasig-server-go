package routers

import (
    "github.com/gin-gonic/gin"
    "academiasig-api/controllers"
)

func AddRoutesLancamentoMovimentacao(r *gin.RouterGroup) *gin.RouterGroup {

    r.GET("/lancamento-movimentacoes", controllers.GetLancamentoMovimentacaos)
    r.GET("/lancamento-movimentacoes/:id", controllers.GetLancamentoMovimentacao)
    r.DELETE("/lancamento-movimentacoes/:id", controllers.DeleteLancamentoMovimentacao)
    r.POST("/lancamento-movimentacoes", controllers.CreateLancamentoMovimentacao)
    r.PUT("/lancamento-movimentacoes/:id", controllers.UpdateLancamentoMovimentacao)

    return r
}