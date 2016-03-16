package routers

import (
    "github.com/gorilla/mux"
    "AcademiaSIG-API/controllers"
)

func AddRoutesCategoria(r *mux.Router) *mux.Router {

    subRouter := r.PathPrefix("/categorias").Subrouter()

    for _, route := range routesCategoria {
        subRouter.
            Methods(route.Method).
            Path(route.Pattern).
            Name(route.Name).
            Handler(route.HandlerFunc)
    }

    return r
}

var routesCategoria = Routes{
    Route{
        "getCategorias",
        "GET",
        "/lista.json",
        controllers.GetCategorias,
    },
    Route{
        "getCategoria",
        "GET",
        "/{id:[0-9]+}.json",
        controllers.GetCategoria,
    },
    Route{
        "getCategoriaPesquisa",
        "GET",
        "/pesquisa.json",
        controllers.GetCategoriaPesquisa,
    },
    Route{
        "saveCategoria",
        "POST",
        "/save.json",
        controllers.CreateCategoria,
    },
}