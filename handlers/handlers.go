package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"goland/models"
	"io/ioutil"
	"net/http"
	"strconv"
)

// HomeHandler serves the "static/index.html" file as the response to the HTTP request.
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/index.html")
}

// GetCustomers handles the HTTP request to return a list of all customers in JSON format.
func GetCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.Customers)
}

// GetCustomer handles the HTTP request to retrieve a specific customer by their ID.
func GetCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for _, customer := range models.Customers {
		if customer.ID == id {
			json.NewEncoder(w).Encode(customer)
			return
		}
	}

	http.Error(w, "Customer not found", http.StatusNotFound)
}

// AddCustomer handles the HTTP request to add a new customer to the system and responds with the newly created customer in JSON format.
func AddCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newCustomer models.Customer
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &newCustomer)

	models.CustomerMux.Lock()
	newCustomer.ID = models.NextID
	models.NextID++
	models.CustomerMux.Unlock()

	models.Customers = append(models.Customers, newCustomer)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newCustomer)
}

// UpdateCustomer handles the HTTP request to update an existing customer by their ID.
// The customer data is updated with the JSON payload from the request body.
// If the customer ID is not found, the function responds with a "Customer not found" error.
func UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	var updatedCustomer models.Customer
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &updatedCustomer)

	for i, customer := range models.Customers {
		if customer.ID == id {
			updatedCustomer.ID = customer.ID
			models.Customers[i] = updatedCustomer
			json.NewEncoder(w).Encode(updatedCustomer)
			return
		}
	}

	http.Error(w, "Customer not found", http.StatusNotFound)
}

// DeleteCustomer handles the HTTP request to delete a customer by their ID.
// It finds the customer within the mock database and removes it if found.
// Responds with the updated customers list in JSON format if successful, or a "Customer not found" error.
func DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for i, customer := range models.Customers {
		if customer.ID == id {
			models.Customers = append(models.Customers[:i], models.Customers[i+1:]...)
			json.NewEncoder(w).Encode(models.Customers)
			return
		}
	}

	http.Error(w, "Customer not found", http.StatusNotFound)
}

// UpdateCustomersBatch handles the batch update of multiple customer records.
// It reads the JSON payload from the request body, locks the mock database for writing,
// updates the customers, and responds with the updated list of customers in JSON format.
func UpdateCustomersBatch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var updatedCustomers []models.Customer
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &updatedCustomers)

	models.CustomerMux.Lock()
	defer models.CustomerMux.Unlock()

	for _, updatedCustomer := range updatedCustomers {
		for i, customer := range models.Customers {
			if customer.ID == updatedCustomer.ID {
				models.Customers[i] = updatedCustomer
				break
			}
		}
	}

	json.NewEncoder(w).Encode(models.Customers)
}
