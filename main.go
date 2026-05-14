package main

import (
	"log"
	"net/http"

	"go-web-app/handlers"
)

func main() {

	// Static files
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Routes
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/toys", handlers.ToysHandler)
	http.HandleFunc("/cart", handlers.CartHandler)
	http.HandleFunc("/add-to-cart", handlers.AddToCartHandler)
	http.HandleFunc("/place-order", handlers.PlaceOrderHandler)
	http.HandleFunc("/about", handlers.AboutHandler)
	http.HandleFunc("/contact", handlers.ContactHandler)

	log.Println("Toy Shop Server started on :8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}