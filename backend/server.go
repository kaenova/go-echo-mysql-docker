package main

import (
	"go_echo_rest/db"
	"go_echo_rest/routes"
)

func main() {
	db.Init()
	e := routes.Init()

	e.Logger.Fatal((e.Start(":1323")))
}
