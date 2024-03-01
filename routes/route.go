package routes

import (
	"e_real_estate/controllers"
	"e_real_estate/db"
	"e_real_estate/middleware"
	"e_real_estate/services"
	"fmt"

	"github.com/gorilla/mux"
)




func Router() *mux.Router{
	database, err := db.NewDb() 
	if err != nil {
		fmt.Println(err)
	}
	userService := services.NewUserService(database)
	userController := controllers.NewUserController(userService)

	 router := mux.NewRouter()

	 router.HandleFunc("/", controllers.Test).Methods("GET")
	 router.HandleFunc("/user", userController.CreateAccount).Methods("POST")
	 router.HandleFunc("/dd", middleware.CreateStock).Methods("GET")

	return router
	}