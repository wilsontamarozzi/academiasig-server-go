package routers

import (
    "github.com/gorilla/mux"
    "AcademiaSIG-API/controllers"
)

func AddRoutesConta(r *mux.Router) *mux.Router {

    subRouter := r.PathPrefix("/contas").Subrouter()

    for _, route := range routesConta {
        subRouter.
            Methods(route.Method).
            Path(route.Pattern).
            Name(route.Name).
            Handler(route.HandlerFunc)
    }

    return r
}

var routesConta = Routes{
    Route{
        "getContas",
        "GET",
        "/lista.json",
        controllers.GetContas,
    },
    Route{
        "getConta",
        "GET",
        "/{id:[0-9]+}.json",
        controllers.GetConta,
    },
    Route{
        "getContaPesquisa",
        "GET",
        "/pesquisa.json",
        controllers.GetContaPesquisa,
    },
}