module github.com/AdairHdz/auth-server

go 1.16

require (
	github.com/go-redis/redis/v8 v8.11.3
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/golang-jwt/jwt/v4 v4.1.0
	github.com/joho/godotenv v1.4.0 // indirect
	github.com/rs/cors v1.8.0
	golang.org/x/crypto v0.0.0-20210921155107-089bfa567519 // indirect
)

replace github.com/AdairHdz/auth-server => ./
