package models

type PessoaJuridica struct {
	PessoaJuridicaId 	int64		`gorm:"primary_key; AUTO_INCREMENT" json:"pessoaJuridicaId"`
	PessoaId 			int64 		`json:"pessoaId"`
	EmpresaSistema		bool		`json:"empresaSistema"`
	RazaoSocial			string		`json:"razaoSocial"`
	Cnpj 				string		`json:"cnpj"`
	TelefoneComercial	string		`json:"telefoneComercial"`
	Fax					string		`json:"fax"`
	Website				string		`json:"website"`
	InscricaoEstadual	string		`json:"inscricaoEstadual"`
	InscricaoMunicipal	string		`json:"inscricaoMunicipal"`
}

type PessoasJuridica []PessoaJuridica