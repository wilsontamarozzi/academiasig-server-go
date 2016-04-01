package services

import (
	"academiasig-api/services/models"
)

func GetPessoasFisica() models.PessoasFisica {

	var pessoasFisica models.PessoasFisica

	Con.Preload("Usuario").
		Find(&pessoasFisica)

    return pessoasFisica
}

func GetPessoaFisica(pessoaFisicaId int64) models.PessoaFisica {

	var pessoaFisica models.PessoaFisica

	Con.Preload("Usuario").
		First(&pessoaFisica, pessoaFisicaId)

	return pessoaFisica
}

func GetPessoasFisicaPesquisa(pessoaFisicaId int64) models.PessoasFisica {

	/*if nome != "" {
		nome = "%" + nome + "%"
	}*/

	var pessoasFisica models.PessoasFisica

	Con.Preload("Usuario").
		Where("pessoa_id = ?", pessoaFisicaId).
		First(&pessoasFisica)

	return pessoasFisica
}