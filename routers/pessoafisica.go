package routers

import (
    "github.com/gorilla/mux"
    "AcademiaSIG-API/controllers"
)

func AddRoutesPessoaFisica(r *mux.Router) *mux.Router {

    subRouter := r.PathPrefix("/pessoas-fisica").Subrouter()

    for _, route := range routesPessoaFisica {
        subRouter.
            Methods(route.Method).
            Path(route.Pattern).
            Name(route.Name).
            Handler(route.HandlerFunc)
    }

    return r
}

var routesPessoaFisica = Routes{
    Route{
        "getPessoas",
        "GET",
        "/lista.json",
        controllers.GetPessoasFisica,
    },
    Route{
        "getPessoa",
        "GET",
        "/{id:[0-9]+}.json",
        controllers.GetPessoaFisica,
    },
    Route{
        "getPessoaPesquisa",
        "GET",
        "/pesquisa.json",
        controllers.GetPessoasFisicaPesquisa,
    },
}