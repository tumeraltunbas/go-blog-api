package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/tumeraltunbas/go-blog/handlers"
)

func AuthRouter() http.Handler {
	router := chi.NewRouter()

	router.Post("/login", handlers.Login)

	return router
}
