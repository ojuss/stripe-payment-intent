package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/stripe/stripe-go/v81"
	"github.com/stripe/stripe-go/v81/paymentintent"
)

func main() {
	godotenv.Load()

	// Set Stripe API key
	stripe.Key = os.Getenv("STRIPE_KEY")

	// Register handlers
	http.HandleFunc("/create-payment-intent", handleCreatePaymentIntent)
	http.HandleFunc("/health", handleHealth)

	// Get the port from the environment variable (Render provides this)
	port := os.Getenv("PORT")
	if port == "" {
		port = "4242" // Default port for local development
	}

	// Start the server
	log.Printf("Listening on port :%s\n", port)
	err := http.ListenAndServe(":"+port, nil)
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

	// Create a new PaymentIntentParams object for stripe
	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(calculateOrderAmount(requestPayload.ProductId)),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
		AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
			Enabled: stripe.Bool(true),
		},
	}

	// Create a new PaymentIntent object for stripe (Calling Stripe's Go API client)
	paymentIntent, err := paymentintent.New(params)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}

	// Upon successful creation of the PaymentIntent
	fmt.Println(paymentIntent.ClientSecret)

	var response struct {
		ClientSecret string `json:"clientSecret"`
	}

	// Encode the response into a buffer
	var buffer bytes.Buffer

	// send the client the Client Secret for the PaymentIntent
	response.ClientSecret = paymentIntent.ClientSecret
	err = json.NewEncoder(&buffer).Encode(response)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}

	// use writer.Header to write the response to the client as a json
	writer.Header().Set("Content-Type", "application/json")

	_, err = io.Copy(writer, &buffer)

	if err != nil {
		fmt.Println(err)
	}
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

func calculateOrderAmount(productId string) int64 {
	switch productId {
	case "Forever Pants":
		return 26000
	case "Forever Shirt":
		return 15500
	case "Forever Shorts":
		return 30000
	}

	return 0
}
