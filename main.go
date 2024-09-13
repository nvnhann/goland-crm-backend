package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"goland/handlers"
)

// main initializes the HTTP server and routes, and starts listening on port 8080.
// It defines routes for handling customer-related operations using handler functions from the handlers package.
func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", handlers.HomeHandler).Methods("GET")
	router.HandleFunc("/customers", handlers.GetCustomers).Methods("GET")
	router.HandleFunc("/customers/{id}", handlers.GetCustomer).Methods("GET")
	router.HandleFunc("/customers", handlers.AddCustomer).Methods("POST")
	router.HandleFunc("/customers/{id}", handlers.UpdateCustomer).Methods("PUT")
	router.HandleFunc("/customers/{id}", handlers.DeleteCustomer).Methods("DELETE")
	router.HandleFunc("/customers-batch", handlers.UpdateCustomersBatch).Methods("PUT")

	fmt.Println("Server started at :8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
