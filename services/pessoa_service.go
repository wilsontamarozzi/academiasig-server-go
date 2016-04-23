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

	Con.Where(commit).Find(&pessoas)

	/*
	for i, _ := range pessoas {
        Con.Preload("Logradouro.Bairro.Cidade.Estado").
		Preload("Usuario").
		First(&pessoas[i])
    }
    */

    return pessoas
}

func GetPessoa(pessoaId int64) models.Pessoa {

	var pessoa models.Pessoa

	Con.Preload("Logradouro.Bairro.Cidade.Estado").
		Preload("Usuario").
		First(&pessoa, pessoaId)

	return pessoa
}

func DeletePessoa(pessoaId int64) error {
	err := Con.Where("id = ?", pessoaId).Delete(&models.Pessoa{}).Error

	return err
}

func CreatePessoa(pessoa models.Pessoa) models.Pessoa {
	err := Con.Set("gorm:save_associations", false).Create(&pessoa).Error

	if err != nil {
		panic(err)
	}

	return pessoa
}

func UpdatePessoa(pessoa models.Pessoa) error {
	err := Con.Set("gorm:save_associations", false).Save(&pessoa).Error

	return err
}