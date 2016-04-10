package routers

import (
    "github.com/gorilla/mux"
    "academiasig-api/controllers"
)

func AddRoutesPessoa(r *mux.Router) *mux.Router {

    subRouter := r.PathPrefix("/pessoas").Subrouter()

    for _, route := range routesPessoa {
        subRouter.
            Methods(route.Method).
            Path(route.Pattern).
            Name(route.Name).
            Handler(route.HandlerFunc)
    }

    return r
}

var routesPessoa = Routes{
    Route{
        "getPessoas",
        "GET",
        "/lista.json",
        controllers.GetPessoas,
    },
    Route{
        "getPessoa",
        "GET",
        "/{id:[0-9]+}.json",
        controllers.GetPessoa,
    },
}