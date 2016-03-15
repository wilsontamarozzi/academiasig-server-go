package main

import (
	"fmt"
	"log"
	"net/http"

	//"github.com/goji/httpauth"
	"github.com/gorilla/mux"
	"AcademiaSIG-API/pessoa"
	"AcademiaSIG-API/pessoa/fisica"
	"AcademiaSIG-API/banco"
	"AcademiaSIG-API/categoria"
	"AcademiaSIG-API/categoriagrupo"
	"AcademiaSIG-API/cidade"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router = router.PathPrefix("/v1/").Subrouter()

	router = pessoa.AddRoutes(router)
    router = pessoafisica.AddRoutes(router)
    router = banco.AddRoutes(router)
    router = categoria.AddRoutes(router)
    router = categoriagrupo.AddRoutes(router)
    router = cidade.AddRoutes(router)

    //http.Handle("/", httpauth.SimpleBasicAuth("w", "1")(router))

	fmt.Println("Server Iniciado -p 8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}