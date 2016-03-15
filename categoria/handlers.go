package categoria

import (
	"net/http"
	"encoding/json"
	"strconv"

	"AcademiaSIG-API/database"
	"AcademiaSIG-API/categoriagrupo"
	"github.com/gorilla/mux"
)

/*	@autor: Wilson T.J.

	Método responsável por buscar todas as Categorias

	Rota: /categorias.json
*/
func GetCategorias(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	con := database.GetConnection()

	var categorias Categorias	

	con.Table("categoria_lancamento").Find(&categorias)

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

	var categoria Categoria

	con.Table("categoria_lancamento").First(&categoria, id)

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

	var categorias Categorias

	con.Table("categoria_lancamento").
		Where("categoria_id = ?", categoriaId).
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
	
	var grupoCategoria categoriagrupo.GrupoCategoria

	con := database.GetConnection()

	con.Table("categoria_lancamento_grupo").First(&grupoCategoria, categoriaGrupoId)

	categoria := Categoria{
		CategoriaGrupoId: categoriaGrupoId,
		Nome: nome,
		Tipo: grupoCategoria.Tipo,
	}

	con.Table("categoria_lancamento").Create(&categoria)

	json.NewEncoder(w).Encode(categoria)
}