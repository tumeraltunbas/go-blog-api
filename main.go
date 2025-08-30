package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/tumeraltunbas/go-blog/config"
	"github.com/tumeraltunbas/go-blog/database"
	"github.com/tumeraltunbas/go-blog/routes"
	"github.com/tumeraltunbas/go-blog/utils"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln(err)
	}

	databaseConfig := config.Get().Database

	pool := database.Connect(databaseConfig.DbConnectionString, databaseConfig.DbConnectionTimeout)
	defer pool.Close()

	router := chi.NewRouter()
	router.Mount("/api", routes.IndexRouter())

	http.ListenAndServe(utils.GetPort(), router)
}
