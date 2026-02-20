# Ecommerce Backend

A RESTful backend API for an ecommerce application built with Go.

## Tech Stack

- **Language:** Go
- **Framework:** [Gin](https://github.com/gin-gonic/gin)
- **ORM:** [GORM](https://gorm.io)
- **Database:** PostgreSQL
- **Live Reload:** [Air](https://github.com/air-verse/air)

## Getting Started

### Prerequisites

- Go 1.21+
- PostgreSQL

### Setup

1. Clone the repository:
   ```bash
   git clone https://github.com/Dav16Akin/ecommerce-backend.git
   cd ecommerce-backend
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Create a `.env` file in the project root:
   ```env
   DATABASE_URL=postgres://<user>:<password>@localhost:5432/ecommerce?sslmode=disable
   ```

4. Run the server:
   ```bash
   go run main.go
   ```

   Or with live reload using Air:
   ```bash
   air
   ```

The server starts on port **8000**.

## API Endpoints

All routes are prefixed with `/v1`.

### Users

| Method | Endpoint        | Description        |
|--------|-----------------|--------------------|
| POST   | `/v1/user/`     | Create a new user  |
| GET    | `/v1/user/:id`  | Get a user by ID   |
| PATCH  | `/v1/user/:id`  | Update a user      |
| DELETE | `/v1/user/:id`  | Delete a user      |

### Create User — Request Body

```json
{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "secret123",
  "phone_number": "08012345678"
}
```

### Update User — Request Body

All fields are optional.

```json
{
  "name": "Jane Doe",
  "email": "jane@example.com",
  "phone_number": "08098765432"
}
```
