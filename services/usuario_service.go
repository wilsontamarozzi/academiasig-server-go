package services

import (
	"academiasig-api/services/models"
)

func AuthenticationUser(user string, password string) models.Pessoa {

	var pessoa models.Pessoa

	Con.Preload("PessoaFisica").
		Preload("PessoaFisica.Usuario").
		Preload("PessoaJuridica").
		Joins("JOIN pessoa_fisica ON pessoa_fisica.pessoa_id = pessoa.id").
		Joins("JOIN usuario ON usuario.pessoa_fisica_id = pessoa_fisica.pessoa_fisica_id").
		Where("usuario.login = ? AND usuario.senha = ?", user, password).
		First(&pessoa)

	return pessoa
}