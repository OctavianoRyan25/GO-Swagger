package main

import (
	"OctavianoRyan25/GOSwagger/database"
	"OctavianoRyan25/GOSwagger/router"
)

func main()  {
	database.StartDB()
	PORT := ":8080"
	router.StartServer().Run(PORT)
}