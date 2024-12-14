# Simple Go CRUD APP
Simple REST app in Golang 

## Features:
- Use of  Gin Web Framework
- Connection to Postgres 
- Use of GORM library
- Swagger/OpenAPI documentation
- docker-compose.yml
- App Config (local/docker .env)

## Run
Configure environment variables in `.env` as provided in .env.example and then run the following command to serve the app: 

```
go run cmd/main.go
```

## Tests

```
go test -v ./...
```