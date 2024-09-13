package main

import (
	"bytes"
	"encoding/json"
	"goland/handlers"
	"goland/models"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

var router *mux.Router

// setupRouter initializes the router and sets up all the necessary routes for handling customer-related requests.
func setupRouter() {
	router = mux.NewRouter()
	router.HandleFunc("/", handlers.HomeHandler).Methods("GET")
	router.HandleFunc("/customers", handlers.GetCustomers).Methods("GET")
	router.HandleFunc("/customers/{id}", handlers.GetCustomer).Methods("GET")
	router.HandleFunc("/customers", handlers.AddCustomer).Methods("POST")
	router.HandleFunc("/customers/{id}", handlers.UpdateCustomer).Methods("PUT")
	router.HandleFunc("/customers/{id}", handlers.DeleteCustomer).Methods("DELETE")
	router.HandleFunc("/customers-batch", handlers.UpdateCustomersBatch).Methods("PUT")
}

// resetCustomers initializes the Customers slice with default customer entries and sets the NextID to 4.
func resetCustomers() {
	models.Customers = []models.Customer{
		{ID: 1, Name: "John Doe", Role: "Manager", Email: "john@example.com", Phone: "123-456-7890", Contacted: true},
		{ID: 2, Name: "Jane Smith", Role: "Developer", Email: "jane@example.com", Phone: "987-654-3210", Contacted: false},
		{ID: 3, Name: "Emily Jones", Role: "Designer", Email: "emily@example.com", Phone: "555-555-5555", Contacted: true},
	}
	models.NextID = 4
}

// TestHomeHandler tests the HomeHandler function to ensure it returns the correct HTTP status and response body for the root endpoint.
func TestHomeHandler(t *testing.T) {
	setupRouter()
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>CRM Backend</title>
</head>
<body>
<h1>Welcome to CRM Backend</h1>
<p>This is a simple CRM backend built with Go. Available API endpoints:</p>
<ul>
    <li>GET /customers - Retrieve all customers</li>
    <li>GET /customers/{id} - Retrieve a single customer by ID</li>
    <li>POST /customers - Create a new customer</li>
    <li>PUT /customers/{id} - Update an existing customer by ID</li>
    <li>DELETE /customers/{id} - Delete a customer by ID</li>
    <li>PUT /customers-batch - Update customers in batch</li>
</ul>
</body>
</html>`
	if strings.Trim(rr.Body.String(), "\n") != strings.Trim(expected, "\n") {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

// TestGetCustomers verifies that the GetCustomers handler returns the correct HTTP status and the expected number of customers.
func TestGetCustomers(t *testing.T) {
	resetCustomers()
	setupRouter()
	req, err := http.NewRequest("GET", "/customers", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var customers []models.Customer
	err = json.Unmarshal(rr.Body.Bytes(), &customers)
	if err != nil {
		t.Fatal(err)
	}

	expectedLength := 3
	if len(customers) != expectedLength {
		t.Errorf("handler returned wrong number of customers: got %v want %v",
			len(customers), expectedLength)
	}
}

// TestAddCustomer validates the functionality of adding a new customer via a POST request and checks if the response is as expected.
func TestAddCustomer(t *testing.T) {
	resetCustomers()
	setupRouter()

	newCustomer := &models.Customer{
		Name:      "Alice Wonderland",
		Role:      "Tester",
		Email:     "alice@example.com",
		Phone:     "111-222-3333",
		Contacted: false,
	}

	body, err := json.Marshal(newCustomer)
	if err != nil {
		t.Fatal(err)
	}
	req, err := http.NewRequest("POST", "/customers", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	var customer models.Customer
	err = json.Unmarshal(rr.Body.Bytes(), &customer)
	if err != nil {
		t.Fatal(err)
	}

	if customer.Name != newCustomer.Name {
		t.Errorf("handler returned unexpected customer name: got %v want %v",
			customer.Name, newCustomer.Name)
	}
}

// TestUpdateCustomer validates the updating of a customer record via the PUT method.
// It checks that the response status code is 200 OK and verifies the updated customer details in the response.
func TestUpdateCustomer(t *testing.T) {
	resetCustomers()
	setupRouter()

	updatedCustomer := &models.Customer{
		Name:      "John Doe Updated",
		Role:      "Manager Updated",
		Email:     "john_updated@example.com",
		Phone:     "123-456-7890",
		Contacted: true,
	}

	body, err := json.Marshal(updatedCustomer)
	if err != nil {
		t.Fatal(err)
	}
	req, err := http.NewRequest("PUT", "/customers/1", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var customer models.Customer
	err = json.Unmarshal(rr.Body.Bytes(), &customer)
	if err != nil {
		t.Fatal(err)
	}

	if customer.Name != updatedCustomer.Name {
		t.Errorf("handler returned unexpected customer name: got %v want %v",
			customer.Name, updatedCustomer.Name)
	}
}

// TestDeleteCustomer tests the deletion of a customer with ID 1, ensuring the status code is HTTP 200 and two customers remain.
func TestDeleteCustomer(t *testing.T) {
	resetCustomers()
	setupRouter()

	req, err := http.NewRequest("DELETE", "/customers/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var customers []models.Customer
	err = json.Unmarshal(rr.Body.Bytes(), &customers)
	if err != nil {
		t.Fatal(err)
	}

	expectedLength := 2 // Số lượng khách hàng còn lại sau khi xóa
	if len(customers) != expectedLength {
		t.Errorf("handler returned wrong number of customers: got %v want %v",
			len(customers), expectedLength)
	}
}
