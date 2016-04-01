package routers

import (
    _ "goji.io"
    _ "goji.io/pat"
    "github.com/gorilla/mux"
    "academiasig-api/controllers"
)

func AddRoutesBanco(r *mux.Router) *mux.Router {

    /*
    subRouter := goji.SubMux()
    r.HandleC(pat.New(PathPrefix + "/bancos/*"), subRouter)

    for _, route := range routesBanco {
        if route.Method == "GET" {
            subRouter.HandleFunc(pat.Get(route.Pattern), route.HandlerFunc)
        } else {
            subRouter.HandleFunc(pat.Post(route.Pattern), route.HandlerFunc)
        }
    }
    */

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