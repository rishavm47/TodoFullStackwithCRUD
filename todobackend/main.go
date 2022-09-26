package main

import (
	"example/data-access/models"
	"example/data-access/routes"
)

func main() {
	models.InitDb()
	routes.RouterInit()
}
