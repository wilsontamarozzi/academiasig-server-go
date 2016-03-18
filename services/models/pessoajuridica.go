package models

type PessoaJuridica struct {
	PessoaJuridicaId 	int64		`gorm:"primary_key; AUTO_INCREMENT"`
	PessoaId 			int64
	EmpresaSistema		bool
	RazaoSocial			string
	Cnpj 				string
	TelefoneComercial	string
	Fax					string
	Website				string
	InscricaoEstadual	string
	InscricaoMunicipal	string
}

type PessoasJuridica []PessoaJuridica