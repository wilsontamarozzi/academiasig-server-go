package services

import (
	"academiasig-api/services/models"
)

func GetLogradouros() models.Logradouros {

	var logradouros models.Logradouros

	Con.Preload("Bairro.Cidade.Estado").
		Find(&logradouros)

    return logradouros
}

func GetLogradouro(logradouroCEP int64) models.Logradouro {

	var logradouro models.Logradouro

	Con.Preload("Bairro.Cidade.Estado").
		Where("endereco_cep = ?", logradouroCEP).
		First(&logradouro)

	return logradouro
}