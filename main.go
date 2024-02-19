package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"

	"e_real_estate/controllers"
	"e_real_estate/db"
	"e_real_estate/services"
)




func main() {

	godotenv.Load(".env")
	portString := os.Getenv("PORT")

	if portString == "" {
		log.Fatal("PORT is not found in the environment")
	}

	//database connection
	 database := db.NewDb()
	 err := database.Connect()

	 defer database.Db.Close()
	 err = database.Db.Ping()
	 if err != nil { 
		panic(err)
	 }
	 fmt.Println("we are connected to postgres")
	 
	// multiplexer inspects URL request, redirects to correct handler
	router := mux.NewRouter()

	//handler

	userService := services.NewUserService(database)
	userController:= controllers.NewUserController(userService)

	router.HandleFunc("/", controllers.Test).Methods("Get")
	router.HandleFunc("/user", userController.CreateAccount).Methods("Post")

	handler := cors.Default().Handler(router)

	server := &http.Server{
		Handler: handler,
		Addr: ":" + portString,
	}

	server.Handler = router

	log.Printf("Server starting on port %v", portString)
	err = server.ListenAndServe()
	if err != nil { 
		log.Fatal(err)
	}

	fmt.Println("Port:", portString)
}