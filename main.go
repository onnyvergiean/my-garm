package main

import (
	"my-garm/database"
	"my-garm/router"
	"os"
)


func main() {	

	database.StartDB()
		
	var PORT = os.Getenv("PORT")
	r:= router.StartApp()

	r.Run(":" + PORT)
}