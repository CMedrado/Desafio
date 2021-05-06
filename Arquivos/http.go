package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func NovoServidorConta(armazenamento MetodosDeArmazenamento) *ServidorConta {
	s := new(ServidorConta)

	s.armazenamento = armazenamento

	roteador := http.NewServeMux()
	roteador.Handle("/accounts/{name}/balance", http.HandlerFunc(s.AçãoSaldo))
	roteador.Handle("/accounts", http.HandlerFunc(s.AçãoMostrarContas))
	roteador.Handle("/accounts/", http.HandlerFunc(s.AçãoCriarConta))

	s.Handler = roteador

	return s
}

func (s *ServidorConta) AçãoSaldo(w http.ResponseWriter, r *http.Request) {
	conta := r.URL.Path[len("/accounts/"):]
	balance := s.armazenamento.MostrarSaldo(conta)
	fmt.Fprint(w, balance)
}

func (s *ServidorConta) AçãoCriarConta(w http.ResponseWriter, r *http.Request) {
	conta := r.URL.Path[len("/accounts/"):]
	//var conta []Conta
	reqBody, _ := ioutil.ReadAll(r.Body)
	//_ = json.NewDecoder(r.Body).Decode(&conta)
	fmt.Fprint(w, "%+v", string(reqBody))
	//fmt.FpBbrint(w, s.armazenamento.ObterContas(conta))
}

func (s *ServidorConta) AçãoMostrarContas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	//json.NewDecoder(w).Encode(s.())
}

func (s *ServidorConta) RegistrarConta(w http.ResponseWriter, conta string) {

	// s.armazenamento.CriarConta
}
