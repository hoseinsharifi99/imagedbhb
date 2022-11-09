package main

import (
	"fmt"
	"imagedb/database"
	"imagedb/handler"
	"imagedb/model"
	"log"
)

func createTables() {
	database.AutoMigrate(model.GameImage{})
}

func main() {
	fmt.Println("Run Server")
	err := database.Connect("localhost", 5432, "postgres", "1", "among_us")
	if err != nil {
		log.Fatalf("%v", err)
	}
	createTables()

	app := handler.Newhandler()
	app.Start()
}
