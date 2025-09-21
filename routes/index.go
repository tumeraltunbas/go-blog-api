package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/tumeraltunbas/go-blog/handlers"
	"github.com/tumeraltunbas/go-blog/repositories"
	"github.com/tumeraltunbas/go-blog/services"
)

func IndexRouter(pool *pgxpool.Pool) http.Handler {
	router := chi.NewRouter()

	userRepository := repositories.NewUserRepository(pool)
	userService := services.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService)

	router.Mount("/auth", AuthRouter(userHandler))

	return router
}
