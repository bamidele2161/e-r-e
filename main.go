package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/cors"

	"e_real_estate/db"
	"e_real_estate/routes"
)

func main() {

	godotenv.Load(".env")
	portString := os.Getenv("PORT")

	if portString == "" {
		log.Fatal("PORT is not found in the environment")
	}

	//database connection
	database, err := db.NewDb()
	if err != nil {
		fmt.Println(err)
	}
	database.Connect()

	defer database.Db.Close()
	err = database.Db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("we are connected to postgres")

	// multiplexer inspects URL request, redirects to correct handler

	//handler
	router := routes.Router()
	handler := cors.Default().Handler(router)

	server := &http.Server{
		Handler: handler,
		Addr:    ":" + portString,
	}

	server.Handler = router

	log.Printf("Server starting on port %v", portString)
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Port:", portString)
}
