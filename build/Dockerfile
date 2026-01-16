FROM golang:alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main.go

FROM alpine:latest

WORKDIR /root/

RUN apk --no-cache add ca-certificates

COPY --from=builder /app/main .

COPY --from=builder /app/configs ./configs

# Ensure .env exists by copying .env.example if .env is missing
#RUN if [ ! -f configs/.env ]; then cp configs/.env.example configs/.env; fi

EXPOSE 8080

CMD ["./main"]