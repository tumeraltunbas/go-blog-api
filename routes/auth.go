package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/tumeraltunbas/go-blog/handlers"
	"github.com/tumeraltunbas/go-blog/middlewares"
	"github.com/tumeraltunbas/go-blog/models/dtos/request"
)

func AuthRouter(userHandler *handlers.UserHandler) http.Handler {
	router := chi.NewRouter()

	router.Group(func(r chi.Router) {
		r.Use(middlewares.DtoValidation[request.RegisterReqDto])
		r.Post("/register", userHandler.Register)
	})

	return router
}
