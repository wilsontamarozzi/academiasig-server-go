package services

import (
	"academiasig-api/services/models"
)

func GetCidades() models.Cidades {

	var cidades models.Cidades	

	Con.Find(&cidades)

    for i, _ := range cidades {
        Con.Model(&cidades[i]).Related(&cidades[i].Estado)
        Con.Model(&cidades[i].Estado).Related(&cidades[i].Estado.Pais)
    }

    return cidades
}

func GetCidade(cidadeId int64) models.Cidade {

	var cidade models.Cidade

	Con.First(&cidade, cidadeId)
	Con.Model(&cidade).Related(&cidade.Estado)
    Con.Model(&cidade.Estado).Related(&cidade.Estado.Pais)

	return cidade
}

func GetCidadePesquisa(cidadeId int64, nome string) models.Cidades {

	if nome != "" {
		nome = "%" + nome + "%"
	}

	var cidades models.Cidades

	Con.Where("cidade_id = ?", cidadeId).
		Or("nome LIKE ?", nome).
		Find(&cidades)

    for i, _ := range cidades {
        Con.Model(&cidades[i]).Related(&cidades[i].Estado)
        Con.Model(&cidades[i].Estado).Related(&cidades[i].Estado.Pais)
    }

	return cidades
}