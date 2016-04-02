package main

import (
	"log"
	"net/http"
	"os"

	"github.com/goji/httpauth"
	"academiasig-api/security"
	"academiasig-api/routers"
)

func main() {
	port := os.Getenv("PORT")

	//Inicia todas as rotas
	router := routers.InitRoutes()

	//Habilita autenticação
	http.Handle("/", httpauth.BasicAuth(security.GetAuthOpts())(router))

	//Printa a Inicialização
	log.Printf("Server Iniciado -p 8080")
	//Inicia o Server
    log.Fatal(http.ListenAndServe(":" + port, nil))
}