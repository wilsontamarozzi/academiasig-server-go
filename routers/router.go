package routers

import (
	"net/http"
	"github.com/gorilla/mux"
)

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

func InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router = router.PathPrefix("/v1/").Subrouter()

	router = AddRoutesBanco(router)
	router = AddRoutesPessoa(router)
    router = AddRoutesPessoaFisica(router)
    router = AddRoutesCategoria(router)
    router = AddRoutesCategoriaGrupo(router)
    router = AddRoutesCidade(router)
    router = AddRoutesConta(router)
	
	return router
}