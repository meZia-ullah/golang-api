# Go API

A RESTful API built with Go for managing tasks, staff, and user authentication.

## Overview

This project is a backend API service that provides endpoints for:
- User authentication (registration & login)
- Staff management
- Todo/task management

The API uses **PostgreSQL** as the database and **GORM** as the ORM for database operations.

## Tech Stack

- **Language**: Go 1.25.6
- **Database**: PostgreSQL 15
- **ORM**: GORM
- **Environment Management**: godotenv
- **Container**: Docker & Docker Compose

## Project Structure

```
.
├── cmd/
│   └── main.go              # Application entry point
├── db/
│   ├── db.go                # Database connection and configuration
│   └── seed.go              # Database seeding logic
├── handlers/
│   ├── auth.go              # Authentication endpoints (register, login)
│   ├── staff.go             # Staff management endpoints
│   └── todo.go              # Todo/task management endpoints
├── models/
│   ├── staff.go             # Staff data model
│   ├── todo.go              # Todo data model
│   └── user.go              # User data model
├── routes/
│   └── index.go             # Route registration and handler mapping
├── docker-compose.yml       # Docker Compose configuration
├── go.mod                   # Go module dependencies
└── .env                     # Environment variables (not included, create locally)
```

## Getting Started

### Prerequisites

- Go 1.25.6 or higher
- Docker & Docker Compose
- PostgreSQL 15 (or use Docker)

### Installation

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd golang-api
   ```

2. **Create `.env` file**
   ```bash
   cp .env.example .env  # If example provided, or create manually
   ```

3. **Start PostgreSQL with Docker Compose**
   ```bash
   docker-compose up -d
   ```

4. **Build and run the application**
   ```bash
   go run ./cmd/main.go
   ```

The API will start on `http://localhost:8080`

## API Endpoints

### Authentication
- `POST /register` - Register a new user
- `POST /login` - Login user

### Staff Management
- `GET /staff` - Get all staff members
- `POST /staff` - Create a new staff member

### Todo Management
- `GET /todos` - Get all todos
- `POST /todos` - Create a new todo
- `PUT /todos/{id}` - Update a todo
- `DELETE /todos/{id}` - Delete a todo

## Environment Variables

Create a `.env` file in the project root with the following variables:

```env
DATABASE_URL=postgres://postgres:postgres@localhost:5432/todo_db
PORT=8080
```

## Database

The application uses PostgreSQL with GORM for ORM operations. Database models are automatically migrated on startup.

### Models
- **User**: User account information
- **Staff**: Staff member details
- **Todo**: Task/todo items

### Starting PostgreSQL

Using Docker Compose:
```bash
docker-compose up -d
```

Default credentials:
- **Username**: postgres
- **Password**: postgres
- **Database**: todo_db
- **Port**: 5432

## Development

### Running Locally

```bash
go run ./cmd/main.go
```

### Building

```bash
go build -o bin/api ./cmd/main.go
```

### Dependencies

All dependencies are defined in `go.mod`. To install:

```bash
go mod download
go mod tidy
```

## Database Seeding

The application automatically seeds the database on startup via the `db.Seed()` function.

## License

This project is licensed under the MIT License.

## Contributing

Feel free to fork this repository and submit pull requests for any improvements.
