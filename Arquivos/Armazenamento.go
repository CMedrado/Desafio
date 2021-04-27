package main

func NovoArmazenamentoContaNaMemoria() *ArmazenamentoContaNaMemoria {
	return &ArmazenamentoContaNaMemoria{map[string]int{}}
}

type ArmazenamentoContaNaMemoria struct {
	armazenamento map[string]int
}

func (a *ArmazenamentoContaNaMemoria) ObterLiga() []Conta {
	var liga []Conta
	for nome, numero := range a.armazenamento {
		liga = append(liga, Conta{nome, numero})
	}
	return liga
}

func (s *ArmazenamentoContaNaMemoria) RegistrarConta(nome string) {
	s.armazenamento[nome]++
}

func (a *ArmazenamentoContaNaMemoria) ObterConta(nome string) int {
	return a.armazenamento[nome]
}
