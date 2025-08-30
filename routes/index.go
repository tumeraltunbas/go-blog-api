package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func IndexRouter() http.Handler {
	router := chi.NewRouter()

	router.Mount("/auth", AuthRouter())

	return router
}
