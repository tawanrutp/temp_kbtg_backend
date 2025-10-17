# API Usage Guide

> Complete API documentation for KBTG Backend API with request/response examples.

## 📋 Table of Contents

- [Base URL](#base-url)
- [Authentication](#authentication)
- [Response Format](#response-format)
- [Error Handling](#error-handling)
- [Endpoints](#endpoints)
  - [General](#general-endpoints)
  - [Customers](#customer-endpoints)
  - [Orders](#order-endpoints)
- [Testing Examples](#testing-examples)

## 🌐 Base URL

```
http://localhost:3000
```

## 🔐 Authentication

Currently, this API does not require authentication. All endpoints are publicly accessible.

> **Note**: In production, implement proper authentication (JWT, OAuth, etc.)

## 📊 Response Format

### Success Response

All successful responses follow this format:

```json
{
  "success": true,
  "data": { ... }
}
```

### Error Response

All error responses follow this format:

```json
{
  "error": "Error message description"
}
```

## ⚠️ Error Handling

| HTTP Status | Description |
|-------------|-------------|
| `200 OK` | Request successful |
| `201 Created` | Resource created successfully |
| `400 Bad Request` | Invalid request body or parameters |
| `404 Not Found` | Resource not found |
| `500 Internal Server Error` | Server error |

---

## 📡 Endpoints

### General Endpoints

#### Get Welcome Message

```http
GET /
```

**Response:**
```json
{
  "message": "Welcome to KBTG Backend API",
  "version": "1.0.0",
  "endpoints": {
    "customers": "/api/v1/customers",
    "orders": "/api/v1/orders"
  }
}
```

#### Get Hello World

```http
GET /hello
```

**Response:**
```json
{
  "message": "Hello World",
  "status": "success"
}
```

---

### Customer Endpoints

#### 1. Get All Customers

```http
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

#### 2. Get Customer by ID

```http
GET /api/v1/customers/:id
```

**Parameters:**
| Parameter | Type | Description |
|-----------|------|-------------|
| `id` | integer | Customer ID |

**Example:**
```bash
GET /api/v1/customers/1
```

**Response (Success):**
```json
{
  "success": true,
  "data": {
    "id": 1,
    "name": "John Doe",
    "email": "john@example.com",
    "phone": "0812345678",
    "created_at": "2025-10-17T10:00:00Z",
    "updated_at": "2025-10-17T10:00:00Z"
  }
}
```

**Response (Not Found):**
```json
{
  "error": "Customer not found"
}
```

#### 3. Create Customer

```http
POST /api/v1/customers
Content-Type: application/json
```

**Request Body:**
```json
{
  "name": "John Doe",
  "email": "john@example.com",
  "phone": "0812345678"
}
```

**Request Body Schema:**
| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `name` | string | ✅ Yes | Customer full name |
| `email` | string | ✅ Yes | Customer email (must be unique) |
| `phone` | string | ❌ No | Customer phone number |

**Response (Success - 201):**
```json
{
  "success": true,
  "data": {
    "id": 1,
    "name": "John Doe",
    "email": "john@example.com",
    "phone": "0812345678",
    "created_at": "2025-10-17T10:00:00Z",
    "updated_at": "2025-10-17T10:00:00Z"
  }
}
```

**Response (Error - 400):**
```json
{
  "error": "Invalid request body"
}
```

#### 4. Update Customer

```http
PUT /api/v1/customers/:id
Content-Type: application/json
```

**Parameters:**
| Parameter | Type | Description |
|-----------|------|-------------|
| `id` | integer | Customer ID |

**Request Body:**
```json
{
  "name": "John Updated",
  "email": "john.updated@example.com",
  "phone": "0887654321"
}
```

**Response (Success):**
```json
{
  "success": true,
  "data": {
    "id": 1,
    "name": "John Updated",
    "email": "john.updated@example.com",
    "phone": "0887654321",
    "created_at": "2025-10-17T10:00:00Z",
    "updated_at": "2025-10-17T10:30:00Z"
  }
}
```

#### 5. Delete Customer

```http
DELETE /api/v1/customers/:id
```

**Parameters:**
| Parameter | Type | Description |
|-----------|------|-------------|
| `id` | integer | Customer ID |

**Response (Success):**
```json
{
  "success": true,
  "message": "Customer deleted successfully"
}
```

---

### Order Endpoints

#### 1. Get All Orders

```http
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

#### 2. Get Order by ID

```http
GET /api/v1/orders/:id
```

**Parameters:**
| Parameter | Type | Description |
|-----------|------|-------------|
| `id` | integer | Order ID |

**Example:**
```bash
GET /api/v1/orders/1
```

**Response (Success):**
```json
{
  "success": true,
  "data": {
    "id": 1,
    "customer_id": 1,
    "order_date": "2025-10-17T10:00:00Z",
    "status": "pending",
    "total_price": 1500.00,
    "created_at": "2025-10-17T10:00:00Z",
    "updated_at": "2025-10-17T10:00:00Z"
  }
}
```

#### 3. Create Order

```http
POST /api/v1/orders
Content-Type: application/json
```

**Request Body:**
```json
{
  "customer_id": 1,
  "status": "pending",
  "total_price": 1500.00
}
```

**Request Body Schema:**
| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `customer_id` | integer | ✅ Yes | Customer ID (must exist) |
| `status` | string | ❌ No | Order status (default: "pending") |
| `total_price` | decimal | ❌ No | Total order price (default: 0.00) |

**Valid Status Values:**
- `pending` - Order is pending
- `processing` - Order is being processed
- `shipped` - Order has been shipped
- `delivered` - Order has been delivered
- `cancelled` - Order has been cancelled

**Response (Success - 201):**
```json
{
  "success": true,
  "data": {
    "id": 1,
    "customer_id": 1,
    "order_date": "2025-10-17T10:00:00Z",
    "status": "pending",
    "total_price": 1500.00,
    "created_at": "2025-10-17T10:00:00Z",
    "updated_at": "2025-10-17T10:00:00Z"
  }
}
```

#### 4. Update Order

```http
PUT /api/v1/orders/:id
Content-Type: application/json
```

**Parameters:**
| Parameter | Type | Description |
|-----------|------|-------------|
| `id` | integer | Order ID |

**Request Body:**
```json
{
  "status": "completed",
  "total_price": 1600.00
}
```

**Response (Success):**
```json
{
  "success": true,
  "data": {
    "id": 1,
    "customer_id": 1,
    "order_date": "2025-10-17T10:00:00Z",
    "status": "completed",
    "total_price": 1600.00,
    "created_at": "2025-10-17T10:00:00Z",
    "updated_at": "2025-10-17T10:30:00Z"
  }
}
```

#### 5. Delete Order

```http
DELETE /api/v1/orders/:id
```

**Parameters:**
| Parameter | Type | Description |
|-----------|------|-------------|
| `id` | integer | Order ID |

**Response (Success):**
```json
{
  "success": true,
  "message": "Order deleted successfully"
}
```

---

## 🧪 Testing Examples

### Using cURL

#### Create a Customer
```bash
curl -X POST http://localhost:3000/api/v1/customers \
  -H "Content-Type: application/json" \
  -d '{"name":"Alice Smith","email":"alice@example.com","phone":"0891234567"}'
```

#### Get All Customers
```bash
curl http://localhost:3000/api/v1/customers
```

#### Get Customer by ID
```bash
curl http://localhost:3000/api/v1/customers/1
```

#### Update Customer
```bash
curl -X PUT http://localhost:3000/api/v1/customers/1 \
  -H "Content-Type: application/json" \
  -d '{"name":"Alice Updated","email":"alice.updated@example.com"}'
```

#### Delete Customer
```bash
curl -X DELETE http://localhost:3000/api/v1/customers/1
```

#### Create an Order
```bash
curl -X POST http://localhost:3000/api/v1/orders \
  -H "Content-Type: application/json" \
  -d '{"customer_id":1,"status":"pending","total_price":2500.50}'
```

#### Get All Orders
```bash
curl http://localhost:3000/api/v1/orders
```

#### Update Order Status
```bash
curl -X PUT http://localhost:3000/api/v1/orders/1 \
  -H "Content-Type: application/json" \
  -d '{"status":"shipped"}'
```

### Using PowerShell (Windows)

#### Create a Customer
```powershell
Invoke-RestMethod -Uri "http://localhost:3000/api/v1/customers" `
  -Method Post `
  -ContentType "application/json" `
  -Body '{"name":"Bob Johnson","email":"bob@example.com","phone":"0823456789"}'
```

#### Get All Customers
```powershell
Invoke-RestMethod -Uri "http://localhost:3000/api/v1/customers" -Method Get
```

### Using JavaScript (Fetch API)

#### Create a Customer
```javascript
fetch('http://localhost:3000/api/v1/customers', {
  method: 'POST',
  headers: {
    'Content-Type': 'application/json',
  },
  body: JSON.stringify({
    name: 'Carol White',
    email: 'carol@example.com',
    phone: '0834567890'
  })
})
.then(response => response.json())
.then(data => console.log(data))
.catch(error => console.error('Error:', error));
```

#### Get All Customers
```javascript
fetch('http://localhost:3000/api/v1/customers')
  .then(response => response.json())
  .then(data => console.log(data))
  .catch(error => console.error('Error:', error));
```

---

## 📝 Notes

1. **Email Uniqueness**: Customer email must be unique. Duplicate emails will result in an error.
2. **Timestamps**: `created_at` and `updated_at` are automatically managed by the system.
3. **Order Date**: If not provided, `order_date` defaults to the current timestamp.
4. **Data Validation**: All required fields must be provided, or the API will return a 400 error.

---

## 🔗 Related Documentation

- [README](README.md) - Project overview and setup
- [Database Schema](database.md) - Database structure and relationships

---

<div align="center">
  <p>For more information, visit the <a href="https://github.com/tawanrutp/temp_kbtg_backend">GitHub Repository</a></p>
</div>
