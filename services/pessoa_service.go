package services

import (
	"academiasig-api/services/models"
)

func GetPessoas(pessoaId string, ativo string, nome string, email string, tipoPessoa string) models.Pessoas {

	pessoaIdQuery 		:= (map[bool]string{true: "id = '" 			+ pessoaId + 	"' AND ", false: ""})	[pessoaId != ""]
	ativoQuery 			:= (map[bool]string{true: "ativo = 	'" 		+ ativo + 		"' AND ", false: ""})	[ativo != ""]
	nomeQuery	 		:= (map[bool]string{true: "nome LIKE '%" 	+ nome + 		"%' AND ", false: ""})	[nome != ""]
	emailQuery 			:= (map[bool]string{true: "email LIKE '%" 	+ email + 		"%' AND ", false: ""})	[email != ""]
	tipoPessoaQuery 	:= (map[bool]string{true: "tipo_pessoa = '" + tipoPessoa + 	"' AND ", false: ""})	[tipoPessoa != ""]	

	commit := pessoaIdQuery + ativoQuery + nomeQuery + emailQuery + tipoPessoaQuery
	if commit != "" {
		commit = commit[:len(commit)-4]
	}

	var pessoas models.Pessoas

	Con.Preload("Usuario").
		Where(commit).
		Find(&pessoas)

    return pessoas
}

func GetPessoa(pessoaId int64) models.Pessoa {

	var pessoa models.Pessoa

	Con.Preload("Usuario").
		First(&pessoa, pessoaId)

	return pessoa
}

func DeletePessoa(pessoaId int64) error {
	err := Con.Where("id = ?", pessoaId).Delete(&models.Pessoa{}).Error

	return err
}

func CreatePessoa(pessoa models.Pessoa) error {
	err := Con.Set("gorm:save_associations", true).Create(&pessoa).Error

	return err
}

func UpdatePessoa(pessoa models.Pessoa) error {
	err := Con.Save(&pessoa).Error

	return err
}