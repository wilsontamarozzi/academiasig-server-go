package services

import (
	"academiasig-api/services/models"
)

func GetBancos() models.Bancos {

	var bancos models.Bancos

	Con.Find(&bancos)

    return bancos
}

func GetBanco(bancoId int64) models.Banco {

	var banco models.Banco

	Con.First(&banco, bancoId)

	return banco
}

func GetBancoPesquisa(bancoId int64, nome string, numero string) models.Bancos {

	if nome != "" {
		nome = "%" + nome + "%"
	}

	var bancos models.Bancos

	Con.Where("banco_id = ?", bancoId).
		Or("nome LIKE ?", nome).
		Or("numero = ?", numero).
		Find(&bancos)

	return bancos
}