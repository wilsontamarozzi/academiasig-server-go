package services

import (
	"academiasig-api/services/models"
)

func GetPessoas(pessoaId string, ativo string, nome string, email string, tipoPessoa string, usuarioSistema string) models.Pessoas {

	var pessoas models.Pessoas

	db := Con

	if pessoaId != "" {
		db = db.Where("id = ?", pessoaId)
	}

	if ativo != "" {
		db = db.Where("ativo = ?", ativo)
	}

	if nome != "" {
		db = db.Where("nome LIKE ?", "%" + nome + "%")	
	}

	if email != "" {
		db = db.Where("email LIKE ?", "%" + email + "%")
	}

	if usuarioSistema != "" {
		db = db.Where("usuario_sistema = ?", usuarioSistema)
	}

	if tipoPessoa != "" {
		db = db.Where("tipo_pessoa = ?", tipoPessoa)
	}
	
	db.Find(&pessoas)

    return pessoas
}

func GetPessoasByFullTextSearch(text string, tipoPessoa string, ativo string) models.Pessoas {

	var pessoas models.Pessoas

	db := Con

	db = db.Table("pessoa").Select("*")

	if tipoPessoa != "" {
		db = db.Where("tipo_pessoa = ?", tipoPessoa)
	}

	if ativo != "" {
		db = db.Where("ativo = ?", ativo)
	}

	db = db.Where(`MATCH(nome, observacao, suporte, complemento, 
			cpf, rg, telefone_celular, telefone_empresa, 
			telefone_residencial, cnpj, fax, 
			inscricao_estadual, inscricao_municipal, 
			razao_social, telefone_comercial, website) 
		AGAINST(?)`, text)

	db.Scan(&pessoas)

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