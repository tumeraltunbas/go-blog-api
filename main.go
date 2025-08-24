package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/tumeraltunbas/go-blog/constants/errors"
	"github.com/tumeraltunbas/go-blog/utils"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalln(err)
		panic(errors.NewInternalServerError())
	}

	http.ListenAndServe(utils.GetPort(), nil)
}
