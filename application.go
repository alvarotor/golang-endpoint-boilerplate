package main

import (
	"fmt"
	"golang-endpoint-boilerplate/db"
	"golang-endpoint-boilerplate/routing" // --> Change this directory path to the name of your folder / project
)

func main() {
	fmt.Println("Starting application")

	db.InitialMigration()

	routing.InitialRouting()
}
