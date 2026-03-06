# Wallet Core

A digital wallet microservice built with Go and Clean Architecture. Manages clients, accounts, and fund transfers between accounts using an event-driven approach and the Unit of Work pattern.

## Architecture

```
cmd/                        # Application entrypoint
internal/
  entity/                   # Domain entities (Client, Account, Transaction)
  gateway/                  # Repository interfaces (ports)
  database/                 # MySQL repository implementations (adapters)
  usecase/                  # Application use cases
    createclient/
    createaccount/
    createtransaction/
    mocks/                  # Shared test mocks
  event/                    # Domain events
  web/                      # HTTP handlers
pkg/
  events/                   # Event dispatcher (generic)
  uow/                      # Unit of Work implementation
```

The project follows Clean Architecture principles: entities at the core with no external dependencies, use cases orchestrating business logic, and gateways as port interfaces implemented by the database layer.

## Tech Stack

- **Go 1.22** - standard library HTTP server (`net/http`)
- **MySQL 5.7** - persistence
- **Docker Compose** - local infrastructure
- **SQLite** - used in integration tests
- **testify** - test assertions and mocks

## Features

- Client registration
- Account creation linked to a client
- Fund transfers between accounts with atomic balance updates (Unit of Work)
- Event dispatching on transaction creation

## API Endpoints

All endpoints accept and return JSON.

| Method | Path            | Description                     |
|--------|-----------------|---------------------------------|
| POST   | `/clients`      | Create a new client             |
| POST   | `/accounts`     | Create an account for a client  |
| POST   | `/transactions` | Transfer funds between accounts |

### POST /clients

**Request:**
```json
{
  "name": "John Doe",
  "email": "john@mail.com"
}
```

**Response (201):**
```json
{
  "id": "uuid",
  "name": "John Doe",
  "email": "john@mail.com",
  "createdAt": "2024-01-01T00:00:00Z",
  "updatedAt": "2024-01-01T00:00:00Z"
}
```

### POST /accounts

**Request:**
```json
{
  "clientId": "client-uuid"
}
```

**Response (201):**
```json
{
  "id": "account-uuid"
}
```

### POST /transactions

**Request:**
```json
{
  "accountFromId": "source-account-uuid",
  "accountToId": "destination-account-uuid",
  "amount": 100
}
```

**Response (201):**
```json
{
  "id": "transaction-uuid",
  "accountFromId": "source-account-uuid",
  "accountToId": "destination-account-uuid",
  "amount": 100
}
```

## Getting Started

### Prerequisites

- [Go 1.22+](https://go.dev/dl/)
- [Docker](https://docs.docker.com/get-docker/) and Docker Compose

### Setup

1. Clone the repository:
   ```bash
   git clone https://github.com/marcioecom/wallet-core.git
   cd wallet-core
   ```

2. Start MySQL:
   ```bash
   docker compose up -d
   ```
   This creates the `wallet` database and initializes the schema automatically.

3. Install dependencies:
   ```bash
   go mod download
   ```

4. Run the application:
   ```bash
   go run cmd/walletcore/main.go
   ```
   The server starts on port `8080` by default. Set the `PORT` environment variable to override (e.g., `PORT=:3000`).

## Running Tests

```bash
go test ./...
```
