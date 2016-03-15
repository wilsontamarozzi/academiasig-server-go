package categoria

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

    subRouter := r.PathPrefix("/categorias").Subrouter()

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
        "getCategorias",
        "GET",
        "/lista.json",
        GetCategorias,
    },
    Route{
        "getCategoria",
        "GET",
        "/{id:[0-9]+}.json",
        GetCategoria,
    },
    Route{
        "getCategoriaPesquisa",
        "GET",
        "/pesquisa.json",
        GetCategoriaPesquisa,
    },
    Route{
        "saveCategoria",
        "POST",
        "/save.json",
        CreateCategoria,
    },
}