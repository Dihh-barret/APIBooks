package main

import (
	"github.com/Dihh-barret/APIBooks/database"
	"github.com/Dihh-barret/APIBooks/server"
)

func main() {
	database.StartDB()
	s := server.NewServer()

	s.Run()
}
