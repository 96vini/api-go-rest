package main

import (
	"fmt"

	"github.com/96vini/api-go-rest/models"
	"github.com/96vini/api-go-rest/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Nome 1", Historia: "História 1"},
		{Id: 2, Nome: "Nome 2", Historia: "História 2"},
	}
	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleRequest()
}
