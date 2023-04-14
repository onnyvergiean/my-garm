package main

import (
	"my-garm/database"
	"my-garm/router"
	"os"

	"github.com/gin-contrib/cors"
)


func main() {	

	database.StartDB()
	
	var PORT = os.Getenv("PORT")
	r:= router.StartApp()
	r.Use(cors.Default())
	r.Run(":" + PORT)
}