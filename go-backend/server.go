package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {

	// HandleFunc is a method that registers the handler function for the given pattern in the DefaultServeMux
	http.HandleFunc("/create-payment-intent", handleCreatePaymentIntent)
	http.HandleFunc("/health", handleHealth)

	// Main method that listens to the port and serves the requests
	log.Println("Listening to the port localhost:3000")
	var err error = http.ListenAndServe("localhost:3000", nil)

	if err != nil {
		log.Fatal(err)
	}
}

func handleCreatePaymentIntent(writer http.ResponseWriter, request *http.Request) {

	// request.Method checks if the request is a POST request
	if request.Method != "POST" {

		// if true ^ then throw an error back, no need to use writer.Write method

		// StatusMehodNotAllowed is a constant that maps to currect corresponding status code

		// StatusText can accept even an int
		http.Error(writer, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}

	// Define the structure of the data that gets sent to this endpoint
	// Client is required to conform to this structure
	var requestPayload struct {
		ProductId string `json:"product_id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Address1  string `json:"address_1"`
		Address2  string `json:"address_2"`
		City      string `json:"city"`
		State     string `json:"state"`
		Zip       string `json:"zip"`
		Country   string `json:"country"`
	}

    // Decode the request body into the requestPayload struct
    err := json.NewDecoder(request.Body).Decode(&requestPayload)

    // if err gets returned its possible that the request body is not in the correct format
    if err != nil {
        
        // Send an error back to the client
        http.Error(writer, err.Error(), http.StatusInternalServerError)
        return 
    }

    fmt.Println(requestPayload)
    
}

func handleHealth(writer http.ResponseWriter, request *http.Request) {

	// a response is always sent back to the client in []byte form in go
	response := []byte("Server is running")

	// Write is method that writes the response to the client
	_, err := writer.Write(response)

	if err != nil {
		fmt.Println(err)
	}

}
