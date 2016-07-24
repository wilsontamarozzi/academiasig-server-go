package models

import(
	"fmt"
	"time"
	"github.com/satori/go.uuid"
)

const(
	CONTA_PAGAR = "Conta Pagar"
	CONTA_RECEBER = "Conta Receber"
)

type Lancamento struct {
	UUID 				uuid.UUID 		`gorm:"primary_key:true"`
	Descricao			string 
	Tipo				string
	ValorFinal 			float64
	DataCadastro 		*time.Time
	DataVencimento 		*time.Time 
	DataEmissao 		*time.Time
	DataLiquidacao 		*time.Time
	Observacao			string

	PessoaUUID 			uuid.UUID
	CadastradoPorUUID 	uuid.UUID

	Pessoa 				*Pessoa 			`gorm:"ForeignKey:PessoaUUID; AssociationForeignKey:UUID"`
	CadastradoPor 		Pessoa 				`gorm:"ForeignKey:CadastradoPorUUID; AssociationForeignKey:UUID"`

	Categorias 			LancamentoCategorias 	`gorm:"ForeignKey:LancamentoUUID; AssociationForeignKey:UUID"`
	Movimentacoes		LancamentoMovimentacoes `gorm:"ForeignKey:LancamentoUUID; AssociationForeignKey:UUID"`
}

type Lancamentos []Lancamento

func (l Lancamento) IsValid() map[string][]string {

	err := make(map[string][]string)

	if l.Tipo == "" {
		err["tipo"] = append(err[""], "Tipo não pode ser vázio.")	
	}

	/* Validações do lançamento */
	if l.Descricao == "" {
		err["descricao"] = append(err["descricao"], "Descrição não pode ser vázio.")
	}

	if l.ValorFinal == 0 {
		err["valorFinal"] = append(err["valorFinal"], "Valor final não pode ser zero.")	
	}

	if l.DataEmissao == nil {
		err["dataEmissao"] = append(err["dataEmissao"], "Data de emissão não pode estar vázia.")		
	}

	if l.DataVencimento == nil {
		err["dataVencimento"] = append(err["dataVencimento"], "Data de vencimento não pode estar vázia.")
	}

	if calculaSaldoCategoiras(l) != 0 {
        err["saldoCategoria"] = append(err["saldoCategoria"], "Valor das categorias não batem com o valor final do lançamento.")
    }

    if l.DataLiquidacao != nil {
    	if calculaSaldoMovimentacoes(l) != 0 {
    		err["saldoMovimentacao"] = append(err["saldoMovimentacao"], "Valor das movimentações não batem com o valor final do lançamento.")		
    	}
    }

	/*
    fmt.Println("Valor Final: ", l.ValorFinal)
    fmt.Println("Categoria Saldo: ", calculaSaldoCategoiras(l))
    fmt.Println("Categoria: ", calculaCategorias(l.Categorias))
    fmt.Println("Movimentacao Saldo: ", calculaSaldoMovimentacoes(l))
    fmt.Println("Movimentacao: ", calculaMovimentacoes(l.Movimentacoes))
	*/
	
    /*
	Faltam 3 validações
	1- Pessoa não pode ser vázio
	2- Data de vencimento não pode ser inferior a data de emissão
	3- Data de Liquidação não pode ser inferior a data de emissão
	*/

	return err
}

func calculaCategorias(categorias LancamentoCategorias) float64 {

	var valorCategoria float64

	for _, categoria := range categorias {
		if categoria.Valor < 0 {
			valorCategoria = valorCategoria + (categoria.Valor * (-1))
		} else {
			valorCategoria = valorCategoria - categoria.Valor
		}
	}

	return valorCategoria
}

func calculaSaldoCategoiras(l Lancamento) float64 {

	var valorFinal float64
    var valorCategoria float64
    var resultadoFinal float64
    
    if l.Tipo == CONTA_PAGAR {
    	valorFinal = l.ValorFinal * (-1)
    } else {
    	valorFinal = l.ValorFinal
    }
            
    valorCategoria = calculaCategorias(l.Categorias)
    
    fmt.Println("final ",valorFinal)
    fmt.Println("cat ", valorCategoria)

    if l.Tipo == CONTA_PAGAR {
    	resultadoFinal = valorFinal - valorCategoria
    } else {
    	resultadoFinal = valorFinal + valorCategoria
    }
    
    fmt.Println("res ",resultadoFinal)

    return resultadoFinal
}

func calculaMovimentacoes(movimentacoes LancamentoMovimentacoes) float64 {

	var valorMovimentacao float64

	for _, movimentacao := range movimentacoes {
		valorMovimentacao = valorMovimentacao + movimentacao.Valor
	}

	return valorMovimentacao
}

func calculaSaldoMovimentacoes(l Lancamento) float64 {

	var valorFinal float64
    var valorMovimentacao float64
    var resultadoFinal float64
    
    if l.Tipo == CONTA_PAGAR {
    	valorFinal = l.ValorFinal * 1
    } else {
    	valorFinal = l.ValorFinal
    }

    valorMovimentacao = calculaMovimentacoes(l.Movimentacoes)
    
    resultadoFinal = valorFinal - valorMovimentacao

    return resultadoFinal
}