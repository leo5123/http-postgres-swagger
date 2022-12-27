package routes

import (
	"log"
	"net/http"
	"sipubtech-desafio/controllers"
	"sipubtech-desafio/middleware"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)

	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/tarefas", controllers.TodasAsTarefas).Methods("Get")
	r.HandleFunc("/api/tarefas/{id}", controllers.TarefaPorId).Methods("Get")
	r.HandleFunc("/api/tarefas/tarefa", controllers.PersonalidadeFiltrada).Methods("Post")
	r.HandleFunc("/api/tarefas", controllers.CriaUmaTarefa).Methods("Post")
	r.HandleFunc("/api/tarefas/{id}", controllers.EditarUmaTarefa).Methods("Put")
	r.HandleFunc("/api/tarefas/{id}", controllers.DeletaUmaTarefa).Methods("Delete")
	headersOk := handlers.AllowedHeaders([]string{"Accept", "Accept-Language", "Content-Type", "Content-Language", "Origin"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "DELETE"})

	// start server listen
	// with error handling
	log.Fatal(http.ListenAndServe(":5050", handlers.CORS(originsOk, headersOk, methodsOk)(r)))

}
