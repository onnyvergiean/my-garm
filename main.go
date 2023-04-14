package main

import (
	"my-garm/database"
	"my-garm/router"
)	

func main() {	

	database.StartDB()
	
	r:= router.StartApp()
	r.Run(":8080")
}