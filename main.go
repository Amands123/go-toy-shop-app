import (
    "log"
    "net/http"

    "go-web-app/handlers"
)

func main() {

    fs := http.FileServer(http.Dir("./static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    http.HandleFunc("/", handlers.HomeHandler)
    http.HandleFunc("/toys", handlers.ToysHandler)
    http.HandleFunc("/about", handlers.AboutHandler)
    http.HandleFunc("/contact", handlers.ContactHandler)

    log.Println("Server started on :8080")
    http.ListenAndServe(":8080", nil)
}