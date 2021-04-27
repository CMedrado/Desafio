package main

import (
	"log"
	"net/http"
)

func main() {
	servidor := NovoServidorConta(NovoArmazenamentoContaNaMemoria())

	if err := http.ListenAndServe(":4020", servidor); err != nil {
		log.Fatal("não foi possivel escutar na porta 4000 ", err)
	}
}
