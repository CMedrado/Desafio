package main

import "net/http"

type Conta struct {
	ID         int    `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	CPF        string `json:"cpf,omitempty"`
	Secret     string `json:"secret,omitempty"`
	Balance    int    `json:"balance,omitempty"`
	Created_at string `json:"created_At,omitempty"`
}

type ArmazenamentoDeContas struct {
	armazenamento map[string]Conta
}

type MetodosDeArmazenamento interface {
	ObterContas() []Conta
	MostrarSaldo(string) int
	CriarConta(string, string, string)
}

type ServidorConta struct {
	armazenamento MetodosDeArmazenamento
	http.Handler
}
