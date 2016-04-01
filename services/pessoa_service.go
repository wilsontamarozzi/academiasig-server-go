package services

import (
	"academiasig-api/services/models"
)

func GetPessoas() models.Pessoas {

	var pessoas models.Pessoas

	Con.Preload("PessoaFisica").
		Preload("PessoaFisica.Usuario").
		Preload("PessoaJuridica").
		Find(&pessoas)

    return pessoas
}

func GetPessoa(pessoaId int64) models.Pessoa {

	var pessoa models.Pessoa

	Con.Preload("PessoaFisica").
		Preload("PessoaFisica.Usuario").
		Preload("PessoaJuridica").
		First(&pessoa, pessoaId)

	return pessoa
}

func GetPessoaPesquisa(pessoaId int64, nome string, ativo bool, email string, tipoPessoa string) models.Pessoas {

	if nome != "" {
		nome = "%" + nome + "%"
	}

	var pessoas models.Pessoas

	Con.Preload("PessoaFisica").
		Preload("PessoaFisica.Usuario").
		Preload("PessoaJuridica").
		Where("pessoa_id = ?", pessoaId).
			Or("nome LIKE ?", nome).
			//Or("ativo = ?", ativo).
			Or("email = ?", email).
			Or("tipo_pessoa = ?", tipoPessoa).
		Find(&pessoas)

	return pessoas
}