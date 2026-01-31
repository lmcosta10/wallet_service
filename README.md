# Digital Wallet API

A RESTful API for a digital wallet system, built with Go.

## Tech Stack
- Go (with net/http and Gin)
- PostgreSQL
- Docker & Docker Compose

## Running Locally

WARNING: Docker files contain mock credentials and /db/schema.sql contains mock data for development only.

Run the project locally with:

```
docker compose up --build
```

The API will be available at http://localhost:7878

### Database Access

To access the db directly in the terminal:

```
docker compose exec db psql -U postgres -d postgresdb
```

