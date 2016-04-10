package routers

import (
    "github.com/gorilla/mux"
    "academiasig-api/controllers"
)

func AddRoutesUsuario(r *mux.Router) *mux.Router {

    subRouter := r.PathPrefix("/usuarios").Subrouter()

    for _, route := range routesUsuario {
        subRouter.
            Methods(route.Method).
            Path(route.Pattern).
            Name(route.Name).
            Handler(route.HandlerFunc)
    }

    return r
}

var routesUsuario = Routes{
    Route{
        "auth",
        "POST",
        "/auth",
        controllers.AuthenticationUser,
    },
}