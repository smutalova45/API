package main

import (
	"fmt"
	"log"
	"net/http"

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

	// http.HandleFunc("/users", con.Users)
	// fmt.Println("listening at port :8080....")
	// http.ListenAndServe(":8080", nil)

	// http.HandleFunc("/orders", con.Orders)
	// fmt.Println("listening at port :8080....")
	// http.ListenAndServe(":8080", nil)

	// http.HandleFunc("/products", con.Products)
	// fmt.Println("listening at port :8080....")
	// http.ListenAndServe(":8080", nil)

	http.HandleFunc("/orderproducts", con.OrderProducts)
	fmt.Println("Listeningg at port :808....")
	http.ListenAndServe(":808", nil)

}
