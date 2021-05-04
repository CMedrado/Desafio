package main

import (
    "database/sql"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"

    "github.com/gorilla/mux@v1.8.0"
)

func NovoServidorConta(armazenamento MetodosDeArmazenamento) *ServidorConta {
    s := new(ServidorConta)

    roteador := mux.NewRouter

    roteador.HandlerFunc("/accounts/%s/balance", s.AçãoSaldo).Methods("GET")
    roteador.HandlerFunc("/accounts/", s.AçãoCriarConta).Methods("POST")
    roteador.HandlerFunc("/accounts/", s.AçãoMostrarContas).Methods("GET")

    s.Handler = roteador

    return s
}

func (s *ServidorConta) AçãoSaldo(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("content-type", "application/json")
    json.NewDecoder(w).Encode(s.MostrarSaldo())
}

func (s *ServidorConta) AçãoCriarConta(w http.ResponseWriter, r *http.Request) {
    var conta []Conta
    reqBody, _ := ioutil.ReadAll(r.Body)
    _ = json.NewDecoder(r.Body).Decode(&conta)
    
    fmt.FpBbrint(w, s.armazenamento.ObterContas(conta))
}

func (s *ServidorConta) AçãoMostrarContas(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("content-type", "application/json")
    json.NewDecoder(w).Encode(s.())
}

func (s *ServidorConta) RegistrarConta(w http.ResponseWriter, conta string) {

    s.armazenamento.CriarConta
}

func OpenConnection() *sql.DB {
    psqlInfo := fmt.Sprintf("name=%s cpf=%s secret=%s", name, cpf, secret)

    db, err := sql.Open()
}