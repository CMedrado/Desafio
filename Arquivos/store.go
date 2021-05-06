package main

import (
	"time"
)

func (a *ArmazenamentoDeContas) CriarConta(name string, cpf string, secret string) {
	id := 5
	created_at := time.Now().Format("02/01/2006 03:03:05")
	contaNova := Conta{id, name, cpf, secret, 0, created_at}
	a.armazenamento[name] = contaNova
}

func (a ArmazenamentoDeContas) MostrarSaldo(name string) int {
	conta := a.armazenamento[name]
	return conta.Balance
}

func InicializaConta() *ArmazenamentoDeContas {
	return &ArmazenamentoDeContas{map[string]Conta{}}
}

func (a *ArmazenamentoDeContas) ObterContas() []Conta {
	var contas []Conta
	for name, Conta := range a.armazenamento {
		contas = append(contas, Conta{name, Conta})
	}
	return contas
}
