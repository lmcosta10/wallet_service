# Digital Wallet API

A RESTful API for a digital wallet system, built with Go.

WARNING: Docker files contain mock credentials and /db/schema.sql contains mock data for development only.

## Tech Stack
- Go (with net/http and Gin)
- PostgreSQL
- Docker & Docker Compose

## Endpoints

- GET /users/:id: get user whose id is ":id"
- POST /wallets: transaction. Required fields: FromWalletID, ToWalletID, Amount.

### Examples

- POST /wallets:
```
curl -X POST -H "Content-Type: application/json" -d '{"FromWalletID": 1, "ToWalletID": 2, "Amount": 100}' localhost:7878/wallets
```

## Running Locally

Run the project locally with:

```
docker compose up --build
```

The API will be available at http://localhost:7878

## Database Access

To access the db directly in the terminal:

```
docker compose exec db psql -U postgres -d postgresdb
```

