# Order API

This project implements a simple Order API using Go, Docker, and PostgreSQL. It follows a clean architecture with separation of concerns for configuration, database, handlers, models, repositories, services, and response formatting.

## Features
- Create and manage orders
- RESTful API endpoints
- PostgreSQL database integration
- Dockerized for easy deployment

## Project Structure
```
docker-compose.yml
Dockerfile
go.mod
cmd/
  server/
    main.go
internal/
  config/
    config.go
  database/
    postgres.go
  handler/
    order_handler.go
  model/
    order.go
  repository/
    order_repository.go
  service/
    order_service.go
pkg/
  response/
    response.go
```

## Getting Started

### Prerequisites
- [Go](https://golang.org/dl/)
- [Docker](https://www.docker.com/get-started)
- [Git](https://git-scm.com/)

### Running with Docker
1. Build and start the services:
   ```sh
   docker-compose up --build
   ```
2. The API will be available at `http://localhost:8080` (or your configured port).

### Running Locally
1. Copy `.env.example` to `.env` and set your environment variables.
2. Start PostgreSQL locally or use Docker Compose.
3. Run the server:
   ```sh
   go run cmd/server/main.go
   ```

## API Endpoints
- `POST /orders` - Create a new order
- `GET /orders/{id}` - Get order by ID
- `GET /orders` - List all orders

## License
MIT
