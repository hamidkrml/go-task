# Build Stage
FROM golang:alpine AS builder

WORKDIR /app
# Kaynak kodları kopyala
COPY . .

# Yerel go olmadığı için manuel indiriyoruz ve go.mod'u güncelliyoruz
RUN go get github.com/spf13/viper
RUN go get github.com/golang-jwt/jwt/v5
RUN go get golang.org/x/crypto/bcrypt
RUN go get github.com/lib/pq
RUN go mod tidy

# Uygulamayı derle
RUN go build -o main cmd/api/main.go

# Run Stage
FROM alpine:latest

WORKDIR /app

# Build aşamasından sadece binary'i al
COPY --from=builder /app/main .
COPY --from=builder /app/.env . 

EXPOSE 8080

CMD ["./main"]
