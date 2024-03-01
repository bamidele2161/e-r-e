package middleware

import (
	"e_real_estate/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func CreateStock(w http.ResponseWriter, r *http.Request) {

	// create an empty stock of type models.Stock
	var stock models.Stock

	// decode the json request to stock
	err := json.NewDecoder(r.Body).Decode(&stock)

	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	// call insert stock function and pass the stock


fmt.Println("dddddd", stock)
}
