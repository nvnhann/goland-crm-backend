# CRM Backend

This project is a simple CRM Backend built with Go. It provides a RESTful API to manage customer data.

## Requirements

- Go 1.15 or higher

## Installation

Clone the repository and navigate to the project directory:

```sh
git clone https://github.com/yourusername/crm-backend.git
cd crm-backend
```

## Usage

To start the server, run:

```sh
go run main.go
```

The server will start at `http://localhost:8080`.

## API Endpoints

- `GET /`: Serve the home HTML page with API overview.
- `GET /customers`: Retrieve all customers.
- `GET /customers/{id}`: Retrieve a single customer by ID.
- `POST /customers`: Create a new customer.
- `PUT /customers/{id}`: Update an existing customer by ID.
- `DELETE /customers/{id}`: Delete a customer by ID.
- `PUT /customers/batch`: Update customers in batch.

## Example Requests

### Get All Customers

```sh
curl -X GET http://localhost:8080/customers
```

### Get a Single Customer

```sh
curl -X GET http://localhost:8080/customers/1
```

### Create a New Customer

```sh
curl -X POST http://localhost:8080/customers -H 'Content-Type: application/json' -d '{
  "Name": "Alice Johnson",
  "Role": "Engineer",
  "Email": "alice@example.com",
  "Phone": "444-555-6666",
  "Contacted": true
}'
```

### Update a Customer

```sh
curl -X PUT http://localhost:8080/customers/1 -H 'Content-Type: application/json' -d '{
  "Name": "John Updated",
  "Role": "Manager",
  "Email": "john_updated@example.com",
  "Phone": "123-456-7890",
  "Contacted": false
}'
```

### Delete a Customer

```sh
curl -X DELETE http://localhost:8080/customers/1
```

### Update Customers in Batch

```sh
curl -X PUT http://localhost:8080/customers-batch -H 'Content-Type: application/json' -d '[{
  "ID": 2,
  "Name": "Jane Updated",
  "Role": "Developer",
  "Email": "jane_updated@example.com",
  "Phone": "987-654-3210",
  "Contacted": true
}, {
  "ID": 3,
  "Name": "Emily Updated",
  "Role": "Designer",
  "Email": "emily_updated@example.com",
  "Phone": "555-555-5555",
  "Contacted": false
}]'
```

## Running Tests

To ensure that the API works as expected, you can run the unit tests included in the project.

Navigate to the project directory and run the tests using `go test`:

```sh
go test -v .
```

The `-v` flag is for verbose output, giving you more information about the tests being run. The `./...` syntax tells Go to find and run all tests in the current directory and subdirectories.

Example output:

```sh
=== RUN   TestHomeHandler
--- PASS: TestHomeHandler (0.00s)
=== RUN   TestGetCustomers
--- PASS: TestGetCustomers (0.00s)
=== RUN   TestAddCustomer
--- PASS: TestAddCustomer (0.00s)
=== RUN   TestUpdateCustomer
--- PASS: TestUpdateCustomer (0.00s)
=== RUN   TestDeleteCustomer
--- PASS: TestDeleteCustomer (0.00s)
PASS
ok  	goland	0.025s
```

By running these tests, you can verify that all endpoints are working correctly.