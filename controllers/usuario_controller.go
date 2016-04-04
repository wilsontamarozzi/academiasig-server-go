package controllers

import (
	"fmt"
	"net/http"
	"encoding/json"

	"academiasig-api/services"
)

/*	@autor: Wilson T.J.

	Método responsável por buscar um Usuarios pelo login e senha

	Rota: /usuarios.json
*/
func AuthenticationUser(w http.ResponseWriter, r *http.Request) {

	login := r.FormValue("login")
	senha := r.FormValue("senha")
	
	fmt.Println(login)
	fmt.Println(senha)

	pessoas := services.AuthenticationUser(login, senha)

	json.NewEncoder(w).Encode(pessoas)
}