package routes

import (
	"fmt"
	"net/http"

	"example.com/go-poc/controllers"
	"example.com/go-poc/middlewares"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)


func HandleRequest(port string) {
	fmt.Printf("Server starting on port: " + port + "\n")
	r := chi.NewRouter()

	// Middlewares
	r.Use(middleware.Logger)
	r.Use(middlewares.ContentTypeMiddleware)

	// Prefix
	r.Route("/api", func(r chi.Router) {

		// Books
		r.Route("/books", func(r chi.Router) {
			r.Get("/", controllers.GetAllBooks)
			r.Get("/{id}", controllers.GetBookById)
			r.Post("/", controllers.CreateBook)
			r.Delete("/{id}", controllers.DeleteBook)
			r.Put("/{id}", controllers.UpdateBook)
		})

		// Authors
		r.Route("/authors", func(r chi.Router) {
			r.Get("/", func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("authors"))
			})
		})
	})

	http.ListenAndServe(":" + port, r)
}