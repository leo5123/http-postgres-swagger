package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"sipubtech-desafio/controllers"
	"sipubtech-desafio/database"
	"sipubtech-desafio/models"
	"strconv"
	"testing"

	"github.com/goccy/go-json"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func SetupDasRotasDeTeste() *mux.Router {
	rotas := mux.NewRouter()
	return rotas
}

var ID int

func CriaTarefaMock() {
	aluno := models.Tarefa{Nome: "Tarefa de teste", Filtro: "React", Pronto: true}
	database.DB.Create(&aluno)
	ID = int(aluno.Id)
}

func DeletaTarefaMock() {
	var aluno models.Tarefa
	database.DB.Delete(&aluno, ID)
}

func TestListandoTodasAsTarefas(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaTarefaMock()
	defer DeletaTarefaMock()
	r := SetupDasRotasDeTeste()
	r.HandleFunc("/api/tarefas", controllers.TodasAsTarefas)
	req, _ := http.NewRequest("GET", "/api/tarefas", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code, "listar todos os alunos")
	// fmt.Println(resposta.Body)
}

func TestBuscaTarefaPeloFiltro(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaTarefaMock()
	defer DeletaTarefaMock()
	r := SetupDasRotasDeTeste()
	r.HandleFunc("/api/tarefas/tarefa", controllers.TarefaFiltrada)
	req, _ := http.NewRequest("POST", "/api/tarefas/tarefa?filtro=React", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code, "listar múltiplas tarefas que correspondem com o filtro")
	// fmt.Println(resposta.Body)
}

func TestBuscaTarefaPeloID(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaTarefaMock()
	defer DeletaTarefaMock()
	r := SetupDasRotasDeTeste()

	r.HandleFunc("/api/tarefas/{id}", controllers.TarefaPorId)
	pathDaBusca := "/api/tarefas/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("GET", pathDaBusca, nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code, "listar única tarefa que corresponde ao id")
	//fmt.Println(resposta.Body)
}

func TestDeletaTarefaPorID(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaTarefaMock()

	r := SetupDasRotasDeTeste()

	r.HandleFunc("/api/tarefas/{id}", controllers.DeletaUmaTarefa)
	pathDaBusca := "/api/tarefas/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("DELETE", pathDaBusca, nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code, "deletar única tarefa que corresponde ao id")
	// fmt.Println(resposta.Body)
}

func TestEditaTarefaPorID(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaTarefaMock()
	defer DeletaTarefaMock()
	r := SetupDasRotasDeTeste()

	r.HandleFunc("/api/tarefas/{id}", controllers.EditarUmaTarefa)

	aluno := models.Tarefa{Nome: "TESTE TESTE TESTE", Filtro: "Javascript", Pronto: true}
	valorJson, _ := json.Marshal(aluno)

	pathDaBusca := "/api/tarefas/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("PUT", pathDaBusca, bytes.NewBuffer(valorJson))
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)

	assert.Equal(t, http.StatusOK, resposta.Code, "editar única tarefa que corresponde ao id")
	//fmt.Println(resposta.Body)
}

func TestCriaUmaTarefa(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaTarefaMock()
	defer DeletaTarefaMock()
	r := SetupDasRotasDeTeste()

	r.HandleFunc("/api/tarefas", controllers.CriaUmaTarefa)

	aluno := models.Tarefa{Nome: "TAREFA CRIADA", Filtro: "Javascript", Pronto: true}
	valorJson, _ := json.Marshal(aluno)

	req, _ := http.NewRequest("PUT", "/api/tarefas", bytes.NewBuffer(valorJson))
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)

	assert.Equal(t, http.StatusOK, resposta.Code, "criar única tarefa que corresponde")
	//fmt.Println(resposta.Body)
}
