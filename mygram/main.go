package main

import (
	"mygram/config"
	"mygram/api/routes"
)

func main() {
	config.ConnectDatabase()

	r := routes.StartApp()
	
	r.Run(":8080")
}

