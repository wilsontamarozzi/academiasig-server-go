package routers

import (
    "github.com/gorilla/mux"
    "AcademiaSIG-API/controllers"
)

func AddRoutesBanco(r *mux.Router) *mux.Router {

    subRouter := r.PathPrefix("/bancos").Subrouter()

    for _, route := range routesBanco {
        subRouter.
            Methods(route.Method).
            Path(route.Pattern).
            Name(route.Name).
            Handler(route.HandlerFunc)
    }

    return r
}

var routesBanco = Routes{
    Route{
        "getBancos",
        "GET",
        "/lista.json",
        controllers.GetBancos,
    },
    Route{
        "getBanco",
        "GET",
        "/{id:[0-9]+}.json",
        controllers.GetBanco,
    },
    Route{
        "getBancoPesquisa",
        "GET",
        "/pesquisa.json",
        controllers.GetBancoPesquisa,
    },
}