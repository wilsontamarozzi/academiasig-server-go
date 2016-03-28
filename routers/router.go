package routers

import (
	"net/http"

    _ "goji.io"
    "github.com/gorilla/mux"
    _ "github.com/goji/httpauth"
    _ "AcademiaSIG-API/security"
)

//const PathPrefix = "/v1"

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

func InitRoutes() *mux.Router {

    //Cria rota Principal
    router := mux.NewRouter().StrictSlash(true)

    //Adiciona o Prefixo da versão da API (v1)
    router = router.PathPrefix("/v1/").Subrouter()

    //Adiciona as Sub rotas na rota Principal
    router = AddRoutesBanco(router)
	router = AddRoutesPessoa(router)
    router = AddRoutesPessoaFisica(router)
    router = AddRoutesCategoria(router)
    router = AddRoutesCategoriaGrupo(router)
    router = AddRoutesCidade(router)
    router = AddRoutesConta(router)

	return router
}