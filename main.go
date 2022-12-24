package main

import (
	"fmt"
	"sipubtech-desafio/database"
	"sipubtech-desafio/routes"
)

func main() {
	fmt.Println("Iniciando servidor rest com o Go.")

	database.ConectaComBancoDeDados()
	routes.HandleRequest()

}
