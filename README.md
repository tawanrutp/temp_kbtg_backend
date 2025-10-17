# KBTG Backend API

Backend project using Go, Fiber framework, and SQLite database with CRUD operations.

## Prerequisites

- Go 1.21 or higher
- Git (optional)

## Features

✅ RESTful API with Fiber framework  
✅ SQLite database with GORM  
✅ CRUD operations for Customers and Orders  
✅ Auto-migration database schema  
✅ CORS enabled  
✅ Request logging middleware  

## Installation

1. Install dependencies:
```bash
go mod tidy
```

## Running the Application

```bash
go run main.go
```

The server will start on `http://localhost:3000`

## Database

The application uses SQLite database (`kbtg.db`) with the following tables:
- **customers** - Customer information
- **delivery_addresses** - Customer delivery addresses
- **orders** - Customer orders
- **line_items** - Order line items

Database will be auto-created and migrated on first run.

## API Endpoints

### General
- `GET /` - Welcome message and API info
- `GET /hello` - Hello World endpoint

### Customers (CRUD)
- `GET /api/v1/customers` - Get all customers
- `GET /api/v1/customers/:id` - Get customer by ID
- `POST /api/v1/customers` - Create new customer
- `PUT /api/v1/customers/:id` - Update customer
- `DELETE /api/v1/customers/:id` - Delete customer

### Orders (CRUD)
- `GET /api/v1/orders` - Get all orders
- `GET /api/v1/orders/:id` - Get order by ID
- `POST /api/v1/orders` - Create new order
- `PUT /api/v1/orders/:id` - Update order
- `DELETE /api/v1/orders/:id` - Delete order

See [API_USAGE.md](API_USAGE.md) for detailed API documentation and examples.

## Example Usage

### Create a customer
```bash
curl -X POST http://localhost:3000/api/v1/customers \
  -H "Content-Type: application/json" \
  -d "{\"name\":\"John Doe\",\"email\":\"john@example.com\",\"phone\":\"0812345678\"}"
```

### Get all customers
```bash
curl http://localhost:3000/api/v1/customers
```

### Create an order
```bash
curl -X POST http://localhost:3000/api/v1/orders \
  -H "Content-Type: application/json" \
  -d "{\"customer_id\":1,\"status\":\"pending\",\"total_price\":1500.00}"
```

## Build

To build the executable:

```bash
go build -o app.exe main.go
```

Then run:

```bash
.\app.exe
```

## Project Structure

```
temp_kbtg_backend/
├── main.go              # Main application entry point
├── go.mod               # Go module dependencies
├── kbtg.db              # SQLite database (auto-created)
├── models/
│   └── customer.go      # Database models
├── database/
│   └── database.go      # Database connection and initialization
├── handlers/
│   ├── customer_handler.go  # Customer CRUD handlers
│   └── order_handler.go     # Order CRUD handlers
├── routes/
│   └── routes.go        # API routes setup
├── API_USAGE.md         # Detailed API documentation
└── README.md            # Project documentation
```

## Technologies

- [Go](https://golang.org/) - Programming language
- [Fiber](https://docs.gofiber.io/) - Web framework
- [GORM](https://gorm.io/) - ORM library
- [SQLite](https://www.sqlite.org/) - Database

## Database Schema

Based on the ERD diagram provided:

- **CUSTOMER** → places → **ORDER**
- **CUSTOMER** → uses → **DELIVERY-ADDRESS**
- **ORDER** → contains → **LINE-ITEM**

All relationships are properly defined in the models with foreign keys.
