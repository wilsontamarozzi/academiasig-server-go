package services

import (
	"academiasig-api/services/models"
)

func GetCidades() models.Cidades {

	var cidades models.Cidades

	Con.//Preload("Estado").
		Find(&cidades)

    return cidades
}

func GetCidade(cidadeCEP int64) models.Cidade {

	var cidade models.Cidade

	Con.//Preload("Estado").
		Where("cidade_cep = ?", cidadeCEP).
		First(&cidade)

	return cidade
}