package main

import (
	"pdv/database"
	"pdv/routes"
)


func main() {
	database.InitDB()
	
	routes.InitRouter()
}
