package routers

import (
    "github.com/gorilla/mux"
    "academiasig-api/controllers"
)

func AddRoutesCategoriaGrupo(r *mux.Router) *mux.Router {

    subRouter := r.PathPrefix("/grupos-categoria").Subrouter()

    for _, route := range routesCategoriaGrupo {
        subRouter.
            Methods(route.Method).
            Path(route.Pattern).
            Name(route.Name).
            Handler(route.HandlerFunc)
    }

    return r
}

var routesCategoriaGrupo = Routes{
    Route{
        "getGruposCategoria",
        "GET",
        "/lista.json",
        controllers.GetGruposCategoria,
    },
    Route{
        "getGrupoCategoria",
        "GET",
        "/{id:[0-9]+}.json",
        controllers.GetGrupoCategoria,
    },
    Route{
        "getGrupoCategoriaPesquisa",
        "GET",
        "/pesquisa.json",
        controllers.GetGrupoCategoriaPesquisa,
    },
}