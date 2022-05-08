package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"linkedin/internal/util"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	router := mux.NewRouter()
	godotenv.Load(".env")

	missingEnviron := util.MissingEnvironList([]string{
		"PORT",
		"APP_URL",
		"SESSIONS_KEY",
		"MONGO_URL",
		"DATABASE_NAME",
	})

	if len(missingEnviron) > 0 {
		log.Fatalf("the following environ is missing : %v", missingEnviron)
	}
	log.Println("Initializing routes...")
	port := os.Getenv("PORT")

	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), router); err != nil {
		log.Println("Failed starting server | %s", err.Error())
		return
	}
}
