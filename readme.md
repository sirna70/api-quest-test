# Book API - Gin + JWT

## Run Project

go mod init book-api
go mod tidy
go run cmd/server/main.go

Server running at:
http://localhost:8080

---

# API List (8 Endpoint)

1. GET /ping
2. POST /echo
3. POST /auth/token
4. POST /books
5. GET /books (JWT Required)
6. GET /books?author=X · GET /books?page=1&limit=2 (JWT Required)
7. GET /books/:id
8. PUT /books/:id
9. DELETE /books/:id

---

# Testing With Curl

## 1. Ping
curl http://localhost:8080/ping

## 2. Generate Token
curl -X POST http://localhost:8080/auth/token

## 3. Create Book
curl -X POST http://localhost:8080/books \
-H "Content-Type: application/json" \
-d '{"title":"Go Clean Code","author":"Rikki"}'

## 4. Get Books (With Token)
curl http://localhost:8080/books \
-H "Authorization: Bearer <TOKEN>"

## 5. Filter
http://localhost:8080/books?author=Rikki&page=1&limit=5