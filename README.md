# KBTG Backend API

Backend project using Go and Fiber framework.

## Prerequisites

- Go 1.21 or higher
- Git (optional)

## Installation

1. Install dependencies:
```bash
go mod download
```

## Running the Application

```bash
go run main.go
```

The server will start on `http://localhost:3000`

## API Endpoints

- `GET /` - Welcome message
- `GET /hello` - Hello World endpoint

## Example Usage

```bash
# Test the hello endpoint
curl http://localhost:3000/hello

# Test the root endpoint
curl http://localhost:3000/
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
├── main.go          # Main application file
├── go.mod           # Go module file
└── README.md        # Project documentation
```

## Technologies

- [Go](https://golang.org/) - Programming language
- [Fiber](https://docs.gofiber.io/) - Web framework
