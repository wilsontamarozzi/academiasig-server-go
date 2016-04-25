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

	return r
}