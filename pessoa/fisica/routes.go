package pessoafisica

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

func AddRoutes(r *mux.Router) *mux.Router {

    subRouter := r.PathPrefix("/pessoas-fisica").Subrouter()

    for _, route := range routes {
        subRouter.
            Methods(route.Method).
            Path(route.Pattern).
            Name(route.Name).
            Handler(route.HandlerFunc)
    }

    return r
}

var routes = Routes{
    Route{
        "getPessoas",
        "GET",
        "/lista.json",
        GetPessoas,
    },
    Route{
        "getPessoa",
        "GET",
        "/{id:[0-9]+}.json",
        GetPessoa,
    },
    Route{
        "getPessoaPesquisa",
        "GET",
        "/pesquisa.json",
        GetPessoaPesquisa,
    },
}