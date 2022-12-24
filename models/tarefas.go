package models

type Tarefa struct {
	Id     int    `json:"id"`
	Nome   string `json:"nome"`
	Filtro string `json:"filtro"`
	Pronto bool   `json:"pronto"`
}

var Tarefas []Tarefa
