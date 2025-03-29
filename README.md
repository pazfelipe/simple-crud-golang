# Simple Product CRUD API

A simple RESTful API built with Go and Gin framework for managing products. This project demonstrates basic CRUD operations (Create, Read, Update, Delete) using a PostgreSQL database.

## Features

- List all products
- Get product details by ID
- Create new products
- Delete products

## Tech Stack

- Go (Golang)
- Gin Web Framework
- PostgreSQL
- Docker & Docker Compose

## Project Structure

The project follows a clean architecture approach with the following components:

- **Controller**: Handles HTTP requests and responses
- **UseCase**: Contains business logic
- **Repository**: Manages data access
- **Model**: Defines data structures

## API Endpoints

| Method | Endpoint      | Description           |
| ------ | ------------- | --------------------- |
| GET    | /ping         | Health check endpoint |
| GET    | /products     | Get all products      |
| GET    | /products/:id | Get product by ID     |
| POST   | /products     | Create a new product  |
| DELETE | /products/:id | Delete a product      |

## Getting Started

### Prerequisites

- Go 1.24.1 or higher
- Docker and Docker Compose

### Running Locally

1. Clone the repository:

```bash
git clone git@github.com:pazfelipe/simple-crud-golang.git
cd simple-crud-golang
```

2. Start the PostgreSQL database using Docker Compose:

```bash
docker-compose up -d postgres
```

3. Run the application:

```bash
go run cmd/main.go
```

4. The API will be available at http://localhost:8000

### Using Docker

You can also run the entire application using Docker:

```bash
docker compose up -d
```

## API Usage Examples

### Create a product

```bash
curl -X POST http://localhost:8000/products \
  -H "Content-Type: application/json" \
  -d '{"name":"Product Name","price":29.99}'
```

### Get all products

```bash
curl http://localhost:8000/products
```

### Get product by ID

```bash
curl http://localhost:8000/products/1
```

### Delete product

```bash
curl -X DELETE http://localhost:8000/products/1
```
