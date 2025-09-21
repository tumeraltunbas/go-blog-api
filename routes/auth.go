package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/tumeraltunbas/go-blog/handlers"
)

func AuthRouter(userHandler *handlers.UserHandler) http.Handler {
	router := chi.NewRouter()

	router.Post("/register", userHandler.Register)

	return router
}
