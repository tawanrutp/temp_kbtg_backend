# GitHub Copilot Instructions for KBTG Backend API

## Project Overview
This is a Go-based RESTful API backend using Fiber framework and SQLite database. The system manages users, customers, orders, and a point transfer system with transaction ledger tracking.

## Tech Stack
- **Language**: Go 1.21+
- **Framework**: Fiber v2 (gofiber/fiber/v2)
- **Database**: SQLite with GORM ORM
- **API Spec**: OpenAPI 3.0 (swagger.yml)

## Code Style Guidelines

### 1. File Organization
- Follow the existing project structure:
  ```
  temp_kbtg_backend/
  ├── main.go           # Application entry point
  ├── models/           # Database models
  ├── handlers/         # HTTP request handlers
  ├── routes/           # Route definitions
  └── database/         # Database initialization
  ```

### 2. Naming Conventions
- **Packages**: lowercase, singular (e.g., `handler`, `model`, `database`)
- **Files**: snake_case with suffix (e.g., `user_handler.go`, `transfer_handler.go`)
- **Structs**: PascalCase (e.g., `User`, `Transfer`, `CreateTransferRequest`)
- **Functions**: PascalCase for exported, camelCase for private
- **Variables**: camelCase (e.g., `userID`, `fromUser`, `totalAmount`)

### 3. Handler Function Pattern
All handlers must follow this structure:
```go
func HandlerName(c *fiber.Ctx) error {
    // 1. Parse parameters/body
    // 2. Validate input
    // 3. Database operation
    // 4. Return standardized response
    
    // Success response:
    return c.JSON(fiber.Map{
        "success": true,
        "data": result,
    })
    
    // Error response:
    return c.Status(statusCode).JSON(fiber.Map{
        "error": "Error message",
    })
}
```

### 4. Response Format Standards
**Success Response (200/201)**:
```json
{
    "success": true,
    "data": { /* result object or array */ }
}
```

**Error Response (400/404/500)**:
```json
{
    "error": "Descriptive error message"
}
```

### 5. Database Model Pattern
```go
type ModelName struct {
    ID        uint           `json:"id" gorm:"primaryKey"`
    // Fields...
    CreatedAt time.Time      `json:"created_at"`
    UpdatedAt time.Time      `json:"updated_at"`
    DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
```

### 6. CRUD Operations Guidelines
- **GET all**: Return array in `data` field
- **GET by ID**: Return single object in `data` field, 404 if not found
- **POST**: Return created object with 201 status
- **PUT**: Return updated object
- **DELETE**: Return success message

### 7. Route Registration Pattern
```go
// In routes/routes.go
group := api.Group("/resource")
group.Get("/", handlers.GetResources)
group.Get("/:id", handlers.GetResource)
group.Post("/", handlers.CreateResource)
group.Put("/:id", handlers.UpdateResource)
group.Delete("/:id", handlers.DeleteResource)
```

### 8. Database Query Patterns
```go
// Find all
var items []models.Item
database.DB.Find(&items)

// Find by ID
var item models.Item
database.DB.First(&item, id)

// Create
database.DB.Create(&item)

// Update
database.DB.Save(&item)

// Delete (soft delete)
database.DB.Delete(&item)

// Preload relationships
database.DB.Preload("RelationName").Find(&items)
```

### 9. Transaction Pattern (for transfers)
```go
err := database.DB.Transaction(func(tx *gorm.DB) error {
    // Step 1: Operation
    if err := tx.Create(&item).Error; err != nil {
        return err
    }
    
    // Step 2: Related operation
    if err := tx.Updates(&related).Error; err != nil {
        return err
    }
    
    return nil
})
```

### 10. Idempotency Pattern (for transfers)
```go
// Check for existing transfer with same idempotency key
var existing models.Transfer
result := database.DB.Where("idempotency_key = ?", key).First(&existing)

if result.Error == nil {
    // Return existing transfer
    return c.JSON(fiber.Map{"success": true, "data": existing})
}
```

## API Design Principles

### 1. RESTful Endpoints
- Use plural nouns: `/users`, `/transfers`, `/customers`
- Use HTTP methods correctly: GET (read), POST (create), PUT (update), DELETE (delete)
- Use path parameters for IDs: `/users/:id`
- Use query parameters for filtering: `/transfers?user_id=1&status=completed`

### 2. Status Codes
- **200**: Successful GET/PUT/DELETE
- **201**: Successful POST (creation)
- **400**: Bad request (validation error)
- **404**: Resource not found
- **500**: Internal server error

### 3. Validation Rules
- Required fields must be validated before database operations
- Return 400 with clear error message for validation failures
- Check for duplicate entries (e.g., idempotency keys)
- Validate business logic (e.g., sufficient balance for transfers)

## Special Features

### 1. Point Transfer System
- Must use transactions for atomic operations
- Check sender balance before transfer
- Support idempotency with `idempotency_key`
- Update both sender and receiver balances
- Create ledger entries for both users
- Status: "pending", "completed", "failed"

### 2. Ledger Tracking
- Record all point movements (credit/debit)
- Include reference to transfer
- Track running balance
- Support filtering by user

### 3. CORS Configuration
```go
app.Use(cors.New(cors.Config{
    AllowOrigins: "*",
    AllowMethods: "GET,POST,PUT,DELETE",
    AllowHeaders: "Origin, Content-Type, Accept",
}))
```

## Common Patterns to Follow

### Error Handling
```go
if err := database.DB.First(&item, id).Error; err != nil {
    return c.Status(404).JSON(fiber.Map{
        "error": "Item not found",
    })
}
```

### Request Body Parsing
```go
type RequestBody struct {
    Field1 string `json:"field1"`
    Field2 int    `json:"field2"`
}

body := new(RequestBody)
if err := c.BodyParser(body); err != nil {
    return c.Status(400).JSON(fiber.Map{
        "error": "Invalid request body",
    })
}
```

### Query Parameters
```go
userID := c.Query("user_id")
status := c.Query("status")
```

### Path Parameters
```go
id := c.Params("id")
```

## Testing Commands
- Run server: `go run main.go`
- Build: `go build -o app.exe main.go`
- Test: Use provided batch files in project root

## Important Notes
1. Always use `database.DB` for database operations
2. Enable GORM logging for debugging: `logger.Default.LogMode(logger.Info)`
3. Auto-migrate models in `database/InitDatabase()`
4. Follow swagger.yml specification for API contracts
5. Use fiber.Map for JSON responses
6. Preload relationships when needed to avoid N+1 queries
7. Use soft deletes (gorm.DeletedAt) for all models
8. Always validate business logic before database operations

## When Writing New Code
1. Check existing handlers for similar patterns
2. Follow the same structure and naming conventions
3. Add routes to `routes/routes.go`
4. Add models to appropriate files in `models/`
5. Use consistent error messages
6. Add database migration to `InitDatabase()` if new model
7. Test with provided batch files or curl commands
8. Ensure swagger.yml compliance

## Code Generation Preferences
- Prefer explicit error handling over silent failures
- Use struct tags for JSON and GORM configurations
- Keep handlers thin - extract complex logic to separate functions
- Use meaningful variable names
- Add comments for complex business logic
- Group related functionality together
- Maintain consistency with existing codebase style