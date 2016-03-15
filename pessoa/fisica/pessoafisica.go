package pessoafisica

import (
	"time"
)

type PessoaFisica struct {
	Id 					int64		`gorm:"column:pessoa_fisica_id; primary_key; AUTO_INCREMENT" json:"pessoaFisicaId"`
	PessoaId 			int64 		`gorm:"column:pessoa_id" json:"pessoaId"`
	Sexo				bool		`json:"sexo"`
	DataNascimento		time.Time 	`gorm:"column:data_nascimento: json:"dataNascimento"`
	Cpf 				string		`json:"cpf"`
	Rg					string		`json:"rg"`
	TelefoneResidencial	string		`gorm:"column:telefone_residencial" json:"telefoneResidencial"`
	TelefoneCelular		string		`gorm:"column:telefone_celular" json:"telefoneCelular"`
	TelefoneEmpresa		string		`gorm:"column:telefone_empresa" json:"telefoneEmpresa"`
	UsuarioSistema		bool		`gorm:"column:usuario_sistema" json:"usuarioSistema"`
}

type PessoasFisica []PessoaFisica