package main

import (
	"fmt"
	"log"
	"net/http"

	"AcademiaSIG-API/routers"
)

func main() {
	//Inicia todas as rotas
	router := routers.InitRoutes()

	fmt.Println("Server Iniciado -p 8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}