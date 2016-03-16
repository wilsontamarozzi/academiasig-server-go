package controllers

import (
	//"fmt"
	"net/http"
	"encoding/json"
	"strconv"

	"github.com/gorilla/mux"
	"AcademiaSIG-API/database"
	//"AcademiaSIG-API/services"
	"AcademiaSIG-API/services/models"
)

/*	@autor: Wilson T.J.

	Método responsável por buscar todas as Categorias

	Rota: /categorias.json
*/
func GetGruposCategoria(w http.ResponseWriter, r *http.Request) {
	
	con := database.GetConnection()

	gruposCategoria := models.CategoriaLancamentoGrupos{}

	con.Preload("CategoriasLancamento").Find(&gruposCategoria)

	json.NewEncoder(w).Encode(gruposCategoria)
}

/*	@autor: Wilson T.J.

	Método responsável por buscar uma Categoria especifica pelo ID

	Rota: /categorias/{id:[0-9]+}.json
*/
func GetGrupoCategoria(w http.ResponseWriter, r *http.Request) {
	vars 	:= mux.Vars(r)
	id, _ 	:= strconv.ParseInt(vars["id"], 0, 64)

	con := database.GetConnection()

	var grupoCategoria models.CategoriaLancamentoGrupo

	con.Preload("CategoriasLancamento").First(&grupoCategoria, id)

	json.NewEncoder(w).Encode(grupoCategoria)
}

/*	@autor: Wilson T.J.

	Método responsável por buscar uma Categoria pelo parametros da queryString

	Rota: /categorias/pesquisa.json?id=ID&nome=NOME
*/
func GetGrupoCategoriaPesquisa(w http.ResponseWriter, r *http.Request) {

	categoriaGrupoId, _	:= strconv.ParseInt(r.FormValue("categoriaGrupoId"), 0, 64)
	nome 				:= r.FormValue("nome")
	tipo 				:= r.FormValue("tipo")

	if nome != "" {
		nome = "%" + nome + "%"
	}

	con := database.GetConnection()

	var gruposCategoria models.CategoriaLancamentoGrupos

	con.Preload("CategoriasLancamento").
		Where("categoria_grupo_id = ?", categoriaGrupoId).
			Or("nome LIKE ?", nome).
			Or("tipo = ?", tipo).
		Find(&gruposCategoria)

	json.NewEncoder(w).Encode(gruposCategoria)
}