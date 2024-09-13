package models

import "sync"

// Customer represents a customer with their personal and contact information.
type Customer struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Role      string `json:"role"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Contacted bool   `json:"contacted"`
}

// Customers is a slice of Customer instances representing a collection of customer records with their respective details.
var Customers = []Customer{
	{ID: 1, Name: "John Doe", Role: "Manager", Email: "john@example.com", Phone: "123-456-7890", Contacted: true},
	{ID: 2, Name: "Jane Smith", Role: "Developer", Email: "jane@example.com", Phone: "987-654-3210", Contacted: false},
	{ID: 3, Name: "Emily Jones", Role: "Designer", Email: "emily@example.com", Phone: "555-555-5555", Contacted: true},
}

// NextID is used to assign a unique identifier to new customer entries.
var NextID = 4

// CustomerMux is a mutex used to manage concurrent access to customer-related operations.
var CustomerMux = &sync.Mutex{}
