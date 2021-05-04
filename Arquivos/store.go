package main

import (
    "math/rand"
    "time"
)

func (a *ArmazenamentoDeContas) CriarConta(name string, cpf string, secret string) []Conta {
    id := 5
    created_at := time.Now().Format("02/01/2006 03:03:05")
    contaNova := []Conta{id, name, cpf, secret, 0, created_at}
    a.armazenamento[name] = contaNova
}

func (a ArmazenamentoDeContas) MostrarSaldo(name string) int {
    
    return a.armazenamento[name]Conta.Balance
}

func InicializaConta() *ArmazenamentoDeContas {
    return &ArmazenamentoDeContas{map[String]Conta{}}
}

func (a *ArmazenamentoDeContas) AdicionaConta(name string, conta []Conta) {
    a.armazenamento[name]++
}

func (a *ArmazenamentoDeContas) ObterContas() []Conta {
    var contas []Conta
    for  nome, []Conta := rage a.armazenamento {
        contas = append(contas, Conta{name, []Conta})
    }
    return contas
}