package routers

import (
    "github.com/gorilla/mux"
    "AcademiaSIG-API/controllers"
)

func AddRoutesCidade(r *mux.Router) *mux.Router {

    subRouter := r.PathPrefix("/cidades").Subrouter()

    for _, route := range routesCidade {
        subRouter.
            Methods(route.Method).
            Path(route.Pattern).
            Name(route.Name).
            Handler(route.HandlerFunc)
    }

    return r
}

var routesCidade = Routes{
    Route{
        "getCidades",
        "GET",
        "/lista.json",
        controllers.GetCidades,
    },
    Route{
        "getCidade",
        "GET",
        "/{id:[0-9]+}.json",
        controllers.GetCidade,
    },
    Route{
        "getCidadePesquisa",
        "GET",
        "/pesquisa.json",
        controllers.GetCidadePesquisa,
    },
}