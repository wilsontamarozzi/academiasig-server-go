package routers

import (
	"net/http"

    "github.com/gin-gonic/gin"
)

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

func InitRoutes() *gin.Engine {

    r := gin.Default()

    v1 := r.Group("api/v1")
    v1 = AddRoutesPessoa(v1)
    v1 = AddRoutesUsuario(v1)
    v1 = AddRoutesLogradouro(v1)
    v1 = AddRoutesCidade(v1)
    v1 = AddRoutesBanco(v1)
    v1 = AddRoutesConta(v1)
    v1 = AddRoutesTarefa(v1)
    v1 = AddRoutesTarefaCategoria(v1)
    v1 = AddRoutesFinanceiroCategoriaGrupo(v1)
    v1 = AddRoutesFinanceiroCategoria(v1)
    v1 = AddRoutesLancamento(v1)
    v1 = AddRoutesLancamentoCategoria(v1)
    v1 = AddRoutesLancamentoMovimentacao(v1)

	return r
}