package controllers

import (
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
func GetCategorias(w http.ResponseWriter, r *http.Request) {

	con := database.GetConnection()

	categorias := models.CategoriaLancamento{}

	con.Find(&categorias)

	json.NewEncoder(w).Encode(categorias)
}

/*	@autor: Wilson T.J.

	Método responsável por buscar uma Categoria especifica pelo ID

	Rota: /categorias/{id:[0-9]+}.json
*/
func GetCategoria(w http.ResponseWriter, r *http.Request) {
	vars 	:= mux.Vars(r)
	id, _ 	:= strconv.ParseInt(vars["id"], 0, 64)

	con := database.GetConnection()

	categoria := models.CategoriaLancamento{}

	con.First(&categoria, id)

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

	if nome != "" {
		nome = "%" + nome + "%"
	}

	con := database.GetConnection()

	categorias := models.CategoriaLancamento{}

	con.Where("categoria_id = ?", categoriaId).
			Or("categoria_grupo_id = ?", categoriaGrupoId).
			Or("nome LIKE ?", nome).
			Or("tipo = ?", tipo).
		Find(&categorias)

	json.NewEncoder(w).Encode(categorias)
}

func CreateCategoria(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	
	categoriaGrupoId, _	:= strconv.ParseInt(r.FormValue("categoriaGrupoId"), 0, 64)
	nome 				:= r.FormValue("nome")
	
	var grupoCategoria models.CategoriaLancamentoGrupo

	con := database.GetConnection()

	con.First(&grupoCategoria, categoriaGrupoId)

	categoria := models.CategoriaLancamento{
		CategoriaGrupoId: categoriaGrupoId,
		Nome: nome,
		Tipo: grupoCategoria.Tipo,
	}

	con.Create(&categoria)

	json.NewEncoder(w).Encode(categoria)
}