package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ArmazenamentoContas interface {
	ObterSaldo(nome string) int
	RegistrarConta(nome string)
	ObterContas() []Conta
}

type Conta struct {
	Nome       string
	Codigo     int
	CPF        string
	Secret     string
	Balance    string
	Created_at string
}

type ServidorContas struct {
	armazenamento ArmazenamentoContas
	http.Handler
}

func NovoServidorConta(armazenamento ArmazenamentoContas) *ServidorContas {
	s := new(ServidorContas)

	s.armazenamento = armazenamento

	roteador := http.NewServeMux()
	roteador.Handle("/accounts/{id}/balance", http.HandlerFunc(s.mostrarConta))
	roteador.Handle("/accounts/", http.HandlerFunc(s.manipulaContas))

	s.Handler = roteador

	return s
}

func (s *ServidorContas) mostrarConta(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(s.armazenamento.ObterContas())
}

func (s *ServidorContas) manipulaContas(w http.ResponseWriter, r *http.Request) {
	conta := r.URL.Path[len("/accounts/"):]

	switch r.Method {
	case http.MethodPost:
		s.cadastrarConta(w, conta)
	case http.MethodGet:
		s.mostrarConta(w, r)
	}
}

func (s *ServidorContas) mostrarSaldo(w http.ResponseWriter, conta string) {
	numero := s.armazenamento.ObterSaldo(conta)

	if numero == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, numero)
}

func (s *ServidorContas) cadastrarConta(w http.ResponseWriter, conta string) {
	s.armazenamento.RegistrarConta(conta)
	w.WriteHeader(http.StatusAccepted)
}
