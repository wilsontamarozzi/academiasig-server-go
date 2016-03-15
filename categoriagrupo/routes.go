package categoriagrupo

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

    subRouter := r.PathPrefix("/grupos-categoria").Subrouter()

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
        "getGruposCategoria",
        "GET",
        "/lista.json",
        GetGruposCategoria,
    },
    Route{
        "getGrupoCategoria",
        "GET",
        "/{id:[0-9]+}.json",
        GetGrupoCategoria,
    },
    Route{
        "getGrupoCategoriaPesquisa",
        "GET",
        "/pesquisa.json",
        GetGrupoCategoriaPesquisa,
    },
}