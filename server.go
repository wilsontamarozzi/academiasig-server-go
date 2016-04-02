package main

import (
	"log"
	"net/http"
	"os"

	"github.com/goji/httpauth"
	"academiasig-api/security"
	"academiasig-api/routers"
)

/*
	Metodo responsável por verificar se há porta setada na variavel ambiente
*/
func getPort() string {
	p := os.Getenv("PORT")

	if p != "" {
		return p
	} else {
		return "8080"
	}
}

func main() {
	//Recebe a porta que irá abrir conexão
	port := getPort()

	//Inicia todas as rotas
	router := routers.InitRoutes()

	//Habilita autenticação
	http.Handle("/", httpauth.BasicAuth(security.GetAuthOpts())(router))

	//Printa a Inicialização
	log.Printf("Server Iniciado -p " + port)
	//Inicia o Server
    log.Fatal(http.ListenAndServe(":" + port, nil))
}