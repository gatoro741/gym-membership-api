# 🏋️ Gym Membership & Booking API

A RESTful backend API built with Go for managing gym memberships, class bookings, and user authentication.

## Tech Stack

- **Language:** Go
- **Database:** PostgreSQL 15
- **Router:** Chi
- **Auth:** JWT (argon2id password hashing)
- **Migrations:** golang-migrate
- **Containerization:** Docker + Docker Compose

## Project Structure

```
.
├── cmd/
│   └── main.go
├── internal/
│   ├── config/
│   ├── handlers/
│   ├── middleware/
│   ├── models/
│   ├── router/
│   ├── service/
│   ├── storage/
│   └── worker/
├── migrations/
├── Dockerfile
└── docker-compose.yml
```

## Getting Started

### Prerequisites

- Docker & Docker Compose

### Run with Docker

```bash
# Clone the repository
git clone https://github.com/gatoro741/gym-membership-api.git
cd gym-membership-api

# Create .env file
cp .env.example .env

# Start the application
docker compose up --build
```

The API will be available at `http://localhost:8080`

### Run Migrations

```bash
migrate -path ./migrations -database "postgres://user:password@localhost:5433/gym_db?sslmode=disable" up
```

## Environment Variables

Create a `.env` file in the root directory:

```env
DB_HOST=postgres
DB_PORT=5432
DB_USER=your_user
DB_PASSWORD=your_password
DB_NAME=gym_db
JWT_SECRET=your_secret_key
```

## API Endpoints

### Auth

| Method | Endpoint | Description | Auth Required |
|--------|----------|-------------|---------------|
| POST | `/register` | Register a new user | No |
| POST | `/login` | Login and get JWT token | No |

### Classes

| Method | Endpoint | Description | Auth Required |
|--------|----------|-------------|---------------|
| GET | `/classes` | Get all classes | No |
| POST | `/classes` | Create a class (admin only) | Yes |

### Memberships

| Method | Endpoint | Description | Auth Required |
|--------|----------|-------------|---------------|
| POST | `/memberships` | Buy a membership plan | Yes |
| GET | `/memberships` | Get my membership | Yes |

### Bookings

| Method | Endpoint | Description | Auth Required |
|--------|----------|-------------|---------------|
| POST | `/bookings` | Book a class | Yes |
| GET | `/bookings` | Get my bookings | Yes |
| DELETE | `/bookings/{id}` | Cancel a booking | Yes |

## Authentication

All protected endpoints require a JWT token in the Authorization header:

```
Authorization: Bearer <your_token>
```

Get your token by calling `POST /login`.

## Example Requests

### Register
```json
POST /register
{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "password123"
}
```

### Login
```json
POST /login
{
  "email": "john@example.com",
  "password": "password123"
}
```

### Book a Class
```json
POST /bookings
Authorization: Bearer <token>
{
  "class_id": 1
}
```

## Background Worker

The API includes a background worker that runs every 24 hours to automatically deactivate expired memberships.

## Features

- JWT-based authentication with role-based access control (admin/client)
- Argon2id password hashing
- Membership plan management with automatic expiration
- Class booking with capacity validation
- Background worker for expired membership cleanup
- Graceful shutdown
- Dockerized deployment
