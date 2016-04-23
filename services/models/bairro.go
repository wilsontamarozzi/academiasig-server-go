package models

type Bairro struct {
	BairroCodigo		int64	`gorm:"primary_key:true;" sql:"AUTO_INCREMENT"`
	CidadeCodigo		int64
	BairroDescricao		string
	Cidade 				Cidade 	`gorm:"foreignkey:cidade_codigo;associationforeignkey:cidade_codigo"`
}

type Bairros []Bairro