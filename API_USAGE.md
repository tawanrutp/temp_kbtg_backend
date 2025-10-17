# API Usage Guide

## Base URL
```
http://localhost:3000
```

## Endpoints

### Root
- **GET** `/` - Welcome message and API info
- **GET** `/hello` - Hello World endpoint

### Customers

#### Get all customers
```bash
GET /api/v1/customers
```

**Response:**
```json
{
  "success": true,
  "data": [
    {
      "id": 1,
      "name": "John Doe",
      "email": "john@example.com",
      "phone": "0812345678",
      "created_at": "2025-10-17T10:00:00Z",
      "updated_at": "2025-10-17T10:00:00Z"
    }
  ]
}
```

#### Get customer by ID
```bash
GET /api/v1/customers/:id
```

#### Create customer
```bash
POST /api/v1/customers
Content-Type: application/json

{
  "name": "John Doe",
  "email": "john@example.com",
  "phone": "0812345678"
}
```

#### Update customer
```bash
PUT /api/v1/customers/:id
Content-Type: application/json

{
  "name": "John Updated",
  "email": "john.updated@example.com",
  "phone": "0887654321"
}
```

#### Delete customer
```bash
DELETE /api/v1/customers/:id
```

### Orders

#### Get all orders
```bash
GET /api/v1/orders
```

**Response:**
```json
{
  "success": true,
  "data": [
    {
      "id": 1,
      "customer_id": 1,
      "order_date": "2025-10-17T10:00:00Z",
      "status": "pending",
      "total_price": 1500.00,
      "created_at": "2025-10-17T10:00:00Z",
      "updated_at": "2025-10-17T10:00:00Z"
    }
  ]
}
```

#### Get order by ID
```bash
GET /api/v1/orders/:id
```

#### Create order
```bash
POST /api/v1/orders
Content-Type: application/json

{
  "customer_id": 1,
  "status": "pending",
  "total_price": 1500.00
}
```

#### Update order
```bash
PUT /api/v1/orders/:id
Content-Type: application/json

{
  "status": "completed",
  "total_price": 1600.00
}
```

#### Delete order
```bash
DELETE /api/v1/orders/:id
```

## Testing with cURL

### Create a customer
```bash
curl -X POST http://localhost:3000/api/v1/customers \
  -H "Content-Type: application/json" \
  -d "{\"name\":\"Alice Smith\",\"email\":\"alice@example.com\",\"phone\":\"0891234567\"}"
```

### Get all customers
```bash
curl http://localhost:3000/api/v1/customers
```

### Create an order
```bash
curl -X POST http://localhost:3000/api/v1/orders \
  -H "Content-Type: application/json" \
  -d "{\"customer_id\":1,\"status\":\"pending\",\"total_price\":2500.50}"
```

### Get all orders
```bash
curl http://localhost:3000/api/v1/orders
```

### Update customer
```bash
curl -X PUT http://localhost:3000/api/v1/customers/1 \
  -H "Content-Type: application/json" \
  -d "{\"name\":\"Alice Updated\",\"email\":\"alice.updated@example.com\"}"
```

### Delete customer
```bash
curl -X DELETE http://localhost:3000/api/v1/customers/1
```

## Database Schema

### Customer
- `id` (uint, primary key)
- `name` (string, required)
- `email` (string, unique, required)
- `phone` (string)
- `created_at` (timestamp)
- `updated_at` (timestamp)

### DeliveryAddress
- `id` (uint, primary key)
- `customer_id` (uint, foreign key)
- `address` (string, required)
- `city` (string)
- `postal_code` (string)
- `country` (string)
- `is_default` (boolean)

### Order
- `id` (uint, primary key)
- `customer_id` (uint, foreign key)
- `order_date` (timestamp, required)
- `status` (string, default: 'pending')
- `total_price` (decimal)
- `created_at` (timestamp)
- `updated_at` (timestamp)

### LineItem
- `id` (uint, primary key)
- `order_id` (uint, foreign key)
- `product_name` (string, required)
- `quantity` (int, required)
- `unit_price` (decimal, required)
- `total_price` (decimal)
