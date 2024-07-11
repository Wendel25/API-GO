package models

type Pessoa struct {
	Nome      string `json:"name"`
	Idade     int    `json:"age"`
	Profissao string `json:"occupation"`
}
