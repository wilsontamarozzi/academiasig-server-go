package controllers

import (
	"net/http"
	"encoding/json"
	"strconv"

	"github.com/gorilla/mux"
	"academiasig-api/services"
)

/*	@autor: Wilson T.J.

	Método responsável por buscar todas as Categorias

	Rota: /categorias.json
*/
func GetGruposCategoria(w http.ResponseWriter, r *http.Request) {

	gruposCategoria := services.GetGruposCategoria()

	json.NewEncoder(w).Encode(gruposCategoria)
}

/*	@autor: Wilson T.J.

	Método responsável por buscar uma Categoria especifica pelo ID

	Rota: /categorias/{id:[0-9]+}.json
*/
func GetGrupoCategoria(w http.ResponseWriter, r *http.Request) {

	vars 		:= mux.Vars(r)
	grupoId, _ 	:= strconv.ParseInt(vars["id"], 0, 64)

	grupoCategoria := services.GetGrupoCategoria(grupoId)

	json.NewEncoder(w).Encode(grupoCategoria)
}

/*	@autor: Wilson T.J.

	Método responsável por buscar uma Categoria pelo parametros da queryString

	Rota: /categorias/pesquisa.json?id=ID&nome=NOME
*/
func GetGrupoCategoriaPesquisa(w http.ResponseWriter, r *http.Request) {

	grupoId, _	:= strconv.ParseInt(r.FormValue("categoriaGrupoId"), 0, 64)
	nome 		:= r.FormValue("nome")
	tipo 		:= r.FormValue("tipo")

	gruposCategoria := services.GetGrupoCategoriaPesquisa(grupoId, nome, tipo)

	json.NewEncoder(w).Encode(gruposCategoria)
}