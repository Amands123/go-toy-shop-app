package handlers

import (
	"html/template"
	"net/http"
)

type Category struct {
	Name  string
	Image string
}

type Toy struct {
	ID          int
	Name        string
	Category    string
	Price       string
	Description string
	Image       string
}

var categories = []Category{
	{
		Name:  "Teddy",
		Image: "/static/images/teddy.jpg",
	},
	{
		Name:  "Robot",
		Image: "/static/images/robot.jpg",
	},
	{
		Name:  "Car",
		Image: "/static/images/car.jpg",
	},
	{
		Name:  "Lego",
		Image: "/static/images/lego.jpg",
	},
}

var toys = []Toy{
	{
		ID:          1,
		Name:        "Teddy Bear",
		Category:    "Teddy",
		Price:       "$20",
		Description: "Soft teddy bear for kids",
		Image:       "/static/images/teddy.jpg",
	},
	{
		ID:          2,
		Name:        "Robot Toy",
		Category:    "Robot",
		Price:       "$40",
		Description: "Smart dancing robot",
		Image:       "/static/images/robot.jpg",
	},
	{
		ID:          3,
		Name:        "Remote Car",
		Category:    "Car",
		Price:       "$35",
		Description: "Battery powered racing car",
		Image:       "/static/images/car.jpg",
	},
	{
		ID:          4,
		Name:        "LEGO Set",
		Category:    "Lego",
		Price:       "$55",
		Description: "Creative building blocks set",
		Image:       "/static/images/lego.jpg",
	},
}

var cart []Toy

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("templates/home.html"))

	err := tmpl.Execute(w, categories)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func ToysHandler(w http.ResponseWriter, r *http.Request) {

	category := r.URL.Query().Get("category")

	var filteredToys []Toy

	for _, toy := range toys {

		if toy.Category == category {
			filteredToys = append(filteredToys, toy)
		}
	}

	tmpl := template.Must(template.ParseFiles("templates/toys.html"))

	err := tmpl.Execute(w, filteredToys)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func AddToCartHandler(w http.ResponseWriter, r *http.Request) {

	toyID := r.URL.Query().Get("id")

	for _, toy := range toys {

		if toyID == "1" && toy.ID == 1 {
			cart = append(cart, toy)
		}

		if toyID == "2" && toy.ID == 2 {
			cart = append(cart, toy)
		}

		if toyID == "3" && toy.ID == 3 {
			cart = append(cart, toy)
		}

		if toyID == "4" && toy.ID == 4 {
			cart = append(cart, toy)
		}
	}

	http.Redirect(w, r, "/cart", http.StatusSeeOther)
}

func CartHandler(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("templates/cart.html"))

	err := tmpl.Execute(w, cart)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func PlaceOrderHandler(w http.ResponseWriter, r *http.Request) {

	cart = nil

	tmpl := template.Must(template.ParseFiles("templates/order-success.html"))

	err := tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("templates/about.html"))

	err := tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func ContactHandler(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("templates/contact.html"))

	err := tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}