import (
    "html/template"
    "net/http"
)

type Toy struct {
    ID          int
    Name        string
    Price       string
    Description string
    Image       string
}

var toys = []Toy{
    {
        ID: 1,
        Name: "Teddy Bear",
        Price: "$20",
        Description: "Soft teddy bear for kids",
        Image: "/static/images/teddy.jpg",
    },
    {
        ID: 2,
        Name: "Remote Car",
        Price: "$35",
        Description: "Battery powered racing car",
        Image: "/static/images/car.jpg",
    },
    {
        ID: 3,
        Name: "Robot Toy",
        Price: "$40",
        Description: "Smart dancing robot",
        Image: "/static/images/robot.jpg",
    },
    {
        ID: 4,
        Name: "LEGO Set",
        Price: "$55",
        Description: "Creative building blocks set",
        Image: "/static/images/lego.jpg",
    },
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles("templates/home.html"))
    tmpl.Execute(w, toys)
}

func ToysHandler(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles("templates/toys.html"))
    tmpl.Execute(w, toys)
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles("templates/about.html"))
    tmpl.Execute(w, nil)
}

func ContactHandler(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles("templates/contact.html"))
    tmpl.Execute(w, nil)
}