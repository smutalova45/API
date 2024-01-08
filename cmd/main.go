package main

import (
	"log"

	_ "github.com/lib/pq"
	"main.go/config"
	"main.go/controller"
	"main.go/storage/postgres"
)

func main() {
	cfg := config.Load()
	store, err := postgres.New(cfg)
	if err != nil {
		log.Fatalln("Error connecting", err.Error())
		return
	}
	defer store.DB.Close()
	con := controller.New(store)
	// con.CreateU()
	// con.Getu()
	// con.Deleteu()
	// con.Createp()
	// con.Createo()
	// con.Createopro()
	con.Deleteusers()

}
