package models

type Logradouro struct {
	EnderecoCodigo			int64	`gorm:"primary_key:true;" sql:"AUTO_INCREMENT"`
	BairroCodigo			int64
	EnderecoCep				string
	EnderecoLogradouro 		string
	EnderecoComplemento		string
	Bairro 					Bairro 	`gorm:"foreignkey:bairro_codigo;associationforeignkey:bairro_codigo"`
}

type Logradouros []Logradouro