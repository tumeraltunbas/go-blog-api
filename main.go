package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/tumeraltunbas/go-blog/config"
	"github.com/tumeraltunbas/go-blog/routes"
	"github.com/tumeraltunbas/go-blog/utils"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln(err)
		panic(err)
	}

	router := chi.NewRouter()
	router.Mount("/api", routes.IndexRouter())

	fmt.Println(config.Get().Port)
	http.ListenAndServe(utils.GetPort(), router)
}
