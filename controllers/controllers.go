package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sipubtech-desafio/database"
	"sipubtech-desafio/models"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Bem vindo")
}

func TodasAsTarefas(w http.ResponseWriter, r *http.Request) {
	var t []models.Tarefa

	database.DB.Find(&t)
	json.NewEncoder(w).Encode(t)
}

func CriaUmaTarefa(w http.ResponseWriter, r *http.Request) {
	var novaTarefa models.Tarefa
	json.NewDecoder(r.Body).Decode(&novaTarefa)
	database.DB.Create(&novaTarefa)
	json.NewEncoder(w).Encode(novaTarefa)
}

func EditarUmaTarefa(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var tarefa models.Tarefa
	database.DB.First(&tarefa, id)
	json.NewDecoder(r.Body).Decode(&tarefa)
	database.DB.Save(&tarefa)
	json.NewEncoder(w).Encode(tarefa)

}

func DeletaUmaTarefa(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var tarefaParaDeletar models.Tarefa
	database.DB.Delete(&tarefaParaDeletar, id)
	json.NewEncoder(w).Encode(tarefaParaDeletar)
}

func PersonalidadeFiltrada(w http.ResponseWriter, r *http.Request) {
	var tarefa []models.Tarefa
	filtro := r.URL.Query().Get("filtro")

	database.DB.Where("filtro = ?", filtro).Find(&tarefa)
	json.NewEncoder(w).Encode(&tarefa)

}

func TarefaPorId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var tarefa models.Tarefa

	database.DB.First(&tarefa, id)
	json.NewEncoder(w).Encode(tarefa)
}
