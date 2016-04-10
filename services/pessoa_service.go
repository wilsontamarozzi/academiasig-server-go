package services

import (
	"academiasig-api/services/models"
)

func GetPessoas(pessoaId string, ativo string, nome string, email string, tipoPessoa string) models.Pessoas {

	pessoaIdQuery 		:= (map[bool]string{true: "pessoa_id = '" 	+ pessoaId + 	"'", false: ""})	[pessoaId != ""]
	ativoQuery 			:= (map[bool]string{true: "ativo = 	'" 		+ ativo + 		"'", false: ""})	[ativo != ""]
	nomeQuery	 		:= (map[bool]string{true: "nome LIKE '%" 	+ nome + 		"%'", false: ""})	[nome != ""]
	emailQuery 			:= (map[bool]string{true: "email LIKE '%" 	+ email + 		"%'", false: ""})	[email != ""]
	tipoPessoaQuery 	:= (map[bool]string{true: "tipo_pessoa = '" + tipoPessoa + 	"'", false: ""})	[tipoPessoa != ""]

	var pessoas models.Pessoas

	Con.Preload("PessoaFisica").
		Preload("PessoaFisica.Usuario").
		Preload("PessoaJuridica").
		Where(pessoaIdQuery).
			Or(ativoQuery).
			Or(nomeQuery).
			Or(emailQuery).
			Or(tipoPessoaQuery).
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