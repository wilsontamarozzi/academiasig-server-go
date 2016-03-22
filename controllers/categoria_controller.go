package controllers

import (
	"net/http"
	"encoding/json"
	"strconv"

	"github.com/gorilla/mux"
	"AcademiaSIG-API/services"
)

/*	@autor: Wilson T.J.

	Método responsável por buscar todas as Categorias

	Rota: /categorias.json
*/
func GetCategorias(w http.ResponseWriter, r *http.Request) {

	categorias := services.GetCategorias()

	json.NewEncoder(w).Encode(categorias)
}

/*	@autor: Wilson T.J.

	Método responsável por buscar uma Categoria especifica pelo ID

	Rota: /categorias/{id:[0-9]+}.json
*/
func GetCategoria(w http.ResponseWriter, r *http.Request) {

	vars 			:= mux.Vars(r)
	categoriaId, _ 	:= strconv.ParseInt(vars["id"], 0, 64)

	categoria := services.GetCategoria(categoriaId)

	json.NewEncoder(w).Encode(categoria)
}

/*	@autor: Wilson T.J.

	Método responsável por buscar uma Categoria pelo parametros da queryString

	Rota: /categorias/pesquisa.json?id=ID&nome=NOME
*/
func GetCategoriaPesquisa(w http.ResponseWriter, r *http.Request) {

	categoriaId, _ 		:= strconv.ParseInt(r.FormValue("id"), 0, 64)
	categoriaGrupoId	:= r.FormValue("categoriaGrupoId")
	nome 				:= r.FormValue("nome")
	tipo 				:= r.FormValue("tipo")

	categorias := services.GetCategoriaPesquisa(categoriaId, categoriaGrupoId, nome, tipo)	

	json.NewEncoder(w).Encode(categorias)
}

func SaveCategoria(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Access-Control-Allow-Origin", "*")
	
	categoriaGrupoId, _	:= strconv.ParseInt(r.FormValue("categoriaGrupoId"), 0, 64)
	nome 				:= r.FormValue("nome")
	
	categoria := services.SaveCategoria(categoriaGrupoId, nome)

	json.NewEncoder(w).Encode(categoria)
}