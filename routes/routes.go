package routes

import (
	"fmt"
	"net/http"

	"example.com/go-poc/controllers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)


func HandleRequest(port string) {
	fmt.Printf("Server starting on port: " + port + "\n")
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", controllers.GetAllBooks)
	http.ListenAndServe(":" + port, r)
}